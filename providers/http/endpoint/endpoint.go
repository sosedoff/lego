package endpoint

import (
	"fmt"
	"net/http"
	"strings"
)

type EndpointProvider struct {
	url string
}

func NewEndpointProvider(url string) (*EndpointProvider, error) {
	return &EndpointProvider{url}, nil
}

func (p *EndpointProvider) Present(domain, token, keyAuth string) error {
	url := fmt.Sprintf("%s/%s", p.url, token)
	body := strings.NewReader(keyAuth)

	resp, err := http.Post(url, "text/plain", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Endpoint responded with", resp.StatusCode)
	}

	return nil
}

func (p *EndpointProvider) CleanUp(domain, token, keyAuth string) error {
	return nil
}
