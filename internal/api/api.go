package api

import (
	"io"
	"os"

	"github.com/go-resty/resty/v2"
)

func RunModel(model string, input interface{}, output io.Writer) error {
	client := resty.New()
	resp, err := client.R().
		SetBody(input).
		SetDoNotParseResponse(true).
		SetAuthToken(os.Getenv("CF_API_TOKEN")).
		Post("https://api.cloudflare.com/client/v4/accounts/" + os.Getenv("CF_ACCOUNT_ID") + "/ai/run/" + model)

	if err != nil {
		return err
	}

	bodyReader := resp.RawBody()
	defer bodyReader.Close()

	_, err = io.Copy(output, bodyReader)

	return err
}
