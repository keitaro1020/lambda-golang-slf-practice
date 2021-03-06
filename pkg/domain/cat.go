package domain

import (
	"context"
)

// Cat is domain object
type Cat struct {
	ID     CatID  `json:"id"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

// Cats is domain object array
type Cats []*Cat

// CatID is ID
type CatID string

// CatClient is infrastructure (http client) interface
type CatClient interface {
	Search(ctx context.Context) (Cats, error)
}

// CatRepository is infrastructure (database) interface
type CatRepository interface {
	Get(ctx context.Context, id CatID) (*Cat, error)
	GetAll(ctx context.Context, first int64) (Cats, error)
	CreateInTx(ctx context.Context, tx Tx, cat *Cat) (*Cat, error)
}
