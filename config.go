package Warp10Helper

var warpURI = "/api/v0/update"

// SetURI is changing the classic URI /api/v0/update
func SetURI(uri string) {
	warpURI = uri
}

var warpHeader = "X-Warp10-Token"

// SetHeader is changing the classic Header X-Warp10-Token
func SetHeader(header string) {
	warpHeader = header
}
