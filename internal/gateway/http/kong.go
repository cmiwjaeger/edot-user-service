package http

import (
	"edot-monorepo/services/user-service/internal/model"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type KongClient struct {
	HttpClient *resty.Client
	BaseUrl    string
}

func NewKongClient(httpClient *resty.Client, baseUrl string) *KongClient {
	return &KongClient{
		HttpClient: httpClient,
		BaseUrl:    baseUrl,
	}
}

func (k *KongClient) CreateConsumer(data model.KongConsumerRequest) (*model.KongConsumerResponse, error) {
	response := &model.KongConsumerResponse{}

	resp, err := k.HttpClient.R().EnableTrace().SetBody(data).SetResult(response).Post(fmt.Sprintf("%s/consumers/edot/jwt", k.BaseUrl))
	curlCmdExecuted := resp.Request.GenerateCurlCommand()

	// Explore curl command
	fmt.Println("Curl Command:\n  ", curlCmdExecuted+"\n")

	fmt.Println(resp.StatusCode())
	fmt.Println(resp.Error())
	fmt.Println(resp)

	if err != nil {
		return nil, err
	}

	return response, err
}
