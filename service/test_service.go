package service

import "goscaffold/model/db"

type TestService struct {
}

func (*TestService) Test(id int) (string, error) {

	orm := db.GetOrm()
	orm.Where("id", id).Get()

	t := "test service. 这是一个测试接口"
	return t, nil
}
