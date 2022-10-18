/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package controllers

import (
	"io"
	"log"
	"net/http"
	"tester/src/data"
)

type TemplateParamsController struct {
	baseController
}

// TODO: type endpoint description here
func (t *TemplateParamsController) SendStringArray(strings []string) (
	data.EchoResponse,
	*http.Response,
	error) {
	req := t.prepareRequest("GET", "/%s")
	req.Authenticate(true)
	req.AppendTemplateParams(strings)
	decoder, resp, err := req.CallAsJson()

	var result data.EchoResponse
	for {
		if err := decoder.Decode(&result); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	return result, resp, err
}

// TODO: type endpoint description here
// func (t *TemplateParamsController) SendIntegerArray(integers []int64) (
//     data.EchoResponse,
//     *http.Response,
//     error) {
//     req := t.prepareRequest("GET", "/%s")
//     req.Authenticate(true)
//     req.AppendTemplateParams(integers)
//     decoder, resp, err := req.CallAsJson()

//     var result data.EchoResponse
//     for {
//     	if err := decoder.Decode(&result); err == io.EOF {
//     		break
//     	} else if err != nil {
//     		log.Fatal(err)
//     	}
//     }

//     return result, resp, err
// }
