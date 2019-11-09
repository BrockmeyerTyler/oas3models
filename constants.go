package oasm

type HTTPVerb string

const (
	HttpGet     = "get"
	HttpPut     = "put"
	HttpPost    = "post"
	HttpDelete  = "delete"
	HttpOptions = "options"
	HttpHead    = "head"
	HttpPatch   = "patch"
	HttpTrace   = "trace"
)

const (
	DataTypeInteger  = "integer"
	DataTypeLong     = "long"
	DataTypeFloat    = "float"
	DataTypeDouble   = "double"
	DataTypeString   = "string"
	DataTypeByte     = "byte"
	DataTypeBinary   = "binary"
	DataTypeBoolean  = "boolean"
	DataTypeDate     = "date"
	DataTypeDatetime = "dateTime"
	DataTypePassword = "password"
)

const (
	MimeJson = "application/json"
)

const (
	InPath   = "path"
	InQuery  = "query"
	InHeader = "header"
	InCookie = "cookie"
)

const (
	StyleDefault        = ""
	StyleMatrix         = "matrix"
	StyleLabel          = "label"
	StyleForm           = "form"
	StyleSimple         = "simple"
	StyleSpaceDelimited = "spaceDelimited"
	StylePipeDelimited  = "pipeDelimited"
	StyleDeepObject     = "deepObject"
)

const (
	SecurityApiKey        = "apiKey"
	SecurityHttp          = "http"
	SecurityOauth2        = "oauth2"
	SecurityOpenIdConnect = "openIdConnect"
)

const (
	SecurityInQuery  = "query"
	SecurityInHeader = "header"
	SecurityInCookie = "cookie"
)
const (
	AuthBasic       = "basic"
	AuthBearer      = "bearer"
	AuthDigest      = "digest"
	AuthHoba        = "hoba"
	AuthMutual      = "mutual"
	AuthNegotiate   = "negotiate"
	AuthOauth       = "oauth"
	AuthScramSha1   = "scram-sha-1"
	AuthScramSha256 = "scram-sha-256"
	AuthVapid       = "vapid"
)
const (
	OAuthImplicit          = "implicit"
	OAuthPassword          = "password"
	OAuthClientCredentials = "clientCredentials"
	OAuthAuthorizationCode = "authorizationCode"
)
