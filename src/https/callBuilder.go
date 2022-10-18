/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package https

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
    "net/url"
    "strings"
    "tester/src/exceptions"
    "tester/src/utilities"
)


type Authenticator func(bool) HttpInterceptor

const CONTENT_TYPE_HEADER = "content-type"
const ACCEPT_HEADER = "accept"
const CONTENT_LENGTH_HEADER = "content-length"
const AUTHORIZATION_HEADER = "authorization"
const FORM_URLENCODED_CONTENT_TYPE = "application/x-www-form-urlencoded"
const JSON_CONTENT_TYPE = "application/json"
const TEXT_CONTENT_TYPE = "text/plain; charset=utf-8"
const XML_CONTENT_TYPE = "application/xml"

type CallBuilderFactory func(httpMethod, path string) CallBuilder

type CallBuilder interface {
	AppendPath(path string)
	AppendTemplateParam(param string)
    AppendTemplateParams(params []string)
    BaseUrl(arg string)
	Method(httpMethodName string)
	Accept(acceptHeaderValue string)
	ContentType(contentTypeHeaderValue string)
	Header(name string, value interface{})
	CombineHeaders(headersToMerge map[string]string)
	QueryParam(name string, value interface{})
	QueryParams(parameters map[string]interface{})
	FormParam(name string, value interface{})
	Text(body string)
	FileStream(name, filepath string)
	Json(data interface{})
	intercept(interceptor HttpInterceptor)
	InterceptRequest(interceptor func(httpRequest *http.Request) *http.Request)
	toRequest() *http.Request
	Call() (*HttpContext, error)
	CallAsJson() (*json.Decoder, *http.Response, error)
	CallAsText() (string, *http.Response, error)
	CallAsStream() ([]byte, *http.Response, error)
	Authenticate(requiresAuth bool)
}

type defaultCallBuilder struct {
	path                   string
	baseUrlArg             string
	httpMethod             string
	acceptHeaderValue      string
	contentTypeHeaderValue string
	headers                map[string]string
	query                  url.Values
	form                   url.Values
	body                   string
	streamBody             bytes.Buffer
	httpClient             http.Client
	interceptors           []HttpInterceptor
	requiresAuth           bool
	authProvider           Authenticator
}

func newDefaultCallBuilder(
	httpClient http.Client,
	baseUrl,
	httpMethod,
	path string,
	authProvider Authenticator,
) *defaultCallBuilder {
	cb := defaultCallBuilder{
		httpClient:   httpClient,
		path:         path,
		baseUrlArg:   baseUrl,
		httpMethod:   httpMethod,
		authProvider: authProvider,
	}
	return &cb
}

func (cb *defaultCallBuilder) addAuthentication() {
	if cb.authProvider != nil {
		cb.intercept(cb.authProvider(cb.requiresAuth))
	}
}

func (cb *defaultCallBuilder) Authenticate(requiresAuth bool) {
	cb.requiresAuth = requiresAuth
	if cb.requiresAuth == true {
		cb.addAuthentication()
	}
}

func (cb *defaultCallBuilder) AppendPath(path string) {
	if cb.path != "" {
		cb.path = sanitizePath(mergePath(cb.path, path))
	} else {
		cb.path = sanitizePath(path)
	}
}

func (cb *defaultCallBuilder) AppendTemplateParam(param string) {
	if strings.Contains(cb.path, "%s") {
		cb.path = fmt.Sprintf(cb.path, "/"+param)
	} else {
		cb.AppendPath(param)
	}
}

func (cb *defaultCallBuilder) AppendTemplateParams(params []string) {
	for _, param := range params {
		cb.AppendTemplateParam(param)
	}
}

func (cb *defaultCallBuilder) BaseUrl(arg string) {
	cb.baseUrlArg = arg
}

