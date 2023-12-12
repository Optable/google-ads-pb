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

var newAudienceInsightsClientHook clientHook

// AudienceInsightsCallOptions contains the retry settings for each method of AudienceInsightsClient.
type AudienceInsightsCallOptions struct {
	GenerateInsightsFinderReport []gax.CallOption
	ListAudienceInsightsAttributes []gax.CallOption
	ListInsightsEligibleDates []gax.CallOption
	GenerateAudienceCompositionInsights []gax.CallOption
}

func defaultAudienceInsightsGRPCClientOptions() []option.ClientOption {
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

func defaultAudienceInsightsCallOptions() *AudienceInsightsCallOptions {
	return &AudienceInsightsCallOptions{
		GenerateInsightsFinderReport: []gax.CallOption{
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
		ListAudienceInsightsAttributes: []gax.CallOption{
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
		ListInsightsEligibleDates: []gax.CallOption{
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
		GenerateAudienceCompositionInsights: []gax.CallOption{
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

// internalAudienceInsightsClient is an interface that defines the methods available from Google Ads API.
type internalAudienceInsightsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GenerateInsightsFinderReport(context.Context, *servicespb.GenerateInsightsFinderReportRequest, ...gax.CallOption) (*servicespb.GenerateInsightsFinderReportResponse, error)
	ListAudienceInsightsAttributes(context.Context, *servicespb.ListAudienceInsightsAttributesRequest, ...gax.CallOption) (*servicespb.ListAudienceInsightsAttributesResponse, error)
	ListInsightsEligibleDates(context.Context, *servicespb.ListInsightsEligibleDatesRequest, ...gax.CallOption) (*servicespb.ListInsightsEligibleDatesResponse, error)
	GenerateAudienceCompositionInsights(context.Context, *servicespb.GenerateAudienceCompositionInsightsRequest, ...gax.CallOption) (*servicespb.GenerateAudienceCompositionInsightsResponse, error)
}

// AudienceInsightsClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Audience Insights Service helps users find information about groups of
// people and how they can be reached with Google Ads. Accessible to
// allowlisted customers only.
type AudienceInsightsClient struct {
	// The internal transport-dependent client.
	internalClient internalAudienceInsightsClient

	// The call options for this service.
	CallOptions *AudienceInsightsCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AudienceInsightsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AudienceInsightsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AudienceInsightsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GenerateInsightsFinderReport creates a saved report that can be viewed in the Insights Finder tool.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
func (c *AudienceInsightsClient) GenerateInsightsFinderReport(ctx context.Context, req *servicespb.GenerateInsightsFinderReportRequest, opts ...gax.CallOption) (*servicespb.GenerateInsightsFinderReportResponse, error) {
	return c.internalClient.GenerateInsightsFinderReport(ctx, req, opts...)
}

// ListAudienceInsightsAttributes searches for audience attributes that can be used to generate insights.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
func (c *AudienceInsightsClient) ListAudienceInsightsAttributes(ctx context.Context, req *servicespb.ListAudienceInsightsAttributesRequest, opts ...gax.CallOption) (*servicespb.ListAudienceInsightsAttributesResponse, error) {
	return c.internalClient.ListAudienceInsightsAttributes(ctx, req, opts...)
}

// ListInsightsEligibleDates lists date ranges for which audience insights data can be requested.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
func (c *AudienceInsightsClient) ListInsightsEligibleDates(ctx context.Context, req *servicespb.ListInsightsEligibleDatesRequest, opts ...gax.CallOption) (*servicespb.ListInsightsEligibleDatesResponse, error) {
	return c.internalClient.ListInsightsEligibleDates(ctx, req, opts...)
}

// GenerateAudienceCompositionInsights returns a collection of attributes that are represented in an audience of
// interest, with metrics that compare each attribute’s share of the audience
// with its share of a baseline audience.
//
// List of thrown errors:
// AudienceInsightsError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
func (c *AudienceInsightsClient) GenerateAudienceCompositionInsights(ctx context.Context, req *servicespb.GenerateAudienceCompositionInsightsRequest, opts ...gax.CallOption) (*servicespb.GenerateAudienceCompositionInsightsResponse, error) {
	return c.internalClient.GenerateAudienceCompositionInsights(ctx, req, opts...)
}

// audienceInsightsGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type audienceInsightsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AudienceInsightsClient
	CallOptions **AudienceInsightsCallOptions

	// The gRPC API client.
	audienceInsightsClient servicespb.AudienceInsightsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewAudienceInsightsClient creates a new audience insights service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Audience Insights Service helps users find information about groups of
// people and how they can be reached with Google Ads. Accessible to
// allowlisted customers only.
func NewAudienceInsightsClient(ctx context.Context, opts ...option.ClientOption) (*AudienceInsightsClient, error) {
	clientOpts := defaultAudienceInsightsGRPCClientOptions()
	if newAudienceInsightsClientHook != nil {
		hookOpts, err := newAudienceInsightsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AudienceInsightsClient{CallOptions: defaultAudienceInsightsCallOptions()}

	c := &audienceInsightsGRPCClient{
		connPool:    connPool,
		audienceInsightsClient: servicespb.NewAudienceInsightsServiceClient(connPool),
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
func (c *audienceInsightsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *audienceInsightsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *audienceInsightsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *audienceInsightsGRPCClient) GenerateInsightsFinderReport(ctx context.Context, req *servicespb.GenerateInsightsFinderReportRequest, opts ...gax.CallOption) (*servicespb.GenerateInsightsFinderReportResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateInsightsFinderReport[0:len((*c.CallOptions).GenerateInsightsFinderReport):len((*c.CallOptions).GenerateInsightsFinderReport)], opts...)
	var resp *servicespb.GenerateInsightsFinderReportResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.audienceInsightsClient.GenerateInsightsFinderReport(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *audienceInsightsGRPCClient) ListAudienceInsightsAttributes(ctx context.Context, req *servicespb.ListAudienceInsightsAttributesRequest, opts ...gax.CallOption) (*servicespb.ListAudienceInsightsAttributesResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListAudienceInsightsAttributes[0:len((*c.CallOptions).ListAudienceInsightsAttributes):len((*c.CallOptions).ListAudienceInsightsAttributes)], opts...)
	var resp *servicespb.ListAudienceInsightsAttributesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.audienceInsightsClient.ListAudienceInsightsAttributes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *audienceInsightsGRPCClient) ListInsightsEligibleDates(ctx context.Context, req *servicespb.ListInsightsEligibleDatesRequest, opts ...gax.CallOption) (*servicespb.ListInsightsEligibleDatesResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListInsightsEligibleDates[0:len((*c.CallOptions).ListInsightsEligibleDates):len((*c.CallOptions).ListInsightsEligibleDates)], opts...)
	var resp *servicespb.ListInsightsEligibleDatesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.audienceInsightsClient.ListInsightsEligibleDates(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *audienceInsightsGRPCClient) GenerateAudienceCompositionInsights(ctx context.Context, req *servicespb.GenerateAudienceCompositionInsightsRequest, opts ...gax.CallOption) (*servicespb.GenerateAudienceCompositionInsightsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateAudienceCompositionInsights[0:len((*c.CallOptions).GenerateAudienceCompositionInsights):len((*c.CallOptions).GenerateAudienceCompositionInsights)], opts...)
	var resp *servicespb.GenerateAudienceCompositionInsightsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.audienceInsightsClient.GenerateAudienceCompositionInsights(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
