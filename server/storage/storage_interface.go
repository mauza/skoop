package storage

import (
	"context"
	skoopv1 "github.com/mauza/skoop/protos/gen/go/skoop/v1"
)

var Storage StorageInterface

type StorageInterface interface {
	Initialize() (shutdown func(), err error)
	Ready() bool
	UpsertHellos(ctx context.Context, request *skoopv1.UpsertHellosRequest) ([]*skoopv1.Hello, error)
	DeleteHellos(ctx context.Context, ids []string) error
	ListHellos(ctx context.Context, request *skoopv1.ListRequest) ([]*skoopv1.Hello, error)
	GetHellos(ctx context.Context, request *skoopv1.GetRequest) ([]*skoopv1.Hello, error)
	GetHello(ctx context.Context, id string) (*skoopv1.Hello, error)
}
