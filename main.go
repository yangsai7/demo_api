package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/yangsai7/demo_api/client"
	"github.com/yangsai7/demo_api/config"
	"github.com/yangsai7/demo_api/dao"
	"github.com/yangsai7/demo_api/server"
)

var (
	ServiceName = "demo_api"
	Version     = "1.0"
)

var (
	webui bool
)

func main() {

	http.DefaultClient.Timeout = time.Second * 120

	ctx := context.Background()

	//init
	config.Init()
	//db
	if err := dao.InitDao(ctx, &mysql.Config{
		User:                 config.GlobalCfg.MySQL.User,
		Passwd:               config.GlobalCfg.MySQL.Passwd,
		Net:                  config.GlobalCfg.MySQL.Net,
		Addr:                 config.GlobalCfg.MySQL.Addr,
		DBName:               config.GlobalCfg.MySQL.DBName,
		AllowNativePasswords: config.GlobalCfg.MySQL.AllowNativePasswords,
		Collation:            config.GlobalCfg.MySQL.Collation,
	}); err != nil {
		panic(err)
	}

	client.InitElastic()

	//http server
	server := server.NewHTTPServer()
	go server.Run(config.GlobalCfg.Http.Addr)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
}
