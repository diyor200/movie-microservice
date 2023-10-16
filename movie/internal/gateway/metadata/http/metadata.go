package http

import (
	"context"
	"encoding/json"
	"github.com/diyor200/movie-microservice/metadata/pkg/model"
	"github.com/diyor200/movie-microservice/movie/internal/gateway"
	"net/http"
)

type Gateway struct {
	addr string
}

func NewGateway(addr string) *Gateway {
	return &Gateway{addr: addr}
}

func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	req, err := http.NewRequest(http.MethodGet, g.addr+"/metadata", nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", id)
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, gateway.ErrNotFound
	}
	var v *model.Metadata
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return nil, err
	}
	return v, nil
}
