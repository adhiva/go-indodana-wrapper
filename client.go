package indodana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	APIEnvType Env
	APIKey     string
	SecretKey  string
	LogLevel   int
	Logger     *log.Logger
}

func NewClient() Client {
	return Client{
		APIEnvType: Sandbox,
		LogLevel:   2,
		Logger:     log.New(os.Stderr, "", log.LstdFlags),
	}
}

var defHTTPTimeout = 80 * time.Second
var httpClient = &http.Client{Timeout: defHTTPTimeout}

func (c *Client) DoRequest(method, path string, body interface{}, v interface{}) error {
	byteJson, _ := json.Marshal(body)
	logLevel := c.LogLevel
	logger := c.Logger

	req, err := http.NewRequest(method, path, bytes.NewBuffer(byteJson))
	if err != nil {
		if logLevel > 0 {
			logger.Println("Request creation failed: ", err)
		}
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", CreateSignatureHeaders(c.APIKey, c.SecretKey)))
	start := time.Now()

	if logLevel > 1 {
		logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
	}

	// Set Log Request
	log := Logger{
		Method:      req.Method,
		Host:        req.URL.Host,
		Path:        req.URL.Path,
		RequestBody: body,
	}
	logger.Println(Stringify(log))
	res, err := httpClient.Do(req)
	if err != nil {
		logger.Println("Cannot send request : ", err)
		return err
	}

	// Println stdout for checking request body and response
	logger.Println("Completed in ", time.Since(start))
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Println("Cannot read response body: ", err)
		return err
	}

	if v != nil {
		response := json.Unmarshal(resBody, v)
		log.ResponseBody = v
		log.TimeLength = int(time.Since(start))
		logger.Println(Stringify(log))

		return response
	}

	return nil
}
