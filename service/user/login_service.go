package service

import "goscaffold/model/db"

type LoginService struct {
}

func (l *LoginService) Login(id int) (string, error) {

	orm := db.GetOrm()
	orm.Where("id", id).Get()

	t := "user.login 这是一个用户登录接口"
	return t, nil
}
