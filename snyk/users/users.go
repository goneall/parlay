// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	APITokenScopes   = "APIToken.Scopes"
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for PrincipalType.
const (
	PrincipalTypeAppInstance    PrincipalType = "app_instance"
	PrincipalTypeServiceAccount PrincipalType = "service_account"
	PrincipalTypeUser           PrincipalType = "user"
)

// Defines values for UserSettingsType.
const (
	UserSettingsTypeUserSettings UserSettingsType = "user_settings"
)

// ActualVersion Resolved API version
type ActualVersion = string

// AppInstance defines model for AppInstance.
type AppInstance struct {
	// DefaultOrgContext ID of the default org for the service account.
	DefaultOrgContext *openapi_types.UUID `json:"default_org_context,omitempty"`

	// Name The name of the service account.
	Name string `json:"name"`
}

// Error defines model for Error.
type Error struct {
	// Code An application-specific error code, expressed as a string value.
	Code *string `json:"code,omitempty"`

	// Detail A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail"`

	// Id A unique identifier for this particular occurrence of the problem.
	Id *openapi_types.UUID `json:"id,omitempty"`

	// Links A link that leads to further details about this particular occurrance of the problem.
	Links  *ErrorLink              `json:"links,omitempty"`
	Meta   *map[string]interface{} `json:"meta,omitempty"`
	Source *struct {
		// Parameter A string indicating which URI query parameter caused the error.
		Parameter *string `json:"parameter,omitempty"`

		// Pointer A JSON Pointer [RFC6901] to the associated entity in the request document.
		Pointer *string `json:"pointer,omitempty"`
	} `json:"source,omitempty"`

	// Status The HTTP status code applicable to this problem, expressed as a string value.
	Status string `json:"status"`

	// Title A short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
	Title *string `json:"title,omitempty"`
}

// ErrorDocument defines model for ErrorDocument.
type ErrorDocument struct {
	Errors  []Error `json:"errors"`
	Jsonapi JsonApi `json:"jsonapi"`
}

// ErrorLink A link that leads to further details about this particular occurrance of the problem.
type ErrorLink struct {
	About *LinkProperty `json:"about,omitempty"`
}

// JsonApi defines model for JsonApi.
type JsonApi struct {
	// Version Version of the JSON API specification this server supports.
	Version string `json:"version"`
}

// LinkProperty defines model for LinkProperty.
type LinkProperty struct {
	union json.RawMessage
}

// LinkProperty0 A string containing the link’s URL.
type LinkProperty0 = string

// LinkProperty1 defines model for .
type LinkProperty1 struct {
	// Href A string containing the link’s URL.
	Href string `json:"href"`

	// Meta Free-form object that may contain non-standard information.
	Meta *Meta `json:"meta,omitempty"`
}

// Links defines model for Links.
type Links struct {
	First   *LinkProperty `json:"first,omitempty"`
	Last    *LinkProperty `json:"last,omitempty"`
	Next    *LinkProperty `json:"next,omitempty"`
	Prev    *LinkProperty `json:"prev,omitempty"`
	Related *LinkProperty `json:"related,omitempty"`
	Self    *LinkProperty `json:"self,omitempty"`
}

// Meta Free-form object that may contain non-standard information.
type Meta map[string]interface{}

// Principal defines model for Principal.
type Principal struct {
	Attributes Principal_Attributes `json:"attributes"`

	// Id The Snyk ID corresponding to this user, service account or app
	Id openapi_types.UUID `json:"id"`

	// Type Content type.
	Type PrincipalType `json:"type"`
}

// PrincipalAttributes0 defines model for .
type PrincipalAttributes0 struct {
	// AvatarUrl The avatar url of the user.
	AvatarUrl string `json:"avatar_url"`

	// DefaultOrgContext ID of the default org for the user.
	DefaultOrgContext *openapi_types.UUID `json:"default_org_context,omitempty"`

	// Email The email of the user.
	Email string `json:"email"`

	// Name The name of the user.
	Name string `json:"name"`

	// Username The username of the user.
	Username *string `json:"username,omitempty"`
}

