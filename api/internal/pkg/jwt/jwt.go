package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const AccessTokenTTL = 30 * time.Minute

type JWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

var (
	instance *JWT
	once     sync.Once
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

func NewJWT() *JWT {
	once.Do(func() {
		var privData, pubData []byte

		// --- 🔍 環境変数を優先 ---
		privEnv := os.Getenv("JWT_PRIVATE_KEY")
		pubEnv := os.Getenv("JWT_PUBLIC_KEY")

		if privEnv != "" && pubEnv != "" {
			privData = []byte(privEnv)
			pubData = []byte(pubEnv)
			log.Println("[JWT] using keys from environment variables")
		} else {
			privData = rawPrivKey
			pubData = rawPubKey
			log.Println("[JWT] using embedded cert/private.pem and cert/public.pem")
		}

		// --- 🔑 秘密鍵のパース ---
		privateBlock, _ := pem.Decode(privData)
		if privateBlock == nil {
			log.Fatal("failed to parse PEM block containing the private key")
		}
		privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
		if err != nil {
			log.Fatal(err)
		}

		// --- 🔑 公開鍵のパース ---
		publicBlock, _ := pem.Decode(pubData)
		if publicBlock == nil {
			log.Fatal("failed to parse PEM block containing the public key")
		}
		pubKeyIface, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
		if err != nil {
			log.Fatal(err)
		}
		publicKey, ok := pubKeyIface.(*rsa.PublicKey)
		if !ok {
			log.Fatal("not RSA public key")
		}

		instance = &JWT{
			privateKey: privateKey,
			publicKey:  publicKey,
		}
	})
	return instance
}

func (j *JWT) GenerateToken(userID int) (string, error) {
	claims := CustomClaims{
		UserID: strconv.Itoa(userID),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "github.com/ryutaKimu/kakebo/api/",
			Subject:   "access_token",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenTTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(j.privateKey)
}

func (j *JWT) VerifyToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims.(*CustomClaims), nil
}