func (cb *defaultCallBuilder) Method(httpMethodName string) {
	if strings.EqualFold(httpMethodName, http.MethodGet) {
		cb.httpMethod = http.MethodGet
	} else if strings.EqualFold(httpMethodName, http.MethodPut) {
		cb.httpMethod = http.MethodPut
	} else if strings.EqualFold(httpMethodName, http.MethodPost) {
		cb.httpMethod = http.MethodPost
	} else if strings.EqualFold(httpMethodName, http.MethodPatch) {
		cb.httpMethod = http.MethodPatch
	} else if strings.EqualFold(httpMethodName, http.MethodDelete) {
		cb.httpMethod = http.MethodDelete
	} else {
		log.Fatal("Invalid HTTP method given!")
	}
}

func (cb *defaultCallBuilder) Accept(acceptHeaderValue string) {
	cb.acceptHeaderValue = acceptHeaderValue
}

func (cb *defaultCallBuilder) ContentType(contentTypeHeaderValue string) {
	cb.contentTypeHeaderValue = contentTypeHeaderValue
}

func (cb *defaultCallBuilder) Header(
	name string,
	value interface{},
) {
	if cb.headers == nil {
		cb.headers = make(map[string]string)
	}
	SetHeaders(cb.headers, name, fmt.Sprintf("%v", value))
}

func (cb *defaultCallBuilder) CombineHeaders(headersToMerge map[string]string) {
	MergeHeaders(cb.headers, headersToMerge)
}

func (cb *defaultCallBuilder) QueryParam(
	name string,
	value interface{},
) {
	if cb.query == nil {
		cb.query = url.Values{}
	}
	cb.query = PrepareFormFields(name, value, cb.query)
}

func (cb *defaultCallBuilder) QueryParams(parameters map[string]interface{}) {
	cb.query = utilities.PrepareQueryParams(cb.query, parameters)
}

func (cb *defaultCallBuilder) FormParam(
	name string,
	value interface{},
) {
	if cb.form == nil {
		cb.form = url.Values{}
	}
	cb.form = PrepareFormFields(name, value, cb.form)
	cb.setContentTypeIfNotSet(FORM_URLENCODED_CONTENT_TYPE)
}

func (cb *defaultCallBuilder) Text(body string) {
	cb.body = body
	cb.setContentTypeIfNotSet(TEXT_CONTENT_TYPE)
}

func (cb *defaultCallBuilder) FileStream(name, filepath string) {
	resp, err := http.Get(filepath)
	if err != nil {
		log.Fatal(err)
	}
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	fw, err := writer.CreateFormFile(name, "image")
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.Copy(fw, resp.Body); err != nil {
		log.Fatal(err)
	}
	writer.Close()

	cb.streamBody = body
	cb.Header("content-type", writer.FormDataContentType())
}

func (cb *defaultCallBuilder) Json(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	cb.body = string(bytes)
	cb.setContentTypeIfNotSet(JSON_CONTENT_TYPE)
}

func (cb *defaultCallBuilder) setContentTypeIfNotSet(contentType string) {
	if cb.headers == nil {
		cb.headers = make(map[string]string)
	}
	cb.headers[CONTENT_TYPE_HEADER] = contentType
}

func (cb *defaultCallBuilder) intercept(interceptor HttpInterceptor) {
	cb.interceptors = append(cb.interceptors, interceptor)
}

func (cb *defaultCallBuilder) InterceptRequest(
	interceptor func(httpRequest *http.Request) *http.Request,
) {
	cb.intercept(
		func(
			req *http.Request,
			next HttpCallExecutor,
		) HttpContext {
			return next(interceptor(req))
		})
}

