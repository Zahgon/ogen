package ogen

import (
	"encoding/json"
)

// NewSpec returns a new Spec.
func NewSpec() *Spec {
	_ = "STUB: not implemented"

	// SetOpenAPI sets the OpenAPI Specification version of the document.
	return nil
}

func (s *Spec) SetOpenAPI(v string) *Spec { _ = "STUB: not implemented"; return nil }

// SetInfo sets the Info of the Spec.
func (s *Spec) SetInfo(i *Info) *Spec { _ = "STUB: not implemented"; return nil }

// SetServers sets the Servers of the Spec.
func (s *Spec) SetServers(srvs []Server) *Spec { _ = "STUB: not implemented"; return nil }

// AddServers adds Servers to the Servers of the Spec.
func (s *Spec) AddServers(srvs ...*Server) *Spec { _ = "STUB: not implemented"; return nil }

// SetPaths sets the Paths of the Spec.
func (s *Spec) SetPaths(p Paths) *Spec { _ = "STUB: not implemented"; return nil }

// AddPathItem adds the given PathItem under the given Name to the Paths of the Spec.
func (s *Spec) AddPathItem(n string, p *PathItem) *Spec { _ = "STUB: not implemented"; return nil }

// AddNamedPathItems adds the given namedPaths to the Paths of the Spec.
func (s *Spec) AddNamedPathItems(ps ...*NamedPathItem) *Spec { _ = "STUB: not implemented"; return nil }

// SetComponents sets the Components of the Spec.
func (s *Spec) SetComponents(c *Components) *Spec { _ = "STUB: not implemented"; return nil }

// initPaths ensures the Paths map is allocated.
func (s *Spec) initPaths() { _ = "STUB: not implemented"; return }

// AddSchema adds the given Schema under the given Name to the Components of the Spec.
func (s *Spec) AddSchema(n string, sc *Schema) *Spec { _ = "STUB: not implemented"; return nil }

// AddNamedSchemas adds the given namedSchemas to the Components of the Spec.
func (s *Spec) AddNamedSchemas(scs ...*NamedSchema) *Spec { _ = "STUB: not implemented"; return nil }

// AddResponse adds the given Response under the given Name to the Components of the Spec.
func (s *Spec) AddResponse(n string, sc *Response) *Spec { _ = "STUB: not implemented"; return nil }

// AddNamedResponses adds the given namedResponses to the Components of the Spec.
func (s *Spec) AddNamedResponses(scs ...*NamedResponse) *Spec {
	_ = "STUB: not implemented"
	return nil
}

// AddParameter adds the given Parameter under the given Name to the Components of the Spec.
func (s *Spec) AddParameter(n string, p *Parameter) *Spec { _ = "STUB: not implemented"; return nil }

// AddNamedParameters adds the given namedParameters to the Components of the Spec.
func (s *Spec) AddNamedParameters(ps ...*NamedParameter) *Spec {
	_ = "STUB: not implemented"
	return nil
}

// AddRequestBody adds the given RequestBody under the given Name to the Components of the Spec.
func (s *Spec) AddRequestBody(n string, sc *RequestBody) *Spec {
	_ = "STUB: not implemented"
	return nil
}

// AddNamedRequestBodies adds the given namedRequestBodies to the Components of the Spec.
func (s *Spec) AddNamedRequestBodies(scs ...*NamedRequestBody) *Spec {
	_ = "STUB: not implemented"
	return nil
}

// RefSchema returns a new Schema referencing the given name.
func (s *Spec) RefSchema(n string) *NamedSchema { _ = "STUB: not implemented"; return nil }

// RefResponse returns a new Response referencing the given name.
func (s *Spec) RefResponse(n string) *NamedResponse { _ = "STUB: not implemented"; return nil }

// RefRequestBody returns a new RequestBody referencing the given name.
func (s *Spec) RefRequestBody(n string) *NamedRequestBody { _ = "STUB: not implemented"; return nil }

// initComponents ensures the Components property is non-nil.
func (s *Spec) initComponents() { _ = "STUB: not implemented"; return }

// initParameters ensures the Parameters map is allocated.
func (s *Spec) initParameters() { _ = "STUB: not implemented"; return }

