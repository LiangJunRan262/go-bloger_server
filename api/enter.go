package api

import (
	"bloger_server/api/log_api"
	"bloger_server/api/site_api"
)

type Api struct {
	SiteApi site_api.SiteApi
	LogApi  log_api.LogApi
}

var App = Api{}
