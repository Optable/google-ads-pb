// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.


package clients

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/Optable/google-ads-pb/protogen/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newConversionActionClientHook clientHook

// ConversionActionCallOptions contains the retry settings for each method of ConversionActionClient.
type ConversionActionCallOptions struct {
	MutateConversionActions []gax.CallOption
}

func defaultConversionActionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("googleads.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultConversionActionCallOptions() *ConversionActionCallOptions {
	return &ConversionActionCallOptions{
		MutateConversionActions: []gax.CallOption{
			gax.WithTimeout(14400000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalConversionActionClient is an interface that defines the methods available from Google Ads API.
type internalConversionActionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateConversionActions(context.Context, *servicespb.MutateConversionActionsRequest, ...gax.CallOption) (*servicespb.MutateConversionActionsResponse, error)
}

// ConversionActionClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage conversion actions.
type ConversionActionClient struct {
	// The internal transport-dependent client.
	internalClient internalConversionActionClient

	// The call options for this service.
	CallOptions *ConversionActionCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConversionActionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConversionActionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ConversionActionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateConversionActions creates, updates or removes conversion actions. Operation statuses are
// returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ConversionActionError (at )
// CurrencyCodeError (at )
// DatabaseError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// NewResourceCreationError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// StringLengthError (at )
func (c *ConversionActionClient) MutateConversionActions(ctx context.Context, req *servicespb.MutateConversionActionsRequest, opts ...gax.CallOption) (*servicespb.MutateConversionActionsResponse, error) {
	return c.internalClient.MutateConversionActions(ctx, req, opts...)
}

// conversionActionGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type conversionActionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ConversionActionClient
	CallOptions **ConversionActionCallOptions

	// The gRPC API client.
	conversionActionClient servicespb.ConversionActionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewConversionActionClient creates a new conversion action service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage conversion actions.
func NewConversionActionClient(ctx context.Context, opts ...option.ClientOption) (*ConversionActionClient, error) {
	clientOpts := defaultConversionActionGRPCClientOptions()
	if newConversionActionClientHook != nil {
		hookOpts, err := newConversionActionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ConversionActionClient{CallOptions: defaultConversionActionCallOptions()}

	c := &conversionActionGRPCClient{
		connPool:    connPool,
		conversionActionClient: servicespb.NewConversionActionServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *conversionActionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *conversionActionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *conversionActionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *conversionActionGRPCClient) MutateConversionActions(ctx context.Context, req *servicespb.MutateConversionActionsRequest, opts ...gax.CallOption) (*servicespb.MutateConversionActionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateConversionActions[0:len((*c.CallOptions).MutateConversionActions):len((*c.CallOptions).MutateConversionActions)], opts...)
	var resp *servicespb.MutateConversionActionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionActionClient.MutateConversionActions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