// initSchemas ensures the Schemas map is allocated.
func (s *Spec) initSchemas() { _ = "STUB: not implemented"; return }

// initResponses ensures the Responses map is allocated.
func (s *Spec) initResponses() { _ = "STUB: not implemented"; return }

// initRequestBodies ensures the RequestBodies map is allocated.
func (s *Spec) initRequestBodies() { _ = "STUB: not implemented"; return }

// NewRequestBody returns a new RequestBody.
func NewRequestBody() *RequestBody { _ = "STUB: not implemented"; return nil }

// SetRef sets the Ref of the RequestBody.
func (r *RequestBody) SetRef(ref string) *RequestBody { _ = "STUB: not implemented"; return nil }

// SetDescription sets the Description of the RequestBody.
func (r *RequestBody) SetDescription(d string) *RequestBody { _ = "STUB: not implemented"; return nil }

// SetContent sets the Content of the RequestBody.
func (r *RequestBody) SetContent(c map[string]Media) *RequestBody {
	_ = "STUB: not implemented"
	return nil

	// AddContent adds the given Schema under the MediaType to the Content of the Response.
}

func (r *RequestBody) AddContent(mt string, s *Schema) *RequestBody {
	_ = "STUB: not implemented"
	return nil
}

// SetJSONContent sets the given Schema under the JSON MediaType to the Content of the Response.
func (r *RequestBody) SetJSONContent(s *Schema) *RequestBody { _ = "STUB: not implemented"; return nil }

// initContent ensures the Content map is allocated.
func (r *RequestBody) initContent() { _ = "STUB: not implemented"; return }

// SetRequired sets the Required of the RequestBody.
func (r *RequestBody) SetRequired(req bool) *RequestBody { _ = "STUB: not implemented"; return nil }

// ToNamed returns a NamedRequestBody wrapping the receiver.
func (r *RequestBody) ToNamed(n string) *NamedRequestBody { _ = "STUB: not implemented"; return nil }

// NamedRequestBody can be used to construct a reference to the wrapped RequestBody.
type NamedRequestBody struct {
	RequestBody *RequestBody
	Name        string
}

// NewNamedRequestBody returns a new NamedRequestBody.
func NewNamedRequestBody(n string, p *RequestBody) *NamedRequestBody {
	_ = "STUB: not implemented"
	return nil
}

// AsLocalRef returns a new RequestBody referencing the wrapped RequestBody in the local document.
func (p *NamedRequestBody) AsLocalRef() *RequestBody { _ = "STUB: not implemented"; return nil }

// NewInfo returns a new Info.
func NewInfo() *Info {
	_ = "STUB: not implemented"

	// SetTitle sets the title of the Info.
	return nil
}

func (i *Info) SetTitle(t string) *Info { _ = "STUB: not implemented"; return nil }

// SetDescription sets the description of the Info.
func (i *Info) SetDescription(d string) *Info { _ = "STUB: not implemented"; return nil }

// SetTermsOfService sets the terms of service of the Info.
func (i *Info) SetTermsOfService(t string) *Info { _ = "STUB: not implemented"; return nil }

// SetContact sets the Contact of the Info.
func (i *Info) SetContact(c *Contact) *Info { _ = "STUB: not implemented"; return nil }

// SetLicense sets the License of the Info.
func (i *Info) SetLicense(l *License) *Info { _ = "STUB: not implemented"; return nil }

// SetVersion sets the version of the Info.
func (i *Info) SetVersion(v string) *Info { _ = "STUB: not implemented"; return nil }

// NewContact returns a new Contact.
func NewContact() *Contact { _ = "STUB: not implemented"; return nil }

// SetName sets the Name of the Contact.
func (c *Contact) SetName(n string) *Contact { _ = "STUB: not implemented"; return nil }

// SetURL sets the URL of the Contact.
func (c *Contact) SetURL(url string) *Contact { _ = "STUB: not implemented"; return nil }

// SetEmail sets the Email of the Contact.
func (c *Contact) SetEmail(e string) *Contact { _ = "STUB: not implemented"; return nil }

// NewLicense returns a new License.
func NewLicense() *License { _ = "STUB: not implemented"; return nil }

