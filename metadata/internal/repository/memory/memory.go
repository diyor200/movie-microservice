package memory

import (
	"context"
	"github.com/diyor200/movie-microservice/metadata/internal/repository"
	"github.com/diyor200/movie-microservice/metadata/pkg/model"
	"sync"
)

// Repository defines a memory movie metadata gateway.
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

func NewRepository() *Repository {
	return &Repository{data: make(map[string]*model.Metadata)}
}

func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.Lock()
	defer r.Unlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
