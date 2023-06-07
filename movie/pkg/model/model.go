package model

import model "movieexample.com/metadata/pkg"

// MovieDetails includes movie metadata its aggregated
// rating.
type MovieDetails struct {
	Rating   *float64
	Metadata model.Metadata
}