// SetName sets the Name of the License.
func (l *License) SetName(n string) *License { _ = "STUB: not implemented"; return nil }

// SetURL sets the URL of the License.
func (l *License) SetURL(url string) *License { _ = "STUB: not implemented"; return nil }

// NewServer returns a new Server.
func NewServer() *Server {
	_ = "STUB: not implemented"

	// SetDescription sets the Description of the Server.
	return nil
}

func (s *Server) SetDescription(d string) *Server { _ = "STUB: not implemented"; return nil }

// SetURL sets the URL of the Server.
func (s *Server) SetURL(url string) *Server { _ = "STUB: not implemented"; return nil }

// NewPathItem returns a new PathItem.
func NewPathItem() *PathItem { _ = "STUB: not implemented"; return nil }

// SetRef sets the Ref of the PathItem.
func (p *PathItem) SetRef(r string) *PathItem { _ = "STUB: not implemented"; return nil }

// SetDescription sets the Description of the PathItem.
func (p *PathItem) SetDescription(d string) *PathItem { _ = "STUB: not implemented"; return nil }

// SetGet sets the Get of the PathItem.
func (p *PathItem) SetGet(o *Operation) *PathItem { _ = "STUB: not implemented"; return nil }

// SetPut sets the Put of the PathItem.
func (p *PathItem) SetPut(o *Operation) *PathItem { _ = "STUB: not implemented"; return nil }

// SetPost sets the Post of the PathItem.
func (p *PathItem) SetPost(o *Operation) *PathItem { _ = "STUB: not implemented"; return nil }

// SetDelete sets the Delete of the PathItem.
func (p *PathItem) SetDelete(o *Operation) *PathItem { _ = "STUB: not implemented"; return nil }

// SetOptions sets the Options of the PathItem.
func (p *PathItem) SetOptions(o *Operation) *PathItem { _ = "STUB: not implemented"; return nil }

// SetHead sets the Head of the PathItem.
func (p *PathItem) SetHead(o *Operation) *PathItem { _ = "STUB: not implemented"; return nil }

// SetPatch sets the Patch of the PathItem.
func (p *PathItem) SetPatch(o *Operation) *PathItem { _ = "STUB: not implemented"; return nil }

// SetTrace sets the Trace of the PathItem.
func (p *PathItem) SetTrace(o *Operation) *PathItem { _ = "STUB: not implemented"; return nil }

// SetQuery sets the Query of the PathItem.
func (p *PathItem) SetQuery(o *Operation) *PathItem { _ = "STUB: not implemented"; return nil }

// SetAdditionalOperations sets the AdditionalOperations of the PathItem.
func (p *PathItem) SetAdditionalOperations(ops map[string]*Operation) *PathItem {
	_ = "STUB: not implemented"
	return nil
}

// SetAdditionalOperation sets a single operation in the AdditionalOperations of the PathItem.
func (p *PathItem) SetAdditionalOperation(method string, o *Operation) *PathItem {
	_ = "STUB: not implemented"
	return nil
}

func (p *PathItem) initAdditionalOperations() { _ = "STUB: not implemented"; return }

// SetServers sets the Servers of the PathItem.
func (p *PathItem) SetServers(srvs []Server) *PathItem { _ = "STUB: not implemented"; return nil }

// AddServers adds Servers to the Servers of the PathItem.
func (p *PathItem) AddServers(srvs ...*Server) *PathItem { _ = "STUB: not implemented"; return nil }

// SetParameters sets the Parameters of the PathItem.
func (p *PathItem) SetParameters(ps []*Parameter) *PathItem { _ = "STUB: not implemented"; return nil }

// AddParameters adds Parameters to the Parameters of the PathItem.
func (p *PathItem) AddParameters(ps ...*Parameter) *PathItem { _ = "STUB: not implemented"; return nil }

// ToNamed returns a NamedPathItem wrapping the receiver.
func (p *PathItem) ToNamed(n string) *NamedPathItem { _ = "STUB: not implemented"; return nil }

// NamedPathItem can be used to construct a reference to the wrapped PathItem.
type NamedPathItem struct {
	PathItem *PathItem
	Name     string
}

