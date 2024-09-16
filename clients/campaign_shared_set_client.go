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

var newCampaignSharedSetClientHook clientHook

// CampaignSharedSetCallOptions contains the retry settings for each method of CampaignSharedSetClient.
type CampaignSharedSetCallOptions struct {
	MutateCampaignSharedSets []gax.CallOption
}

func defaultCampaignSharedSetGRPCClientOptions() []option.ClientOption {
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

func defaultCampaignSharedSetCallOptions() *CampaignSharedSetCallOptions {
	return &CampaignSharedSetCallOptions{
		MutateCampaignSharedSets: []gax.CallOption{
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

// internalCampaignSharedSetClient is an interface that defines the methods available from Google Ads API.
type internalCampaignSharedSetClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCampaignSharedSets(context.Context, *servicespb.MutateCampaignSharedSetsRequest, ...gax.CallOption) (*servicespb.MutateCampaignSharedSetsResponse, error)
}

// CampaignSharedSetClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage campaign shared sets.
type CampaignSharedSetClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignSharedSetClient

	// The call options for this service.
	CallOptions *CampaignSharedSetCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignSharedSetClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignSharedSetClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CampaignSharedSetClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCampaignSharedSets creates or removes campaign shared sets. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CampaignSharedSetError (at )
// ContextError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// FieldError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *CampaignSharedSetClient) MutateCampaignSharedSets(ctx context.Context, req *servicespb.MutateCampaignSharedSetsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignSharedSetsResponse, error) {
	return c.internalClient.MutateCampaignSharedSets(ctx, req, opts...)
}

// campaignSharedSetGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignSharedSetGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CampaignSharedSetClient
	CallOptions **CampaignSharedSetCallOptions

	// The gRPC API client.
	campaignSharedSetClient servicespb.CampaignSharedSetServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCampaignSharedSetClient creates a new campaign shared set service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage campaign shared sets.
func NewCampaignSharedSetClient(ctx context.Context, opts ...option.ClientOption) (*CampaignSharedSetClient, error) {
	clientOpts := defaultCampaignSharedSetGRPCClientOptions()
	if newCampaignSharedSetClientHook != nil {
		hookOpts, err := newCampaignSharedSetClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CampaignSharedSetClient{CallOptions: defaultCampaignSharedSetCallOptions()}

	c := &campaignSharedSetGRPCClient{
		connPool:    connPool,
		campaignSharedSetClient: servicespb.NewCampaignSharedSetServiceClient(connPool),
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
func (c *campaignSharedSetGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignSharedSetGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignSharedSetGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignSharedSetGRPCClient) MutateCampaignSharedSets(ctx context.Context, req *servicespb.MutateCampaignSharedSetsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignSharedSetsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateCampaignSharedSets[0:len((*c.CallOptions).MutateCampaignSharedSets):len((*c.CallOptions).MutateCampaignSharedSets)], opts...)
	var resp *servicespb.MutateCampaignSharedSetsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignSharedSetClient.MutateCampaignSharedSets(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