// Principal_Attributes defines model for Principal.Attributes.
type Principal_Attributes struct {
	union json.RawMessage
}

// PrincipalType Content type.
type PrincipalType string

// QueryVersion Requested API version
type QueryVersion = string

// ServiceAccount defines model for ServiceAccount.
type ServiceAccount struct {
	// DefaultOrgContext ID of the default org for the service account.
	DefaultOrgContext *openapi_types.UUID `json:"default_org_context,omitempty"`

	// Name The name of the service account.
	Name string `json:"name"`
}

// User defines model for User.
type User struct {
	// AvatarUrl The avatar url of the user.
	AvatarUrl string `json:"avatar_url"`

	// DefaultOrgContext ID of the default org for the user.
	DefaultOrgContext *openapi_types.UUID `json:"default_org_context,omitempty"`

	// Email The email of the user.
	Email string `json:"email"`

	// Name The name of the user.
	Name string `json:"name"`

	// Username The username of the user.
	Username *string `json:"username,omitempty"`
}

// UserPreferredOrgSettings defines model for UserPreferredOrgSettings.
type UserPreferredOrgSettings struct {
	// PreferredOrg The org to use for operations when there is no explicit org in context
	PreferredOrg *struct {
		// Id The id of the preferred org
		Id openapi_types.UUID `json:"id"`
	} `json:"preferred_org,omitempty"`
}

// UserSettings user settings
type UserSettings struct {
	Attributes *UserSettings_Attributes `json:"attributes,omitempty"`

	// Id The id of the user
	Id openapi_types.UUID `json:"id"`

	// Type The resource type
	Type UserSettingsType `json:"type"`
}

// UserSettings_Attributes defines model for UserSettings.Attributes.
type UserSettings_Attributes struct {
	union json.RawMessage
}

// UserSettingsType The resource type
type UserSettingsType string

// Version Requested API version
type Version = QueryVersion

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody struct {
	Data *struct {
		Attributes struct {
			Membership *struct {
				// Role Role name
				Role *string `json:"role,omitempty"`
			} `json:"membership"`
		} `json:"attributes"`

		// Id The Snyk ID corresponding to this user
		Id openapi_types.UUID `json:"id"`

		// Type Content type
		Type string `json:"type"`
	} `json:"data,omitempty"`
}

// UpdateUserParams defines parameters for UpdateUser.
type UpdateUserParams struct {
	// Version The requested version of the endpoint to process the request
	Version Version `form:"version" json:"version"`
}

// GetUserParams defines parameters for GetUser.
type GetUserParams struct {
	// Version The requested version of the endpoint to process the request
	Version Version `form:"version" json:"version"`
}

// GetSelfParams defines parameters for GetSelf.
type GetSelfParams struct {
	// Version The requested version of the endpoint to process the request
	Version Version `form:"version" json:"version"`
}

// GetUserSettingsParams defines parameters for GetUserSettings.
type GetUserSettingsParams struct {
	// Version The requested version of the endpoint to process the request
	Version Version `form:"version" json:"version"`
}

// UpdateUserJSONRequestBody defines body for UpdateUser for application/vnd.api+json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody

