package postgres

import (
	"context"
	skoopv1 "github.com/mauza/skoop/protos/gen/go/skoop/v1"
	"github.com/joomcode/errorx"
)

func (p PostgresStorage) UpsertHellos(ctx context.Context, request *skoopv1.UpsertHellosRequest) ([]*skoopv1.Hello, error) {
	helloProtos := skoopv1.HelloProtos(request.Hellos)
	err := helloProtos.Upsert(ctx, db, nil, nil, true)
	return helloProtos, err
}

func (p PostgresStorage) DeleteHellos(ctx context.Context, ids []string) error {
	return skoopv1.DeleteHelloGormModels(ctx, db, ids)
}

func (p PostgresStorage) ListHellos(ctx context.Context, request *skoopv1.ListRequest) ([]*skoopv1.Hello, error) {
	protos := skoopv1.HelloProtos{}
	var orderBy string
	if orderBy = request.OrderBy; orderBy == "" {
		orderBy = "created_at"
	}
	err := protos.List(ctx, db, int(request.Limit), int(request.Offset), orderBy)
	return protos, err
}

func (p PostgresStorage) GetHellos(ctx context.Context, request *skoopv1.GetRequest) ([]*skoopv1.Hello, error) {
	protos := skoopv1.HelloProtos{}
	err := protos.GetByIds(ctx, db, request.Ids)
	return protos, err
}

func (p PostgresStorage) GetHello(ctx context.Context, id string) (*skoopv1.Hello, error) {
	var err error
	protos := skoopv1.HelloProtos{}
	if err = protos.GetByIds(ctx, db, []string{id}); err != nil {
		return nil, err
	}
	if len(protos) == 0 {
		return nil, errorx.IllegalArgument.New("hello with id %s not found", id)
	}
	return protos[0], nil
}
