package oas3models

type DataType string
const (
	DataTypeInteger  = DataType("integer")
	DataTypeLong     = DataType("long")
	DataTypeFloat    = DataType("float")
	DataTypeDouble   = DataType("double")
	DataTypeString   = DataType("string")
	DataTypeByte     = DataType("byte")
	DataTypeBinary   = DataType("binary")
	DataTypeBoolean  = DataType("boolean")
	DataTypeDate     = DataType("date")
	DataTypeDatetime = DataType("dateTime")
	DataTypePassword = DataType("password")
)

type MimeType string
const (
	MimeJson = MimeType("application/json")
)

type InRequest string
const (
	InPath   = InRequest("path")
	InQuery  = InRequest("query")
	InHeader = InRequest("header")
	InCookie = InRequest("cookie")
)

type StyleSetting string
const (
	StyleDefault        = StyleSetting("")
	StyleMatrix         = StyleSetting("matrix")
	StyleLabel          = StyleSetting("label")
	StyleForm           = StyleSetting("form")
	StyleSimple         = StyleSetting("simple")
	StyleSpaceDelimited = StyleSetting("spaceDelimited")
	StylePipeDelimited  = StyleSetting("pipeDelimited")
	StyleDeepObject     = StyleSetting("deepObject")
)

type SecurityType string
const (
	SecurityApiKey        = SecurityType("apiKey")
	SecurityHttp          = SecurityType("http")
	SecurityOauth2        = SecurityType("oauth2")
	SecurityOpenIdConnect = SecurityType("openIdConnect")
)

type SecurityInRequest string
const (
	SecurityInQuery  = SecurityInRequest("query")
	SecurityInHeader = SecurityInRequest("header")
	SecurityInCookie = SecurityInRequest("cookie")
)

type HTTPAuthScheme string
const (
	AuthBasic       = HTTPAuthScheme("basic")
	AuthBearer      = HTTPAuthScheme("bearer")
	AuthDigest      = HTTPAuthScheme("digest")
	AuthHoba        = HTTPAuthScheme("hoba")
	AuthMutual      = HTTPAuthScheme("mutual")
	AuthNegotiate   = HTTPAuthScheme("negotiate")
	AuthOauth       = HTTPAuthScheme("oauth")
	AuthScramSha1   = HTTPAuthScheme("scram-sha-1")
	AuthScramSha256 = HTTPAuthScheme("scram-sha-256")
	AuthVapid       = HTTPAuthScheme("vapid")
)

type OAuthFlowType string
const (
	OAuthImplicit          = OAuthFlowType("implicit")
	OAuthPassword          = OAuthFlowType("password")
	OAuthClientCredentials = OAuthFlowType("clientCredentials")
	OAuthAuthorizationCode = OAuthFlowType("authorizationCode")
)
