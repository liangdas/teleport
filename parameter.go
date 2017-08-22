// Copyright 2015-2017 HenryLee. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package teleport

// Packet Header types
const (
	TypeUndefined int32 = 0
	TypePull      int32 = 1
	TypeReply     int32 = 2 // reply to pull
	TypePush      int32 = 4
	// TypeAuth      int32 = 5
	// TypeHeartbeat int32 = 6
)

// Response Header status codes as registered with IANA.
const (
	StatusWriteFailed  = 100
	StatusConnClosed   = 101
	StatusOK           = 200
	StatusBadUri       = 400
	StatusNotFound     = 404
	StatusFailedPlugin = 424

	StatusProcessing = 102 // RFC 2518, 10.1

	StatusCreated              = 201 // RFC 7231, 6.3.2
	StatusAccepted             = 202 // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
	StatusNoContent            = 204 // RFC 7231, 6.3.5
	StatusResetContent         = 205 // RFC 7231, 6.3.6
	StatusPartialContent       = 206 // RFC 7233, 4.1
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently  = 301 // RFC 7231, 6.4.2
	StatusFound             = 302 // RFC 7231, 6.4.3
	StatusSeeOther          = 303 // RFC 7231, 6.4.4
	StatusNotModified       = 304 // RFC 7232, 4.1
	StatusUseProxy          = 305 // RFC 7231, 6.4.5
	StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect = 308 // RFC 7538, 3

	StatusUnauthorized               = 401 // RFC 7235, 3.1
	StatusPaymentRequired            = 402 // RFC 7231, 6.5.2
	StatusForbidden                  = 403 // RFC 7231, 6.5.3
	StatusMethodNotAllowed           = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable              = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired          = 407 // RFC 7235, 3.2
	StatusPullTimeout                = 408 // RFC 7231, 6.5.7
	StatusConflict                   = 409 // RFC 7231, 6.5.8
	StatusGone                       = 410 // RFC 7231, 6.5.9
	StatusLengthRequired             = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed         = 412 // RFC 7232, 4.2
	StatusPullEntityTooLarge         = 413 // RFC 7231, 6.5.11
	StatusPullURITooLong             = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType       = 415 // RFC 7231, 6.5.13
	StatusPulledRangeNotSatisfiable  = 416 // RFC 7233, 4.4
	StatusExpectationFailed          = 417 // RFC 7231, 6.5.14
	StatusTeapot                     = 418 // RFC 7168, 2.3.3
	StatusUnprocessableEntity        = 422 // RFC 4918, 11.2
	StatusLocked                     = 423 // RFC 4918, 11.3
	StatusUpgradeRequired            = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired       = 428 // RFC 6585, 3
	StatusTooManyPulls               = 429 // RFC 6585, 4
	StatusPullHeaderFieldsTooLarge   = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

var statusText = map[int]string{
	StatusWriteFailed: "write failed",
	StatusConnClosed:  "Connection Closed",
	StatusProcessing:  "Processing",

	StatusOK:                   "OK",
	StatusCreated:              "Created",
	StatusAccepted:             "Accepted",
	StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	StatusNoContent:            "No Content",
	StatusResetContent:         "Reset Content",
	StatusPartialContent:       "Partial Content",
	StatusMultiStatus:          "Multi-Status",
	StatusAlreadyReported:      "Already Reported",
	StatusIMUsed:               "IM Used",

	StatusMultipleChoices:   "Multiple Choices",
	StatusMovedPermanently:  "Moved Permanently",
	StatusFound:             "Found",
	StatusSeeOther:          "See Other",
	StatusNotModified:       "Not Modified",
	StatusUseProxy:          "Use Proxy",
	StatusTemporaryRedirect: "Temporary Redirect",
	StatusPermanentRedirect: "Permanent Redirect",

	StatusBadUri:                     "Bad URI",
	StatusUnauthorized:               "Unauthorized",
	StatusPaymentRequired:            "Payment Required",
	StatusForbidden:                  "Forbidden",
	StatusNotFound:                   "Not Found",
	StatusMethodNotAllowed:           "Method Not Allowed",
	StatusNotAcceptable:              "Not Acceptable",
	StatusProxyAuthRequired:          "Proxy Authentication Required",
	StatusPullTimeout:                "Pull Timeout",
	StatusConflict:                   "Conflict",
	StatusGone:                       "Gone",
	StatusLengthRequired:             "Length Required",
	StatusPreconditionFailed:         "Precondition Failed",
	StatusPullEntityTooLarge:         "Pull Entity Too Large",
	StatusPullURITooLong:             "Pull URI Too Long",
	StatusUnsupportedMediaType:       "Unsupported Media Type",
	StatusPulledRangeNotSatisfiable:  "Pulled Range Not Satisfiable",
	StatusExpectationFailed:          "Expectation Failed",
	StatusTeapot:                     "I'm a teapot",
	StatusUnprocessableEntity:        "Unprocessable Entity",
	StatusLocked:                     "Locked",
	StatusFailedPlugin:               "Failed Plugin",
	StatusUpgradeRequired:            "Upgrade Required",
	StatusPreconditionRequired:       "Precondition Required",
	StatusTooManyPulls:               "Too Many Pulls",
	StatusPullHeaderFieldsTooLarge:   "Pull Header Fields Too Large",
	StatusUnavailableForLegalReasons: "Unavailable For Legal Reasons",

	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

// StatusText returns a text for the Response Header status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
