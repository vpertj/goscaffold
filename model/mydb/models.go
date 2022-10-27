package models

type User struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Name       string `xorm:"not null VARCHAR(16)"`
	Password   string `xorm:"not null VARCHAR(50)"`
	Createtime int64  `xorm:"not null BIGINT(10)"`
}
