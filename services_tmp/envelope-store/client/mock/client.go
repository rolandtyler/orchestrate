package clientmock

import (
	"context"

	evlpstore "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/services_tmp/envelope-store"
	mock "gitlab.com/ConsenSys/client/fr/core-stack/service/envelope-store.git/services_tmp/mock"
	"google.golang.org/grpc"
)

// EnvelopeStoreClient is a client that wraps an EnvelopeStoreServer into an EnvelopeStoreClient
type EnvelopeStoreClient struct {
	srv evlpstore.EnvelopeStoreServer
}

func New() *EnvelopeStoreClient {
	return &EnvelopeStoreClient{
		srv: mock.NewEnvelopeStore(),
	}
}

func (client *EnvelopeStoreClient) Store(ctx context.Context, in *evlpstore.StoreRequest, opts ...grpc.CallOption) (*evlpstore.StoreResponse, error) {
	return client.srv.Store(ctx, in)
}

// Load envelope by identifier
func (client *EnvelopeStoreClient) LoadByID(ctx context.Context, in *evlpstore.LoadByIDRequest, opts ...grpc.CallOption) (*evlpstore.StoreResponse, error) {
	return client.srv.LoadByID(ctx, in)
}

// Load Envelope by transaction hash
func (client *EnvelopeStoreClient) LoadByTxHash(ctx context.Context, in *evlpstore.LoadByTxHashRequest, opts ...grpc.CallOption) (*evlpstore.StoreResponse, error) {
	return client.srv.LoadByTxHash(ctx, in)
}

// SetStatus set an envelope status
func (client *EnvelopeStoreClient) SetStatus(ctx context.Context, in *evlpstore.SetStatusRequest, opts ...grpc.CallOption) (*evlpstore.StatusResponse, error) {
	return client.srv.SetStatus(ctx, in)

}

// LoadPending load envelopes of pending transactions
func (client *EnvelopeStoreClient) LoadPending(ctx context.Context, in *evlpstore.LoadPendingRequest, opts ...grpc.CallOption) (*evlpstore.LoadPendingResponse, error) {
	return client.srv.LoadPending(ctx, in)

}
