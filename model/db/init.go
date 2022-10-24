package db

import (
	"goscaffold/model/config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

var _orms map[string]*xorm.Engine = nil
var _ormInitLock sync.Mutex

func init() {
	syncTable()
}

func _initOrm() {
	_ormInitLock.Lock()
	defer _ormInitLock.Unlock()
	if _orms != nil {
		return
	}
	conf := config.GetConfig()
	dbs := conf.DB
	if len(dbs) == 0 {
		logrus.Warn("未配置数据库!")
		return
	}
	_orms = make(map[string]*xorm.Engine, len(dbs))
	for name := range dbs {
		db := dbs[name]
		orm, err := xorm.NewEngine(db.Type, db.Url)
		orm.ShowSQL(conf.Server.ShowSql)
		if err != nil {
			logrus.Error(err)
		}
		_orms[name] = orm
	}
}

func syncTable() {
	orm := GetOrm()
	/**
	在数据库结构同步后注意xorm报的警告信息，有些东西无法同步如下
	[xorm] [warn]  2022/05/09 16:29:55.814139 Table test Column age db default is 1, struct default is 0
	[xorm] [warn]  2022/05/09 16:29:55.817355 Table test_table Column age db default is 1, struct default is 0
	*/
	err := orm.Sync2(
	//new(PmEnterprise), //企业表
	//new(PmManager),    //用户表
	//new(PmRole),       //角色表

	)
	if err != nil {
		logrus.Error("数据库结构同步出错:", err)
	} else {
		logrus.Info("数据库结构同步成功.")
	}
}

func GetOrm(name ...string) *xorm.Engine {
	dbName := config.DefaultDBName
	if len(name) > 0 {
		dbName = name[0]
	}
	if _orms == nil {
		_initOrm()
	}
	return _orms[dbName]
}
