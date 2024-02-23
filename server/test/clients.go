package test

import (
	"context"
	"time"

	skoopv1 "github.com/mauza/skoop/protos/gen/go/skoop/v1"
	"google.golang.org/grpc"
)

var ApiClient = initializeApiClient(5 * time.Second)

func initializeApiClient(timeout time.Duration) skoopv1.ApiClient {
	connectTo := "127.0.0.1:6000"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, connectTo, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return skoopv1.NewApiClient(conn)
}

func UpsertHellos(hellos []*skoopv1.Hello) ([]*skoopv1.Hello, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.UpsertHellos(ctx, &skoopv1.UpsertHellosRequest{Hellos: hellos})
	if err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else {
		return response.Hellos, err
	}
}

func DeleteHellos(ids []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ApiClient.DeleteHellos(ctx, &skoopv1.DeleteRequest{Ids: ids})
	return err
}

func ListHellos(limit, offset int, orderBy string) ([]*skoopv1.Hello, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.ListHellos(ctx, &skoopv1.ListRequest{Limit: int32(limit), Offset: int32(offset), OrderBy: orderBy})
	if err != nil {
		return nil, err
	}
	return response.Hellos, err
}

func GetHellosById(ids []string) ([]*skoopv1.Hello, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.GetHellos(ctx, &skoopv1.GetRequest{Ids: ids})
	return response.Hellos, err
}
