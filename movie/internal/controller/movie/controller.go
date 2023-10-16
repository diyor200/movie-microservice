package movie

import (
	"context"
	"errors"
	metadatamodel "github.com/diyor200/movie-microservice/metadata/pkg/model"
	"github.com/diyor200/movie-microservice/movie/internal/gateway"
	"github.com/diyor200/movie-microservice/movie/pkg/model"
	ratingmodel "github.com/diyor200/movie-microservice/rating/pkg/model"
)

var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}
type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

func NewController(ratingGateway ratingGateway, metadatGateway metadataGateway) *Controller {
	return &Controller{ratingGateway: ratingGateway, metadataGateway: metadatGateway}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	details := &model.MovieDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordMovieType)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {

	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}