// NewNamedPathItem returns a new NamedPathItem.
func NewNamedPathItem(n string, p *PathItem) *NamedPathItem { _ = "STUB: not implemented"; return nil }

// AsLocalRef returns a new PathItem referencing the wrapped PathItem in the local document.
func (p *NamedPathItem) AsLocalRef() *PathItem { _ = "STUB: not implemented"; return nil }

// NewOperation returns a new Operation.
func NewOperation() *Operation { _ = "STUB: not implemented"; return nil }

// SetTags sets the Tags of the Operation.
func (o *Operation) SetTags(ts []string) *Operation { _ = "STUB: not implemented"; return nil }

// AddTags adds Tags to the Tags of the Operation.
func (o *Operation) AddTags(ts ...string) *Operation { _ = "STUB: not implemented"; return nil }

// SetSummary sets the Summary of the Operation.
func (o *Operation) SetSummary(s string) *Operation { _ = "STUB: not implemented"; return nil }

// SetDescription sets the Description of the Operation.
func (o *Operation) SetDescription(d string) *Operation { _ = "STUB: not implemented"; return nil }

// SetOperationID sets the OperationID of the Operation.
func (o *Operation) SetOperationID(id string) *Operation { _ = "STUB: not implemented"; return nil }

// SetParameters sets the Parameters of the Operation.
func (o *Operation) SetParameters(ps []*Parameter) *Operation {
	_ = "STUB: not implemented"
	return nil
}

// AddParameters adds Parameters to the Parameters of the Operation.
func (o *Operation) AddParameters(ps ...*Parameter) *Operation {
	_ = "STUB: not implemented"
	return nil
}

// SetRequestBody sets the RequestBody of the Operation.
func (o *Operation) SetRequestBody(r *RequestBody) *Operation {
	_ = "STUB: not implemented"
	return nil
}

// SetResponses sets the Responses of the Operation.
func (o *Operation) SetResponses(r Responses) *Operation { _ = "STUB: not implemented"; return nil }

// AddResponse adds the given Response under the given Name to the Responses of the Operation.
func (o *Operation) AddResponse(n string, p *Response) *Operation {
	_ = "STUB: not implemented"
	return nil
}

// AddNamedResponses adds the given namedResponses to the Responses of the Operation.
func (o *Operation) AddNamedResponses(ps ...*NamedResponse) *Operation {
	_ = "STUB: not implemented"
	return nil
}

// initResponses ensures the Responses map is allocated.
func (o *Operation) initResponses() { _ = "STUB: not implemented"; return }

// NewParameter returns a new Parameter.
func NewParameter() *Parameter { _ = "STUB: not implemented"; return nil }

// SetRef sets the Ref of the Parameter.
func (p *Parameter) SetRef(r string) *Parameter { _ = "STUB: not implemented"; return nil }

// SetName sets the Name of the Parameter.
func (p *Parameter) SetName(n string) *Parameter { _ = "STUB: not implemented"; return nil }

// SetIn sets the In of the Parameter.
func (p *Parameter) SetIn(i string) *Parameter {
	_ = "STUB: not implemented"

	// InPath sets the In of the Parameter to "path".
	return nil
}

func (p *Parameter) InPath() *Parameter { _ = "STUB: not implemented"; return nil }

// InQuery sets the In of the Parameter to "query".
func (p *Parameter) InQuery() *Parameter { _ = "STUB: not implemented"; return nil }

// InHeader sets the In of the Parameter to "header".
func (p *Parameter) InHeader() *Parameter { _ = "STUB: not implemented"; return nil }

// InCookie sets the In of the Parameter to "cookie".
func (p *Parameter) InCookie() *Parameter { _ = "STUB: not implemented"; return nil }

// SetDescription sets the Description of the Parameter.
func (p *Parameter) SetDescription(d string) *Parameter { _ = "STUB: not implemented"; return nil }

// SetSchema sets the Schema of the Parameter.
func (p *Parameter) SetSchema(s *Schema) *Parameter { _ = "STUB: not implemented"; return nil }

// SetRequired sets the Required of the Parameter.
func (p *Parameter) SetRequired(r bool) *Parameter { _ = "STUB: not implemented"; return nil }

