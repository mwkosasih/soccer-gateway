package main

import (
	"fmt"
	"os"

	"github.com/mwkosasih/soccer-gateway/route"
	"github.com/mwkosasih/soccer-gateway/util"
)

func init() {
	util.Env("./")
}

func main() {
	e := route.Init()
	data, err := util.Json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(fmt.Sprint(err))
	}
	fmt.Println(string(data))

	e.Logger.Fatal(e.Start(":" + os.Getenv("gateway_port")))
}
