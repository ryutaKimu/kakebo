package request

import "errors"

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginUserRequest) Validate() error {
	if r.Email == "" {
		return errors.New("メールアドレスは必須です")
	}

	if r.Password == "" {
		return errors.New("パスワードは必須です")
	}
	return nil
}