// SetDeprecated sets the Deprecated of the Parameter.
func (p *Parameter) SetDeprecated(d bool) *Parameter { _ = "STUB: not implemented"; return nil }

// SetContent sets the Content of the Parameter.
func (p *Parameter) SetContent(c map[string]Media) *Parameter {
	_ = "STUB: not implemented"
	return nil

	// TODO(masseelch): Add Content helpers for Parameter
}

// SetStyle sets the Style of the Parameter.
func (p *Parameter) SetStyle(s string) *Parameter { _ = "STUB: not implemented"; return nil }

// SetExplode sets the Explode of the Parameter.
func (p *Parameter) SetExplode(e bool) *Parameter { _ = "STUB: not implemented"; return nil }

// ToNamed returns a NamedParameter wrapping the receiver.
func (p *Parameter) ToNamed(n string) *NamedParameter { _ = "STUB: not implemented"; return nil }

// NamedParameter can be used to construct a reference to the wrapped Parameter.
type NamedParameter struct {
	Parameter *Parameter
	Name      string
}

// NewNamedParameter returns a new NamedParameter.
func NewNamedParameter(n string, p *Parameter) *NamedParameter {
	_ = "STUB: not implemented"
	return nil
}

// AsLocalRef returns a new Parameter referencing the wrapped Parameter in the local document.
func (p *NamedParameter) AsLocalRef() *Parameter { _ = "STUB: not implemented"; return nil }

// NewResponse returns a new Response.
func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

// SetRef sets the Ref of the Response.
func (r *Response) SetRef(ref string) *Response { _ = "STUB: not implemented"; return nil }

// SetDescription sets the Description of the Response.
func (r *Response) SetDescription(d string) *Response { _ = "STUB: not implemented"; return nil }

// SetHeaders sets the Headers of the Response.
func (r *Response) SetHeaders(h map[string]*Header) *Response {
	_ = "STUB: not implemented"
	return nil

	// SetContent sets the Content of the Response.
}

func (r *Response) SetContent(c map[string]Media) *Response { _ = "STUB: not implemented"; return nil }

// AddContent adds the given Schema under the MediaType to the Content of the Response.
func (r *Response) AddContent(mt string, s *Schema) *Response {
	_ = "STUB: not implemented"
	return nil
}

// SetJSONContent sets the given Schema under the JSON MediaType to the Content of the Response.
func (r *Response) SetJSONContent(s *Schema) *Response { _ = "STUB: not implemented"; return nil }

// initContent ensures the Content map is allocated.
func (r *Response) initContent() { _ = "STUB: not implemented"; return }

// SetLinks sets the Links of the Response.
func (r *Response) SetLinks(l map[string]*Link) *Response { _ = "STUB: not implemented"; return nil }

// ToNamed returns a NamedResponse wrapping the receiver.
func (r *Response) ToNamed(n string) *NamedResponse { _ = "STUB: not implemented"; return nil }

// NamedResponse can be used to construct a reference to the wrapped Response.
type NamedResponse struct {
	Response *Response
	Name     string
}

// NewNamedResponse returns a new NamedResponse.
func NewNamedResponse(n string, p *Response) *NamedResponse { _ = "STUB: not implemented"; return nil }

// AsLocalRef returns a new Response referencing the wrapped Response in the local document.
func (p *NamedResponse) AsLocalRef() *Response { _ = "STUB: not implemented"; return nil }

// TODO(masseelch): Discriminator

// NewSchema returns a new Schema.
func NewSchema() *Schema {
	_ = "STUB: not implemented"

	// SetRef sets the Ref of the Schema.
	return nil
}

func (s *Schema) SetRef(r string) *Schema { _ = "STUB: not implemented"; return nil }

// SetSummary sets the Summary of the Schema.
func (s *Schema) SetSummary(smry string) *Schema { _ = "STUB: not implemented"; return nil }

// SetDescription sets the Description of the Schema.
func (s *Schema) SetDescription(d string) *Schema { _ = "STUB: not implemented"; return nil }

// SetType sets the Type of the Schema.
func (s *Schema) SetType(t string) *Schema { _ = "STUB: not implemented"; return nil }

