package main

import (
	"flag"
	"fmt"
	"github.com/go-things/things/src/webapi/internal/vars"

	"github.com/go-things/things/src/webapi/internal/config"
	"github.com/go-things/things/src/webapi/internal/handler"
	"github.com/go-things/things/src/webapi/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/webapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	vars.Svrctx = ctx
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
