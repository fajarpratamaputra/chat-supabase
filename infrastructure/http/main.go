package http

import (
	"github.com/gojek/heimdall/v7/hystrix"
	"time"
)

type Client struct {
	*hystrix.Client
}

func NewClient() *Client {
	client := hystrix.NewClient(
		hystrix.WithHTTPTimeout(10*time.Millisecond),
		hystrix.WithHystrixTimeout(1000*time.Millisecond),
		hystrix.WithMaxConcurrentRequests(30),
		hystrix.WithErrorPercentThreshold(20),
	)
	return &Client{client}
}