// SetFormat sets the Format of the Schema.
func (s *Schema) SetFormat(f string) *Schema { _ = "STUB: not implemented"; return nil }

// SetProperties sets the Properties of the Schema.
func (s *Schema) SetProperties(p *Properties) *Schema { _ = "STUB: not implemented"; return nil }

// AddOptionalProperties adds the Properties to the Properties of the Schema.
func (s *Schema) AddOptionalProperties(ps ...*Property) *Schema {
	_ = "STUB: not implemented"
	return nil
}

// AddRequiredProperties adds the Properties to the Properties of the Schema and marks them as required.
func (s *Schema) AddRequiredProperties(ps ...*Property) *Schema {
	_ = "STUB: not implemented"
	return nil
}

// SetRequired sets the Required of the Schema.
func (s *Schema) SetRequired(r []string) *Schema { _ = "STUB: not implemented"; return nil }

// SetItems sets the Items of the Schema.
func (s *Schema) SetItems(i *Schema) *Schema { _ = "STUB: not implemented"; return nil }

// SetNullable sets the Nullable of the Schema.
func (s *Schema) SetNullable(n bool) *Schema { _ = "STUB: not implemented"; return nil }

// SetAllOf sets the AllOf of the Schema.
func (s *Schema) SetAllOf(a []*Schema) *Schema { _ = "STUB: not implemented"; return nil }

// SetOneOf sets the OneOf of the Schema.
func (s *Schema) SetOneOf(o []*Schema) *Schema { _ = "STUB: not implemented"; return nil }

// SetAnyOf sets the AnyOf of the Schema.
func (s *Schema) SetAnyOf(a []*Schema) *Schema { _ = "STUB: not implemented"; return nil }

// SetDiscriminator sets the Discriminator of the Schema.
func (s *Schema) SetDiscriminator(d *Discriminator) *Schema { _ = "STUB: not implemented"; return nil }

// SetEnum sets the Enum of the Schema.
func (s *Schema) SetEnum(e []json.RawMessage) *Schema { _ = "STUB: not implemented"; return nil }

// SetMultipleOf sets the MultipleOf of the Schema.
func (s *Schema) SetMultipleOf(m *uint64) *Schema { _ = "STUB: not implemented"; return nil }

// SetMaximum sets the Maximum of the Schema.
func (s *Schema) SetMaximum(m *int64) *Schema { _ = "STUB: not implemented"; return nil }

// SetExclusiveMaximum sets the ExclusiveMaximum of the Schema.
func (s *Schema) SetExclusiveMaximum(e bool) *Schema { _ = "STUB: not implemented"; return nil }

// SetMinimum sets the Minimum of the Schema.
func (s *Schema) SetMinimum(m *int64) *Schema { _ = "STUB: not implemented"; return nil }

// SetExclusiveMinimum sets the ExclusiveMinimum of the Schema.
func (s *Schema) SetExclusiveMinimum(e bool) *Schema { _ = "STUB: not implemented"; return nil }

// SetMaxLength sets the MaxLength of the Schema.
func (s *Schema) SetMaxLength(m *uint64) *Schema { _ = "STUB: not implemented"; return nil }

// SetMinLength sets the MinLength of the Schema.
func (s *Schema) SetMinLength(m *uint64) *Schema { _ = "STUB: not implemented"; return nil }

// SetPattern sets the Pattern of the Schema.
func (s *Schema) SetPattern(p string) *Schema { _ = "STUB: not implemented"; return nil }

// SetMaxItems sets the MaxItems of the Schema.
func (s *Schema) SetMaxItems(m *uint64) *Schema { _ = "STUB: not implemented"; return nil }

// SetMinItems sets the MinItems of the Schema.
func (s *Schema) SetMinItems(m *uint64) *Schema { _ = "STUB: not implemented"; return nil }

// SetUniqueItems sets the UniqueItems of the Schema.
func (s *Schema) SetUniqueItems(u bool) *Schema { _ = "STUB: not implemented"; return nil }

// SetMaxProperties sets the MaxProperties of the Schema.
func (s *Schema) SetMaxProperties(m *uint64) *Schema { _ = "STUB: not implemented"; return nil }

