// Package main
package main

import (
	"fmt"
	"github.com/eya46/go-cqhttp/cmd/gocq"
	"github.com/eya46/go-cqhttp/global/terminal"
	"github.com/eya46/go-cqhttp/internal/base"
	"time"

	_ "github.com/eya46/go-cqhttp/db/leveldb"   // leveldb 数据库支持
	_ "github.com/eya46/go-cqhttp/modules/silk" // silk编码模块
	// 其他模块
	// _ "github.com/eya46/go-cqhttp/db/sqlite3"   // sqlite3 数据库支持
	// _ "github.com/eya46/go-cqhttp/db/mongodb"    // mongodb 数据库支持
	// _ "github.com/eya46/go-cqhttp/modules/pprof" // pprof 性能分析
)

func main() {
	fmt.Printf("\033]0;go-cqhttp "+base.Version+" © 2020 - %d Mrs4s"+"\007", time.Now().Year())
	gocq.InitBase()
	gocq.PrepareData()
	gocq.LoginInteract()
	_ = terminal.DisableQuickEdit()
	_ = terminal.EnableVT100()
	gocq.WaitSignal()
	_ = terminal.RestoreInputMode()
}
