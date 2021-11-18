package service

import (
	"errors"
	"fmt"
)

type AuthReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthReq) error {
	fmt.Println(param.Username)
	fmt.Println("-------")
	auth, err := svc.dao.GetAuth(param.Username, param.Password)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("Auth info does not exist.")
}
