package controllers_test

import (
	"tester/src/configuration"
	"tester/src/controllers"
	"tester/src/testerClient"
)

var responseTypesCntr controllers.ResponseTypesController
var bodyParamsCntr controllers.BodyParamsController
var formParamsCntr controllers.FormParamsController
var queryParamsCntr controllers.QueryParamController

func init() {
	client := testerClient.NewClient(
		configuration.ENUMDEFAULT,
		configuration.DefaultConfig(),
	)

	responseTypesCntr = client.ResponseTypesController
	bodyParamsCntr = client.BodyParamsController
	formParamsCntr = client.FormParamsController
	queryParamsCntr = client.QueryParamController
}
