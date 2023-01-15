package dingtalk

import (
	"context"
	"errors"
	"time"

	"github.com/imroc/req/v3"
)

type Result struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

var (
	myHTTPClient *req.Client
)

const (
	defaultTimeout = 5 * time.Second
)

func init() {
	myHTTPClient = initDefaultHTTPClient()
}

// initDefaultHTTPClient for connection re-use
func initDefaultHTTPClient() *req.Client {
	client := req.C().
		SetTimeout(defaultTimeout)
	return client
}

func doRequest(ctx context.Context, callMethod string, endPoint string, header map[string]string, body []byte) (*Result, error) {
	req := myHTTPClient.R()
	result := Result{}
	var errMsg string
	_, err := req.
		SetHeaders(header).
		SetBodyBytes(body).
		SetResult(&result).
		SetError(&errMsg).
		Post(endPoint)
	if err != nil {
		return nil, err
	}
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}

	return &result, nil
}
