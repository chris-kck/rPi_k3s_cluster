package acquisition

import "context"

// Repository TODO
type Repository interface {
	AcquireDatum(context.Context, Datum) error
}
