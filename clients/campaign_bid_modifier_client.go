// Copyright 2023 Google LLC
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

var newCampaignBidModifierClientHook clientHook

// CampaignBidModifierCallOptions contains the retry settings for each method of CampaignBidModifierClient.
type CampaignBidModifierCallOptions struct {
	MutateCampaignBidModifiers []gax.CallOption
}

func defaultCampaignBidModifierGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCampaignBidModifierCallOptions() *CampaignBidModifierCallOptions {
	return &CampaignBidModifierCallOptions{
		MutateCampaignBidModifiers: []gax.CallOption{
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

// internalCampaignBidModifierClient is an interface that defines the methods available from Google Ads API.
type internalCampaignBidModifierClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCampaignBidModifiers(context.Context, *servicespb.MutateCampaignBidModifiersRequest, ...gax.CallOption) (*servicespb.MutateCampaignBidModifiersResponse, error)
}

// CampaignBidModifierClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage campaign bid modifiers.
type CampaignBidModifierClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignBidModifierClient

	// The call options for this service.
	CallOptions *CampaignBidModifierCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignBidModifierClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignBidModifierClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CampaignBidModifierClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCampaignBidModifiers creates, updates, or removes campaign bid modifiers.
// Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ContextError (at )
// CriterionError (at )
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
func (c *CampaignBidModifierClient) MutateCampaignBidModifiers(ctx context.Context, req *servicespb.MutateCampaignBidModifiersRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignBidModifiersResponse, error) {
	return c.internalClient.MutateCampaignBidModifiers(ctx, req, opts...)
}

// campaignBidModifierGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignBidModifierGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CampaignBidModifierClient
	CallOptions **CampaignBidModifierCallOptions

	// The gRPC API client.
	campaignBidModifierClient servicespb.CampaignBidModifierServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCampaignBidModifierClient creates a new campaign bid modifier service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage campaign bid modifiers.
func NewCampaignBidModifierClient(ctx context.Context, opts ...option.ClientOption) (*CampaignBidModifierClient, error) {
	clientOpts := defaultCampaignBidModifierGRPCClientOptions()
	if newCampaignBidModifierClientHook != nil {
		hookOpts, err := newCampaignBidModifierClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CampaignBidModifierClient{CallOptions: defaultCampaignBidModifierCallOptions()}

	c := &campaignBidModifierGRPCClient{
		connPool:    connPool,
		campaignBidModifierClient: servicespb.NewCampaignBidModifierServiceClient(connPool),
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
func (c *campaignBidModifierGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignBidModifierGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignBidModifierGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignBidModifierGRPCClient) MutateCampaignBidModifiers(ctx context.Context, req *servicespb.MutateCampaignBidModifiersRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignBidModifiersResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateCampaignBidModifiers[0:len((*c.CallOptions).MutateCampaignBidModifiers):len((*c.CallOptions).MutateCampaignBidModifiers)], opts...)
	var resp *servicespb.MutateCampaignBidModifiersResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignBidModifierClient.MutateCampaignBidModifiers(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
