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

var newCampaignGroupClientHook clientHook

// CampaignGroupCallOptions contains the retry settings for each method of CampaignGroupClient.
type CampaignGroupCallOptions struct {
	MutateCampaignGroups []gax.CallOption
}

func defaultCampaignGroupGRPCClientOptions() []option.ClientOption {
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

func defaultCampaignGroupCallOptions() *CampaignGroupCallOptions {
	return &CampaignGroupCallOptions{
		MutateCampaignGroups: []gax.CallOption{
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

// internalCampaignGroupClient is an interface that defines the methods available from Google Ads API.
type internalCampaignGroupClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCampaignGroups(context.Context, *servicespb.MutateCampaignGroupsRequest, ...gax.CallOption) (*servicespb.MutateCampaignGroupsResponse, error)
}

// CampaignGroupClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage campaign groups.
type CampaignGroupClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignGroupClient

	// The call options for this service.
	CallOptions *CampaignGroupCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignGroupClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignGroupClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CampaignGroupClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCampaignGroups creates, updates, or removes campaign groups. Operation statuses are
// returned.
func (c *CampaignGroupClient) MutateCampaignGroups(ctx context.Context, req *servicespb.MutateCampaignGroupsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignGroupsResponse, error) {
	return c.internalClient.MutateCampaignGroups(ctx, req, opts...)
}

// campaignGroupGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignGroupGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CampaignGroupClient
	CallOptions **CampaignGroupCallOptions

	// The gRPC API client.
	campaignGroupClient servicespb.CampaignGroupServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCampaignGroupClient creates a new campaign group service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage campaign groups.
func NewCampaignGroupClient(ctx context.Context, opts ...option.ClientOption) (*CampaignGroupClient, error) {
	clientOpts := defaultCampaignGroupGRPCClientOptions()
	if newCampaignGroupClientHook != nil {
		hookOpts, err := newCampaignGroupClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CampaignGroupClient{CallOptions: defaultCampaignGroupCallOptions()}

	c := &campaignGroupGRPCClient{
		connPool:    connPool,
		campaignGroupClient: servicespb.NewCampaignGroupServiceClient(connPool),
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
func (c *campaignGroupGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignGroupGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignGroupGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignGroupGRPCClient) MutateCampaignGroups(ctx context.Context, req *servicespb.MutateCampaignGroupsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignGroupsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateCampaignGroups[0:len((*c.CallOptions).MutateCampaignGroups):len((*c.CallOptions).MutateCampaignGroups)], opts...)
	var resp *servicespb.MutateCampaignGroupsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignGroupClient.MutateCampaignGroups(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