// AsLinkProperty0 returns the union data inside the LinkProperty as a LinkProperty0
func (t LinkProperty) AsLinkProperty0() (LinkProperty0, error) {
	var body LinkProperty0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLinkProperty0 overwrites any union data inside the LinkProperty as the provided LinkProperty0
func (t *LinkProperty) FromLinkProperty0(v LinkProperty0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLinkProperty0 performs a merge with any union data inside the LinkProperty, using the provided LinkProperty0
func (t *LinkProperty) MergeLinkProperty0(v LinkProperty0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsLinkProperty1 returns the union data inside the LinkProperty as a LinkProperty1
func (t LinkProperty) AsLinkProperty1() (LinkProperty1, error) {
	var body LinkProperty1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLinkProperty1 overwrites any union data inside the LinkProperty as the provided LinkProperty1
func (t *LinkProperty) FromLinkProperty1(v LinkProperty1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLinkProperty1 performs a merge with any union data inside the LinkProperty, using the provided LinkProperty1
func (t *LinkProperty) MergeLinkProperty1(v LinkProperty1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t LinkProperty) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *LinkProperty) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPrincipalAttributes0 returns the union data inside the Principal_Attributes as a PrincipalAttributes0
func (t Principal_Attributes) AsPrincipalAttributes0() (PrincipalAttributes0, error) {
	var body PrincipalAttributes0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPrincipalAttributes0 overwrites any union data inside the Principal_Attributes as the provided PrincipalAttributes0
func (t *Principal_Attributes) FromPrincipalAttributes0(v PrincipalAttributes0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePrincipalAttributes0 performs a merge with any union data inside the Principal_Attributes, using the provided PrincipalAttributes0
func (t *Principal_Attributes) MergePrincipalAttributes0(v PrincipalAttributes0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsServiceAccount returns the union data inside the Principal_Attributes as a ServiceAccount
func (t Principal_Attributes) AsServiceAccount() (ServiceAccount, error) {
	var body ServiceAccount
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromServiceAccount overwrites any union data inside the Principal_Attributes as the provided ServiceAccount
func (t *Principal_Attributes) FromServiceAccount(v ServiceAccount) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeServiceAccount performs a merge with any union data inside the Principal_Attributes, using the provided ServiceAccount
func (t *Principal_Attributes) MergeServiceAccount(v ServiceAccount) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsAppInstance returns the union data inside the Principal_Attributes as a AppInstance
func (t Principal_Attributes) AsAppInstance() (AppInstance, error) {
	var body AppInstance
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAppInstance overwrites any union data inside the Principal_Attributes as the provided AppInstance
func (t *Principal_Attributes) FromAppInstance(v AppInstance) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAppInstance performs a merge with any union data inside the Principal_Attributes, using the provided AppInstance
func (t *Principal_Attributes) MergeAppInstance(v AppInstance) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t Principal_Attributes) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Principal_Attributes) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsUserPreferredOrgSettings returns the union data inside the UserSettings_Attributes as a UserPreferredOrgSettings
func (t UserSettings_Attributes) AsUserPreferredOrgSettings() (UserPreferredOrgSettings, error) {
	var body UserPreferredOrgSettings
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromUserPreferredOrgSettings overwrites any union data inside the UserSettings_Attributes as the provided UserPreferredOrgSettings
func (t *UserSettings_Attributes) FromUserPreferredOrgSettings(v UserPreferredOrgSettings) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeUserPreferredOrgSettings performs a merge with any union data inside the UserSettings_Attributes, using the provided UserPreferredOrgSettings
func (t *UserSettings_Attributes) MergeUserPreferredOrgSettings(v UserPreferredOrgSettings) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t UserSettings_Attributes) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *UserSettings_Attributes) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

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
	// UpdateUser request with any body
	UpdateUserWithBody(ctx context.Context, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateUser(ctx context.Context, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUser request
	GetUser(ctx context.Context, orgId openapi_types.UUID, id openapi_types.UUID, params *GetUserParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSelf request
	GetSelf(ctx context.Context, params *GetSelfParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUserSettings request
	GetUserSettings(ctx context.Context, params *GetUserSettingsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) UpdateUserWithBody(ctx context.Context, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserRequestWithBody(c.Server, groupId, id, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateUser(ctx context.Context, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserRequest(c.Server, groupId, id, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUser(ctx context.Context, orgId openapi_types.UUID, id openapi_types.UUID, params *GetUserParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserRequest(c.Server, orgId, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetSelf(ctx context.Context, params *GetSelfParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSelfRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUserSettings(ctx context.Context, params *GetUserSettingsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserSettingsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewUpdateUserRequest calls the generic UpdateUser builder with application/vnd.api+json body
func NewUpdateUserRequest(server string, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, body UpdateUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateUserRequestWithBody(server, groupId, id, params, "application/vnd.api+json", bodyReader)
}

// NewUpdateUserRequestWithBody generates requests for UpdateUser with any type of body
func NewUpdateUserRequestWithBody(server string, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "group_id", runtime.ParamLocationPath, groupId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/groups/%s/users/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "version", runtime.ParamLocationQuery, params.Version); err != nil {
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

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetUserRequest generates requests for GetUser
func NewGetUserRequest(server string, orgId openapi_types.UUID, id openapi_types.UUID, params *GetUserParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "org_id", runtime.ParamLocationPath, orgId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/orgs/%s/users/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "version", runtime.ParamLocationQuery, params.Version); err != nil {
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

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSelfRequest generates requests for GetSelf
func NewGetSelfRequest(server string, params *GetSelfParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/self")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "version", runtime.ParamLocationQuery, params.Version); err != nil {
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

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetUserSettingsRequest generates requests for GetUserSettings
func NewGetUserSettingsRequest(server string, params *GetUserSettingsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/self/settings")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "version", runtime.ParamLocationQuery, params.Version); err != nil {
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
	// UpdateUser request with any body
	UpdateUserWithBodyWithResponse(ctx context.Context, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error)

	UpdateUserWithResponse(ctx context.Context, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error)

	// GetUser request
	GetUserWithResponse(ctx context.Context, orgId openapi_types.UUID, id openapi_types.UUID, params *GetUserParams, reqEditors ...RequestEditorFn) (*GetUserResponse, error)

	// GetSelf request
	GetSelfWithResponse(ctx context.Context, params *GetSelfParams, reqEditors ...RequestEditorFn) (*GetSelfResponse, error)

	// GetUserSettings request
	GetUserSettingsWithResponse(ctx context.Context, params *GetUserSettingsParams, reqEditors ...RequestEditorFn) (*GetUserSettingsResponse, error)
}

type UpdateUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r UpdateUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSelfResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetSelfResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSelfResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserSettingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetUserSettingsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserSettingsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// UpdateUserWithBodyWithResponse request with arbitrary body returning *UpdateUserResponse
func (c *ClientWithResponses) UpdateUserWithBodyWithResponse(ctx context.Context, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error) {
	rsp, err := c.UpdateUserWithBody(ctx, groupId, id, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserResponse(rsp)
}

func (c *ClientWithResponses) UpdateUserWithResponse(ctx context.Context, groupId openapi_types.UUID, id openapi_types.UUID, params *UpdateUserParams, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error) {
	rsp, err := c.UpdateUser(ctx, groupId, id, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserResponse(rsp)
}

// GetUserWithResponse request returning *GetUserResponse
func (c *ClientWithResponses) GetUserWithResponse(ctx context.Context, orgId openapi_types.UUID, id openapi_types.UUID, params *GetUserParams, reqEditors ...RequestEditorFn) (*GetUserResponse, error) {
	rsp, err := c.GetUser(ctx, orgId, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserResponse(rsp)
}

// GetSelfWithResponse request returning *GetSelfResponse
func (c *ClientWithResponses) GetSelfWithResponse(ctx context.Context, params *GetSelfParams, reqEditors ...RequestEditorFn) (*GetSelfResponse, error) {
	rsp, err := c.GetSelf(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSelfResponse(rsp)
}

// GetUserSettingsWithResponse request returning *GetUserSettingsResponse
func (c *ClientWithResponses) GetUserSettingsWithResponse(ctx context.Context, params *GetUserSettingsParams, reqEditors ...RequestEditorFn) (*GetUserSettingsResponse, error) {
	rsp, err := c.GetUserSettings(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserSettingsResponse(rsp)
}

// ParseUpdateUserResponse parses an HTTP response from a UpdateUserWithResponse call
func ParseUpdateUserResponse(rsp *http.Response) (*UpdateUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetUserResponse parses an HTTP response from a GetUserWithResponse call
func ParseGetUserResponse(rsp *http.Response) (*GetUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetSelfResponse parses an HTTP response from a GetSelfWithResponse call
func ParseGetSelfResponse(rsp *http.Response) (*GetSelfResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSelfResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetUserSettingsResponse parses an HTTP response from a GetUserSettingsWithResponse call
func ParseGetUserSettingsResponse(rsp *http.Response) (*GetUserSettingsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserSettingsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
