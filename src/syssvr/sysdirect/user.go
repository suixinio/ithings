package sysdirect

import (
	client "github.com/i-Things/things/src/syssvr/client/user"
	server "github.com/i-Things/things/src/syssvr/internal/server/user"

	clientMenu "github.com/i-Things/things/src/syssvr/client/menu"
	clientRole "github.com/i-Things/things/src/syssvr/client/role"
	serverMenu "github.com/i-Things/things/src/syssvr/internal/server/menu"
	serverRole "github.com/i-Things/things/src/syssvr/internal/server/role"

	clientLog "github.com/i-Things/things/src/syssvr/client/log"
	serverLog "github.com/i-Things/things/src/syssvr/internal/server/log"

	clientCommon "github.com/i-Things/things/src/syssvr/client/common"
	serverCommon "github.com/i-Things/things/src/syssvr/internal/server/common"

	clientApi "github.com/i-Things/things/src/syssvr/client/api"
	serverApi "github.com/i-Things/things/src/syssvr/internal/server/api"

	clientAuth "github.com/i-Things/things/src/syssvr/client/auth"
	serverAuth "github.com/i-Things/things/src/syssvr/internal/server/auth"
)

func NewUser(runSvr bool) client.User {
	svcCtx := GetSvcCtx()
	if runSvr {
		RunServer(svcCtx)
	}
	return client.NewDirectUser(svcCtx, server.NewUserServer(svcCtx))
}

func NewRole(runSvr bool) clientRole.Role {
	svcCtx := GetSvcCtx()
	if runSvr {
		RunServer(svcCtx)
	}
	return clientRole.NewDirectRole(svcCtx, serverRole.NewRoleServer(svcCtx))
}

func NewMenu(runSvr bool) clientMenu.Menu {
	svcCtx := GetSvcCtx()
	if runSvr {
		RunServer(svcCtx)
	}
	return clientMenu.NewDirectMenu(svcCtx, serverMenu.NewMenuServer(svcCtx))
}

func NewCommon(runSvr bool) clientCommon.Common {
	svcCtx := GetSvcCtx()
	if runSvr {
		RunServer(svcCtx)
	}
	return clientCommon.NewDirectCommon(svcCtx, serverCommon.NewCommonServer(svcCtx))
}

func NewLog(runSvr bool) clientLog.Log {
	svcCtx := GetSvcCtx()
	if runSvr {
		RunServer(svcCtx)
	}
	return clientLog.NewDirectLog(svcCtx, serverLog.NewLogServer(svcCtx))
}

func NewApi(runSvr bool) clientApi.Api {
	svcCtx := GetSvcCtx()
	if runSvr {
		RunServer(svcCtx)
	}
	return clientApi.NewDirectApi(svcCtx, serverApi.NewApiServer(svcCtx))
}

func NewAuth() clientAuth.Auth {
	userSvc := GetSvcCtx()
	return clientAuth.NewDirectAuth(userSvc, serverAuth.NewAuthServer(userSvc))
}
