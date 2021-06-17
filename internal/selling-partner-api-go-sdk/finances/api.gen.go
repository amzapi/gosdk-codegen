// Package finances provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package finances

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListFinancialEventGroups request
	ListFinancialEventGroups(ctx context.Context, params *ListFinancialEventGroupsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListFinancialEventsByGroupId request
	ListFinancialEventsByGroupId(ctx context.Context, eventGroupId string, params *ListFinancialEventsByGroupIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListFinancialEvents request
	ListFinancialEvents(ctx context.Context, params *ListFinancialEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListFinancialEventsByOrderId request
	ListFinancialEventsByOrderId(ctx context.Context, orderId string, params *ListFinancialEventsByOrderIdParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListFinancialEventGroups(ctx context.Context, params *ListFinancialEventGroupsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListFinancialEventGroupsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListFinancialEventsByGroupId(ctx context.Context, eventGroupId string, params *ListFinancialEventsByGroupIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListFinancialEventsByGroupIdRequest(c.Server, eventGroupId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListFinancialEvents(ctx context.Context, params *ListFinancialEventsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListFinancialEventsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListFinancialEventsByOrderId(ctx context.Context, orderId string, params *ListFinancialEventsByOrderIdParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListFinancialEventsByOrderIdRequest(c.Server, orderId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListFinancialEventGroupsRequest generates requests for ListFinancialEventGroups
func NewListFinancialEventGroupsRequest(server string, params *ListFinancialEventGroupsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/finances/v0/financialEventGroups")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if params.MaxResultsPerPage != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "MaxResultsPerPage", runtime.ParamLocationQuery, *params.MaxResultsPerPage); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.FinancialEventGroupStartedBefore != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "FinancialEventGroupStartedBefore", runtime.ParamLocationQuery, *params.FinancialEventGroupStartedBefore); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.FinancialEventGroupStartedAfter != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "FinancialEventGroupStartedAfter", runtime.ParamLocationQuery, *params.FinancialEventGroupStartedAfter); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "NextToken", runtime.ParamLocationQuery, *params.NextToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListFinancialEventsByGroupIdRequest generates requests for ListFinancialEventsByGroupId
func NewListFinancialEventsByGroupIdRequest(server string, eventGroupId string, params *ListFinancialEventsByGroupIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "eventGroupId", runtime.ParamLocationPath, eventGroupId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/finances/v0/financialEventGroups/%s/financialEvents", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if params.MaxResultsPerPage != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "MaxResultsPerPage", runtime.ParamLocationQuery, *params.MaxResultsPerPage); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "NextToken", runtime.ParamLocationQuery, *params.NextToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListFinancialEventsRequest generates requests for ListFinancialEvents
func NewListFinancialEventsRequest(server string, params *ListFinancialEventsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/finances/v0/financialEvents")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if params.MaxResultsPerPage != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "MaxResultsPerPage", runtime.ParamLocationQuery, *params.MaxResultsPerPage); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PostedAfter != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "PostedAfter", runtime.ParamLocationQuery, *params.PostedAfter); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PostedBefore != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "PostedBefore", runtime.ParamLocationQuery, *params.PostedBefore); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "NextToken", runtime.ParamLocationQuery, *params.NextToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListFinancialEventsByOrderIdRequest generates requests for ListFinancialEventsByOrderId
func NewListFinancialEventsByOrderIdRequest(server string, orderId string, params *ListFinancialEventsByOrderIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "orderId", runtime.ParamLocationPath, orderId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/finances/v0/orders/%s/financialEvents", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if params.MaxResultsPerPage != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "MaxResultsPerPage", runtime.ParamLocationQuery, *params.MaxResultsPerPage); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "NextToken", runtime.ParamLocationQuery, *params.NextToken); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ListFinancialEventGroups request
	ListFinancialEventGroupsWithResponse(ctx context.Context, params *ListFinancialEventGroupsParams, reqEditors ...RequestEditorFn) (*ListFinancialEventGroupsResponse, error)

	// ListFinancialEventsByGroupId request
	ListFinancialEventsByGroupIdWithResponse(ctx context.Context, eventGroupId string, params *ListFinancialEventsByGroupIdParams, reqEditors ...RequestEditorFn) (*ListFinancialEventsByGroupIdResponse, error)

	// ListFinancialEvents request
	ListFinancialEventsWithResponse(ctx context.Context, params *ListFinancialEventsParams, reqEditors ...RequestEditorFn) (*ListFinancialEventsResponse, error)

	// ListFinancialEventsByOrderId request
	ListFinancialEventsByOrderIdWithResponse(ctx context.Context, orderId string, params *ListFinancialEventsByOrderIdParams, reqEditors ...RequestEditorFn) (*ListFinancialEventsByOrderIdResponse, error)
}

type ListFinancialEventGroupsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListFinancialEventGroupsResponse
	JSON400      *ListFinancialEventGroupsResponse
	JSON403      *ListFinancialEventGroupsResponse
	JSON404      *ListFinancialEventGroupsResponse
	JSON429      *ListFinancialEventGroupsResponse
	JSON500      *ListFinancialEventGroupsResponse
	JSON503      *ListFinancialEventGroupsResponse
}

// Status returns HTTPResponse.Status
func (r ListFinancialEventGroupsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFinancialEventGroupsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListFinancialEventsByGroupIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListFinancialEventsResponse
	JSON400      *ListFinancialEventsResponse
	JSON403      *ListFinancialEventsResponse
	JSON404      *ListFinancialEventsResponse
	JSON429      *ListFinancialEventsResponse
	JSON500      *ListFinancialEventsResponse
	JSON503      *ListFinancialEventsResponse
}

// Status returns HTTPResponse.Status
func (r ListFinancialEventsByGroupIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFinancialEventsByGroupIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListFinancialEventsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListFinancialEventsResponse
	JSON400      *ListFinancialEventsResponse
	JSON403      *ListFinancialEventsResponse
	JSON404      *ListFinancialEventsResponse
	JSON429      *ListFinancialEventsResponse
	JSON500      *ListFinancialEventsResponse
	JSON503      *ListFinancialEventsResponse
}

// Status returns HTTPResponse.Status
func (r ListFinancialEventsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFinancialEventsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListFinancialEventsByOrderIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListFinancialEventsResponse
	JSON400      *ListFinancialEventsResponse
	JSON403      *ListFinancialEventsResponse
	JSON404      *ListFinancialEventsResponse
	JSON429      *ListFinancialEventsResponse
	JSON500      *ListFinancialEventsResponse
	JSON503      *ListFinancialEventsResponse
}

// Status returns HTTPResponse.Status
func (r ListFinancialEventsByOrderIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFinancialEventsByOrderIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListFinancialEventGroupsWithResponse request returning *ListFinancialEventGroupsResponse
func (c *ClientWithResponses) ListFinancialEventGroupsWithResponse(ctx context.Context, params *ListFinancialEventGroupsParams, reqEditors ...RequestEditorFn) (*ListFinancialEventGroupsResponse, error) {
	rsp, err := c.ListFinancialEventGroups(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListFinancialEventGroupsResponse(rsp)
}

// ListFinancialEventsByGroupIdWithResponse request returning *ListFinancialEventsByGroupIdResponse
func (c *ClientWithResponses) ListFinancialEventsByGroupIdWithResponse(ctx context.Context, eventGroupId string, params *ListFinancialEventsByGroupIdParams, reqEditors ...RequestEditorFn) (*ListFinancialEventsByGroupIdResponse, error) {
	rsp, err := c.ListFinancialEventsByGroupId(ctx, eventGroupId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListFinancialEventsByGroupIdResponse(rsp)
}

// ListFinancialEventsWithResponse request returning *ListFinancialEventsResponse
func (c *ClientWithResponses) ListFinancialEventsWithResponse(ctx context.Context, params *ListFinancialEventsParams, reqEditors ...RequestEditorFn) (*ListFinancialEventsResponse, error) {
	rsp, err := c.ListFinancialEvents(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListFinancialEventsResponse(rsp)
}

// ListFinancialEventsByOrderIdWithResponse request returning *ListFinancialEventsByOrderIdResponse
func (c *ClientWithResponses) ListFinancialEventsByOrderIdWithResponse(ctx context.Context, orderId string, params *ListFinancialEventsByOrderIdParams, reqEditors ...RequestEditorFn) (*ListFinancialEventsByOrderIdResponse, error) {
	rsp, err := c.ListFinancialEventsByOrderId(ctx, orderId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListFinancialEventsByOrderIdResponse(rsp)
}

// ParseListFinancialEventGroupsResponse parses an HTTP response from a ListFinancialEventGroupsWithResponse call
func ParseListFinancialEventGroupsResponse(rsp *http.Response) (*ListFinancialEventGroupsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListFinancialEventGroupsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseListFinancialEventsByGroupIdResponse parses an HTTP response from a ListFinancialEventsByGroupIdWithResponse call
func ParseListFinancialEventsByGroupIdResponse(rsp *http.Response) (*ListFinancialEventsByGroupIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListFinancialEventsByGroupIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseListFinancialEventsResponse parses an HTTP response from a ListFinancialEventsWithResponse call
func ParseListFinancialEventsResponse(rsp *http.Response) (*ListFinancialEventsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListFinancialEventsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseListFinancialEventsByOrderIdResponse parses an HTTP response from a ListFinancialEventsByOrderIdWithResponse call
func ParseListFinancialEventsByOrderIdResponse(rsp *http.Response) (*ListFinancialEventsByOrderIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListFinancialEventsByOrderIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}