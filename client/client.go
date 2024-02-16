package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"api.test.go/go-api-test/common"
)

var (
	dnsResolverIP        = "8.8.8.8:53"
	dnsResolverProto     = "udp"
	dnsResolverTimeoutMs = 2000
)

type ApiClientImpl struct {
	url          string
	httpClient   *http.Client
	httpRequest  *http.Request
	headers      map[string][]string
	httpResponse *http.Response
	response     string
	responseCode int
	timeout      time.Duration
}

func NewApiClient() *ApiClientImpl {
	return &ApiClientImpl{}
}

func (client *ApiClientImpl) WithUrl(url string) *ApiClientImpl {
	client.url = url
	return client
}

func (client *ApiClientImpl) CreateClient() error {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Duration(dnsResolverTimeoutMs) * time.Millisecond,
				}
				return d.DialContext(ctx, dnsResolverProto, dnsResolverIP)
			},
		},
	}
	if client.timeout != 0 {
		dialer.Timeout = client.timeout
	}
	dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
		return dialer.DialContext(ctx, network, address)
	}
	client.httpClient = &http.Client{Transport: &http.Transport{
		DialContext:     dialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	return nil
}

func (client *ApiClientImpl) Get(body []byte) {
	var err error
	client.httpRequest, err = http.NewRequest("GET", client.url, bytes.NewBuffer(body))
	if common.CheckNoError(err) {
		client.httpResponse, err = client.httpClient.Do(client.httpRequest)
	}
	if common.CheckNoError(err) {
		client.responseCode = client.httpResponse.StatusCode
	}
	fmt.Println("Response Code:", client.responseCode)

}