func (cb *defaultCallBuilder) toRequest() *http.Request {
	request := http.Request{
		Method: cb.httpMethod,
	}

	url, _ := url.Parse(mergePath(cb.baseUrlArg, cb.path))
	if len(cb.query) > 0 {
		url.RawQuery = cb.query.Encode()
	}
	url.RawQuery = strings.ReplaceAll(url.RawQuery, "+", "%20")
	request.URL = url

	request.Header = make(http.Header)
	if strings.TrimSpace(cb.acceptHeaderValue) != "" {
		request.Header.Add(ACCEPT_HEADER, cb.acceptHeaderValue)
	}

	if strings.TrimSpace(cb.contentTypeHeaderValue) != "" {
		request.Header.Add(CONTENT_TYPE_HEADER, cb.contentTypeHeaderValue)
	}

	for key, val := range cb.headers {
		if request.Header.Get(key) != "" {
			continue
		} else {
			request.Header.Add(key, val)
		}
	}

	if strings.TrimSpace(cb.body) != "" {
		request.Body = io.NopCloser(bytes.NewBuffer([]byte(cb.body)))
		defer request.Body.Close()
	} else if len(cb.form) > 0 {
		request.Form = cb.form
		replaced := strings.ReplaceAll(cb.form.Encode(), "+", "%20")
		request.Body = io.NopCloser(bytes.NewBuffer([]byte(replaced)))
	} else if cb.streamBody.Bytes() != nil {
		request.Body = io.NopCloser(&cb.streamBody)
	}

	return &request
}

func (cb *defaultCallBuilder) Call() (*HttpContext, error) {
	f := func(request *http.Request) HttpContext {
		client := http.DefaultClient
		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		return HttpContext{
			Request:  request,
			Response: response,
		}
	}

	pipeline := CallHttpInterceptors(cb.interceptors, f)
	context := pipeline(cb.toRequest())

	err := validateResponse(*context.Response)
    if err != nil {
		log.Fatal(err)
	}
	return &context, err
}

func (cb *defaultCallBuilder) CallAsJson() (*json.Decoder, *http.Response, error) {
	f := func(request *http.Request) *http.Request {
		request.Header.Set(ACCEPT_HEADER, JSON_CONTENT_TYPE)
		return request
	}

	cb.InterceptRequest(f)
	result, err := cb.Call()
	if result.Response.Body == http.NoBody {
		log.Fatal("Response body empty!")
	}

	return json.NewDecoder(result.Response.Body), result.Response, err
}

func (cb *defaultCallBuilder) CallAsText() (string, *http.Response, error) {
	result, err := cb.Call()
	if result.Response.Body == http.NoBody {
		log.Fatalln("Response body empty!")
	}

	body, err := ioutil.ReadAll(result.Response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body), result.Response, err
}

func (cb *defaultCallBuilder) CallAsStream() ([]byte, *http.Response, error) {
	result, err := cb.Call()
	if result.Response.Body == http.NoBody {
		log.Fatalln("Response body empty!")
	}

	bytes, err := ioutil.ReadAll(result.Response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return bytes, result.Response, err
}

func validateResponse(response http.Response) error {
	if response.StatusCode == 400 {
		return exceptions.NewApiError(400, "400 Global")
	}
	if response.StatusCode == 402 {
		return exceptions.NewApiError(402, "402 Global")
	}
	if response.StatusCode == 403 {
		return exceptions.NewApiError(403, "403 Global")
	}
	if response.StatusCode == 404 {
		return exceptions.NewApiError(404, "404 Global")
	}
	if response.StatusCode == 412 {
		return exceptions.NewNestedModelException(412, "Precondition Failed")
	}
	if response.StatusCode == 500 {
		return exceptions.NewGlobalTestException(500, "500 Global")
	}
	return nil
}


func mergePath(left, right string) string {
	if right == "" {
		return left
	}

	if left[len(left)-1] == '/' && right[0] == '/' {
		return left + strings.Replace(right, "/", "", -1)
	} else if left[len(left)-1] == '/' || right[0] == '/' {
		return left + right
	} else {
		return left + "/" + right
	}
}

func sanitizePath(path string) string {
	return strings.Replace(path, "//", "/", -1)
}

func CreateCallBuilderFactory(
	baseUrl string,
	auth Authenticator,
	httpClient http.Client,
) CallBuilderFactory {
	return func(
		httpMethod,
		path string,
	) CallBuilder {
		return newDefaultCallBuilder(
			httpClient,
			baseUrl,
			httpMethod,
			path,
			auth,
		)
	}
}

