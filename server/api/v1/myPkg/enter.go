package myPkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	MyApi
}

var (
	myApiService = service.ServiceGroupApp.PkgServiceGroup.MyApiService
)
