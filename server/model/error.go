package model

import "errors"

var (
	ERROR_UESR_NOTEXISTS = errors.New("用户不存在")
	ERROR_UESR_EXISTS = errors.New("用户已存在")
	ERROR_UESR_PWD = errors.New("密码不正确")
)