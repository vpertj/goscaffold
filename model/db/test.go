package db

type Test struct {
	Id      int64  `xorm:"bigint autoincr pk"` // 等同于bigserial pk
	Name    string `xorm:"varchar(20) notnull default('')"`
	Age     uint8  `xorm:"int notnull default(0)"`
	Address string `xorm:"text"`
}

// TableName 获取表名,如果没有这个方法，默认使用xorm表名转换规则产生表名
func (*Test) TableName() string {
	return "test"
}

type TestTable struct {
	Id      int64  `xorm:"bigint autoincr pk"` // 等同于bigserial pk
	Name    string `xorm:"varchar(20) notnull default('')"`
	Age     uint8  `xorm:"int notnull default(0)"`
	Address string `xorm:"text"`
}
