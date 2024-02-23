package handlers

import (
	"context"
	skoopv1 "github.com/mauza/skoop/protos/gen/go/skoop/v1"
	"github.com/mauza/skoop/server/storage"
)

func (a ApiServer) UpsertHellos(ctx context.Context, request *skoopv1.UpsertHellosRequest) (*skoopv1.Hellos, error) {
	upsertedHellos, err := storage.Storage.UpsertHellos(ctx, request)
	return &skoopv1.Hellos{Hellos: upsertedHellos}, err
}

func (a ApiServer) DeleteHellos(ctx context.Context, request *skoopv1.DeleteRequest) (*skoopv1.DeleteResponse, error) {
	return &skoopv1.DeleteResponse{}, storage.Storage.DeleteHellos(ctx, request.Ids)
}

func (a ApiServer) ListHellos(ctx context.Context, request *skoopv1.ListRequest) (*skoopv1.Hellos, error) {
	if request.Limit == 0 {
		request.Limit = 100
	}
	hellos, err := storage.Storage.ListHellos(ctx, request)
	return &skoopv1.Hellos{Hellos: hellos}, err
}

func (a ApiServer) GetHellos(ctx context.Context, request *skoopv1.GetRequest) (*skoopv1.Hellos, error) {
	hellos, err := storage.Storage.GetHellos(ctx, request)
	return &skoopv1.Hellos{Hellos: hellos}, err
}
