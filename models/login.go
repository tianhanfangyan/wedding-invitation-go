package models

import (
	"errors"
	"time"
	"wedding-invitation-go/utils"
)

type Login struct {
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

func (l *Login) LoginAuth(userName, passWord string) (LoginResponse, error) {
	var result LoginResponse

	isExist, err := CheckAuth(userName, passWord)
	if err != nil {
		return result, err
	}

	if !isExist {
		return result, errors.New("用户验证失败")
	}

	token, err := utils.GenerateToken(userName, passWord)
	if err != nil {
		return result, err
	}

	result.AccessToken = token
	result.ExpiresAt = time.Now().Add(time.Hour * 12).Unix()

	return result, nil
}

func CheckAuth(userName, passWord string) (bool, error) {
	var auth Auths

	db := utils.GetDB().Select("id").Where(Auths{Username: userName, Password: passWord}).First(&auth)
	if err := db.Error; err != nil {
		return false, err
	}

	if auth.Id > 0 {
		return true, nil
	}

	return false, nil
}
