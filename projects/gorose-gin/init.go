package main

import (
	"github.com/gohouse/gorose"
	_ "github.com/mattn/go-sqlite3"
)

// BootGorose 初始化gorose, 单例模式
func BootGorose() {
	var err error
	once.Do(func() {
		engin, err = gorose.Open(&gorose.Config{
			Driver: "sqlite3",
			Dsn:    "db.sqlite",
		})
		if err != nil {
			panic(err.Error())
		}
	})
}

// UserInit 初始化用户表
func UserInit() {
	dbsql := `create table if not exists "users" (
		"uid" integer not null primary key autoincrement,
		"username" text not null default "",
		"age" integer not null default 0
	)`

	affectedRows, err := DB().Execute(dbsql)
	if err != nil {
		panic(err.Error())
	}
	if affectedRows == 0 {
		return
	}
}
