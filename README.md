go-indodana-wrapper
===================

This library crated for simplify your code to use Indodana Payment Method API.

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:

	go get github.com/adhiva/go-indodana-wrapper

After it the package is ready to use.


#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/adhiva/go-indodana-wrapper"
```
If you are unhappy to use long `go-indodana`, you can do something like this:
```go
import (
  indodana "github.com/adhiva/go-indodana-wrapper"
)
```

#### List of wrapper indodana client:
```go
func GetInstallment(req *RequestGetInstallment) (ResponseGetInstallment, error)
func CheckStatus(req *RequestCheckStatus) (ResponseCheckStatus, error)
func CheckoutPayment(req *RequestCheckoutPayment) (ResponseCheckoutPayment, error)
```

#### Examples

###### Initial The Client
```go
func init() {
	indodanaClient = indodana.NewClient()
	indodanaClient.APIKey = config.Config.Indodana.APIKey
	indodanaClient.SecretKey = config.Config.Indodana.SecretKey

	if config.Config.Indodana.ENV == "production" {
		indodanaClient.APIEnvType = indodana.Production
	} else {
		indodanaClient.APIEnvType = indodana.Sandbox
	}

	coreGateway = indodana.CoreGateway{
		Client: indodanaClient,
	}
}
```
###### Get Installment
```go
func GetInstallment(request *indodana.RequestGetInstallment) (*indodana.ResponseGetInstallment, error) {
	var (
		resp *indodana.ResponseCharge
		err  error
	)

	resp, err = coreGateway.GetInstallment(request)

	if err != nil {
		return nil, err
    }
    
	return resp, nil
}
```

###### Checkout Payment
```go
func CheckoutPayment(request *indodana.RequestCheckoutPayment) (*indodana.ResponseCheckoutPayment, error) {
	var (
		resp *indodana.ResponseCheckoutPayment
		err  error
	)

	// Generate deeplink for apps and mobile-web only
	resp, err = coreGateway.CheckoutPayment(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
```

###### Check Status
```go
	func CheckStatus(merchantOrderID string) (*indodana.ResponseCheckStatus, error) {
		var (
			request *indodana.RequestCheckStatus
			resp    *indodana.ResponseCheckStatus
			err     error
		)

		request.MerchantOrderId = merchantOrderID

		// Generate deeplink for apps and mobile-web only
		resp, err = coreGateway.CheckStatus(request)

		if err != nil {
			return nil, err
		}
		return resp, nil
	}
```

#### Support
If you do have a contribution to the package, feel free to create a Pull Request or an Issue.

#### What to contribute
If you don't know what to do, there are some features and functions that need to be done
- [ ] Cancel Refund Transaction
- [x] Get Installment
- [x] Checkout Payment
- [x] Check Status

#### Advice
Feel free to create what you want, but keep in mind when you implement new features:
- Code must be clear and readable, names of variables/constants clearly describes what they are doing
- Wrapper they are created must be documented and described in source file and added to README.md to the list of wrapper indodana client

### Contributors 
- [Adhitya Giva Muhammad](https://github.com/adhiva)