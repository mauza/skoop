package handlers

import (
	"context"
	"errors"
	skoopv1 "github.com/mauza/skoop/protos/gen/go/skoop/v1"
	"github.com/mauza/skoop/server/storage"
)

type ApiServer struct{}

func (a ApiServer) Healthy(ctx context.Context, empty *skoopv1.Empty) (*skoopv1.Empty, error) {
	return &skoopv1.Empty{}, nil
}

func (a ApiServer) Ready(ctx context.Context, empty *skoopv1.Empty) (*skoopv1.Empty, error) {
	if storage.Storage.Ready() {
		return &skoopv1.Empty{}, nil
	}
	return nil, errors.New("")
}
