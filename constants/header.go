package constants

import "net/textproto"

var (
	XServiceName  = textproto.CanonicalMIMEHeaderKey("X-Service-Name")
	XApiKey       = textproto.CanonicalMIMEHeaderKey("X-API-Key")
	XRequestAt    = textproto.CanonicalMIMEHeaderKey("X-Request-At")
	Authorization = textproto.CanonicalMIMEHeaderKey("Authorization")
)
