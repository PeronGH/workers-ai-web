package utils

import (
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/valyala/fasthttp"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func HasEnv(key string) bool {
	value := os.Getenv(key)
	return value != ""
}

func PipeResponse(from *resty.Response, to *fasthttp.Response) {
	to.SetStatusCode(from.StatusCode())
	to.Header.SetContentType(from.Header().Get("Content-Type"))

	to.SetBodyStream(from.RawBody(), -1)
}
