package request

import (
	"errors"
	"unicode/utf8"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" {
		return errors.New("名前は必須です")
	}
	if r.Email == "" {
		return errors.New("メールアドレスは必須です")
	}
	if utf8.RuneCountInString(r.Password) < 8 {
		return errors.New("パスワードは8文字以上である必要があります")
	}
	return nil
}
