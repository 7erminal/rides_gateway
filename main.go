package main

import (
	_ "AMC_gateway/routers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	logs.SetLogger(logs.AdapterFile, `{"filename":"../logs/AMC-gateway.log"}`)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174", "https://api.amc-flowpos.com"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		logs.Info("Handling CORS for origin: %s", ctx.Input.Header("Origin"))
	})

	beego.BConfig.Log.AccessLogs = true

	beego.Run()
}
