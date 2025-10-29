package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"strconv"
	"time"

	_ "embed"

	"github.com/golang-jwt/jwt/v5"
)

const AccessTokenTTL = 30 * time.Minute

type JWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	privateBlock, _ := pem.Decode(rawPrivKey)
	if privateBlock == nil {
		log.Fatal("failed to parse PEM block containing the private key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	publicBlock, _ := pem.Decode(rawPubKey)
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

	return &JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (j *JWT) GenerateToken(userID int) (string, error) {
	claims := CustomClaims{
		UserID: strconv.Itoa(userID),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "github.com/ryutaKimu/kakebo",
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
