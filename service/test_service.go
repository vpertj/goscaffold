package service

type TestService struct {
}

func (*TestService) Test() (string, error) {
	t := "test service. 这是一个测试接口"
	return t, nil
}
