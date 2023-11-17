package api

import (
	"os"

	"github.com/PeronGH/workers-ai-web/internal/utils"
	"github.com/go-resty/resty/v2"
	"github.com/valyala/fasthttp"
)

func RunModel(model string, input interface{}) (*resty.Response, error) {
	return resty.New().
		NewRequest().
		SetBody(input).
		SetDoNotParseResponse(true).
		SetAuthToken(os.Getenv("CF_API_TOKEN")).
		Post("https://api.cloudflare.com/client/v4/accounts/" + os.Getenv("CF_ACCOUNT_ID") + "/ai/run/" + model)
}

func PipeRunModelToResponse(model string, input interface{}, response *fasthttp.Response) error {
	resp, err := RunModel(model, input)
	if err != nil {
		return err
	}

	utils.PipeResponse(resp, response)
	return nil
}
