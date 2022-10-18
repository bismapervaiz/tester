/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package testerClient

import (
    "fmt"
    "net/http"
    "tester/src/configuration"
    "tester/src/controllers"
    "tester/src/https"
)

type client struct {
    callBuilderFactory       https.CallBuilderFactory             
    config                   configuration.Configuration          
    ResponseTypesController  controllers.ResponseTypesController  
    FormParamsController     controllers.FormParamsController     
    BodyParamsController     controllers.BodyParamsController     
    ErrorCodesController     controllers.ErrorCodesController     
    QueryParamController     controllers.QueryParamController     
    EchoController           controllers.EchoController           
    HeaderController         controllers.HeaderController         
    TemplateParamsController controllers.TemplateParamsController 
}

// Constructor for client.
func NewClient(
    server configuration.Server,
    config configuration.Configuration) (
    *client) {
    client := &client{config: config}
    
    client.callBuilderFactory = callBuilderHandler(
    	getBaseUri(server, client.config),
    	nil,
    	*http.DefaultClient,
    )
    
    baseController := controllers.NewBaseController(client)
    client.ResponseTypesController = *controllers.NewResponseTypesControllerInterface(*baseController)
    client.FormParamsController = *controllers.NewFormParamsControllerInterface(*baseController)
    client.BodyParamsController = *controllers.NewBodyParamsControllerInterface(*baseController)
    client.ErrorCodesController = *controllers.NewErrorCodesControllerInterface(*baseController)
    client.QueryParamController = *controllers.NewQueryParamControllerInterface(*baseController)
    client.EchoController = *controllers.NewEchoControllerInterface(*baseController)
    client.HeaderController = *controllers.NewHeaderControllerInterface(*baseController)
    client.TemplateParamsController = *controllers.NewTemplateParamsControllerInterface(*baseController)
    return client
}

func (c *client) GetCallBuilder() (
    https.CallBuilderFactory) {
    return c.callBuilderFactory
}

func callBuilderHandler(
    baseUrl string,
    auth https.Authenticator,
    httpClient http.Client) (
    https.CallBuilderFactory) {
    return https.CreateCallBuilderFactory(baseUrl, auth, httpClient)
}

func getBaseUri(
    server configuration.Server,
    config configuration.Configuration) (
    string) {
    if config.Environment == configuration.Environment(configuration.PRODUCTION) {
        if server == configuration.Server(configuration.ENUMDEFAULT) {
            return fmt.Sprintf("http://apimatic.hopto.org:%s", config.Suites)
        }
        if server == configuration.Server(configuration.AUTHSERVER) {
            return fmt.Sprintf("http://apimaticauth.hopto.org:3000", )
        }
    }
    if config.Environment == configuration.Environment(configuration.TESTING) {
        if server == configuration.Server(configuration.ENUMDEFAULT) {
            return fmt.Sprintf("http://localhost:3000", )
        }
        if server == configuration.Server(configuration.AUTHSERVER) {
            return fmt.Sprintf("http://apimaticauth.xhopto.org:3000", )
        }
    }
    return "TODO: Select a valid server."
}

