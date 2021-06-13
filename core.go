package indodana

import (
	"errors"
	"fmt"

	"github.com/google/go-querystring/query"
)

type CoreGateway struct {
	Client Client
}

func (gateway *CoreGateway) GetInstallment(req *RequestGetInstallment) (ResponseGetInstallment, error) {
	host := gateway.Client.APIEnvType.String()
	response := ResponseGetInstallment{}
	path := host + "/v2/payment_calculation"

	err := gateway.Client.DoRequest("POST", path, req, &response)
	if err != nil {
		errMessage := fmt.Sprintf("Indodana failed get installment : %s", err.Error())
		fmt.Println(errMessage)
		return response, errors.New(errMessage)
	}

	if response.Error != (Error{}) {
		errMessage := fmt.Sprintf("Indodana failed get installment : %s", response.Error.Kind)
		fmt.Println(errMessage)
		return response, errors.New(errMessage)
	}

	return response, nil
}

func (gateway *CoreGateway) CheckStatus(req *RequestCheckStatus) (ResponseCheckStatus, error) {
	host := gateway.Client.APIEnvType.String()
	response := ResponseCheckStatus{}
	path := host + "/v1/transactions/check_status"

	if req != nil {
		queryParam, _ := query.Values(req)
		path = fmt.Sprintf("%s%s?%s", host, "/v1/transactions/check_status", queryParam.Encode())
	}

	err := gateway.Client.DoRequest("GET", path, req, &response)
	if err != nil {
		errMessage := fmt.Sprintf("Indodana failed get status : %s", err.Error())
		fmt.Println(errMessage)
		return response, errors.New(errMessage)
	}

	return response, nil
}

func (gateway *CoreGateway) CheckoutPayment(req *RequestCheckoutPayment) (ResponseCheckoutPayment, error) {
	host := gateway.Client.APIEnvType.String()
	response := ResponseCheckoutPayment{}
	path := host + "/v2/checkout_url"

	err := gateway.Client.DoRequest("POST", path, req, &response)
	if err != nil {
		errMessage := fmt.Sprintf("Indodana failed checkout payment : %s", err.Error())
		fmt.Println(errMessage)
		return response, errors.New(errMessage)
	}

	if response.Error != (Error{}) {
		errMessage := fmt.Sprintf("Indodana failed checkout payment : %s", response.Error.Kind)
		fmt.Println(errMessage)
		return response, errors.New(errMessage)
	}

	return response, nil
}
