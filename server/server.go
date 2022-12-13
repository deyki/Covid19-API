package server

import (
	"net/http"

	"github.com/Covid19MicroServiceV2/deyki/v2/controller"
	"github.com/Covid19MicroServiceV2/deyki/v2/service"
)

func Run() {
	service.DownloadDataFromUrl()
	http.ListenAndServe(":8080", controller.GorillaMuxRouter())
}