// SetMinProperties sets the MinProperties of the Schema.
func (s *Schema) SetMinProperties(m *uint64) *Schema { _ = "STUB: not implemented"; return nil }

// SetDefault sets the Default of the Schema.
func (s *Schema) SetDefault(d json.RawMessage) *Schema { _ = "STUB: not implemented"; return nil }

// SetDeprecated sets the Deprecated of the Schema.
func (s *Schema) SetDeprecated(d bool) *Schema { _ = "STUB: not implemented"; return nil }

// ToNamed returns a NamedSchema wrapping the receiver.
func (s *Schema) ToNamed(n string) *NamedSchema { _ = "STUB: not implemented"; return nil }

// Int returns an integer OAS data type (Schema).
func Int() *Schema { _ = "STUB: not implemented"; return nil }

// Int32 returns an 32-bit integer OAS data type (Schema).
func Int32() *Schema { _ = "STUB: not implemented"; return nil }

// Int64 returns an 64-bit integer OAS data type (Schema).
func Int64() *Schema { _ = "STUB: not implemented"; return nil }

// Float returns a float OAS data type (Schema).
func Float() *Schema { _ = "STUB: not implemented"; return nil }

// Double returns a double OAS data type (Schema).
func Double() *Schema { _ = "STUB: not implemented"; return nil }

// String returns a string OAS data type (Schema).
func String() *Schema { _ = "STUB: not implemented"; return nil }

// UUID returns a UUID OAS data type (Schema).
func UUID() *Schema { _ = "STUB: not implemented"; return nil }

// Bytes returns a base64 encoded OAS data type (Schema).
func Bytes() *Schema { _ = "STUB: not implemented"; return nil }

// Binary returns a sequence of octets OAS data type (Schema).
func Binary() *Schema { _ = "STUB: not implemented"; return nil }

// Bool returns a boolean OAS data type (Schema).
func Bool() *Schema { _ = "STUB: not implemented"; return nil }

// Date returns a date as defined by full-date - RFC3339 OAS data type (Schema).
func Date() *Schema { _ = "STUB: not implemented"; return nil }

// DateTime returns a date as defined by date-time - RFC3339 OAS data type (Schema).
func DateTime() *Schema { _ = "STUB: not implemented"; return nil }

// HTTPDate returns a date as defined by HTTP-date - RFC7231 OAS data type (Schema).
func HTTPDate() *Schema { _ = "STUB: not implemented"; return nil }

// Password returns an obscured OAS data type (Schema).
func Password() *Schema { _ = "STUB: not implemented"; return nil }

// schema returns a Schema for a primitive type.
func schema(t, f string) *Schema { _ = "STUB: not implemented"; return nil }

// AsArray returns a new "array" Schema wrapping the receiver.
func (s *Schema) AsArray() *Schema { _ = "STUB: not implemented"; return nil }

// AsEnum returns a new "enum" Schema wrapping the receiver.
func (s *Schema) AsEnum(def json.RawMessage, values ...json.RawMessage) *Schema {
	_ = "STUB: not implemented"
	return nil
}

// ToProperty returns a Property with the given name and with this Schema.
func (s *Schema) ToProperty(n string) *Property { _ = "STUB: not implemented"; return nil }

// NamedSchema can be used to construct a reference to the wrapped Schema.
type NamedSchema struct {
	Schema *Schema
	Name   string
}

// NewNamedSchema returns a new NamedSchema.
func NewNamedSchema(n string, p *Schema) *NamedSchema { _ = "STUB: not implemented"; return nil }

// AsLocalRef returns a new Schema referencing the wrapped Schema in the local document.
func (p *NamedSchema) AsLocalRef() *Schema { _ = "STUB: not implemented"; return nil }

// NewProperty returns a new Property.
func NewProperty() *Property { _ = "STUB: not implemented"; return nil }

// SetName sets the Name of the Property.
func (p *Property) SetName(n string) *Property { _ = "STUB: not implemented"; return nil }

// SetSchema sets the Schema of the Property.
func (p *Property) SetSchema(s *Schema) *Property { _ = "STUB: not implemented"; return nil }

func escapeRef(ref string) string { _ = "STUB: not implemented"; return "" }
