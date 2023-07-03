package grpc

import (
	"context"

	"movieexample.com/gen"
	"movieexample.com/internal/grpcutil"
	"movieexample.com/pkg/discovery"
	model "movieexample.com/rating/pkg"
)

// Gateway defines a gRPC gateway for a rating service.
type Gateway struct {
	registry discovery.Registry
}

// New creates a new gRPC gateway for a rating service.
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

// GetAggregatedRating returns aggregated rating for a record or ErrNotFound if there no ratings for it.
func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "rating", g.registry)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	client := gen.NewRatingServiceClient(conn)
	resp, err := client.GetAggregatedRating(ctx, &gen.GetAggregatedRatingRequest{
		RecordId:   string(recordID),
		RecordType: string(recordType),
	})
	if err != nil {
		return 0, err
	}
	return resp.RatingValue, nil
}
