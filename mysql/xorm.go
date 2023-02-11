package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// NewEngine 创建 client
func NewEngine(src string, showSQL bool) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", src)

	if err != nil {
		return nil, err
	}

	engine.SetMapper(names.GonicMapper{})

	if err != nil {
		return engine, err
	}

	err = engine.Ping()
	if err != nil {
		return engine, err
	}

	// 控制台打印出生成的SQL语句
	engine.ShowSQL(showSQL)

	return engine, nil
}
