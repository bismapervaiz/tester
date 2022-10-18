/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package controllers

import(
	"tester/src/https"

)

type callBuilderFactory interface {
	GetCallBuilder() https.CallBuilderFactory
}

type baseController struct {
	callBuilder    callBuilderFactory
	prepareRequest https.CallBuilderFactory
}

func (b *baseController) callBuilderFactory() https.CallBuilderFactory {
	return b.callBuilder.GetCallBuilder()
}

func NewBaseController(cb callBuilderFactory) *baseController {
	baseController := baseController{callBuilder: cb}
	baseController.prepareRequest = baseController.callBuilder.GetCallBuilder()
	return &baseController
}