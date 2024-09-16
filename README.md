
# Google Ads API Client Library for Golang

[![Go](https://github.com/Optable/google-ads-pb/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Optable/google-ads-pb/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/Optable/google-ads-pb?status.svg)](https://pkg.go.dev/github.com/Optable/google-ads-pb)
[![Go Report Card](https://goreportcard.com/badge/github.com/Optable/google-ads-pb)](https://goreportcard.com/report/github.com/Optable/google-ads-pb)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This library provides a Golang client for the [Google Ads API](https://developers.google.com/google-ads/api/docs/start). It's fully generated from the [googleapis](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads) repository. More information on the generation process can be found [here](https://github.com/Optable/google-ads-pb/blob/main/Makefile).

Although this project isn't official, we deem it as low-risk due to its maturity and our two years of using it in production. However, always consult the [sunset schedule](https://developers.google.com/google-ads/api/docs/sunset-dates) of the Google Ads API.

## Features

- Full support for Google Ads API.
- Source code generated from the official googleapis repository.
- Support for GRPC and HTTP calls using [protojson](https://google.golang.org/protobuf/encoding/protojson).
- Although we won't regularly update this based on the [official repository](https://github.com/googleapis/googleapis), we continually maintain it. Update frequency depends on our needs but usually occurs a month after the official release.

## Version support

| google-ads-pb         | Google Ads API | Sunset date           |
| -----------------     | -------------- | --------------------- |
| Not generated yet     | v18            | End of September 2025 |
| Not generated yet     | v17.1          | End of May 2025       |
| v1.2.0                | v17            | End of May 2025       |
| Skipped               | v16.1          | End of January 2025   |
| Skipped               | v16            | End of January 2025   |
| ~~v1.1.0~~ DEPRECATED | v15            | End of September 2024 |
| ~~v1.0.1~~ DEPRECATED | v14            | End of May 2024       |

## Requirements

- Go 1.21.
- Familiarize yourself with the [OAuth2 guide](https://developers.google.com/google-ads/api/docs/oauth/overview).
- If needed, obtain a [developer token](https://developers.google.com/google-ads/api/docs/first-call/dev-token).

## Installation

```bash
$ go get github.com/Optable/google-ads-pb
```

## Getting started

1. Set your environment variables.

```bash
$ export ACCESS_TOKEN=<your access token>
$ export DEVELOPER_TOKEN=<your developer token>
$ export CUSTOMER_ID=<your customer id>
```

2. Establish a GRPC connection.

```go
ctx := context.Background()

headers := metadata.Pairs(
  "authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"),
  "developer-token", os.Getenv("DEVELOPER_TOKEN"),
  "login-customer-id", os.Getenv("CUSTOMER_ID"),
)
ctx = metadata.NewOutgoingContext(ctx, headers)

cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
conn, err := grpc.Dial("googleads.googleapis.com:443", cred)
if err != nil {
  panic(err)
}
defer conn.Close()
```

3. Start making calls.

```go
customerServiceClient := services.NewCustomerServiceClient(conn)
accessibleCustomers, err := customerServiceClient.ListAccessibleCustomers(
  ctx,
  &services.ListAccessibleCustomersRequest{},
)
if err != nil {
  panic(err)
}

for _, customer := range accessibleCustomers.ResourceNames {
  fmt.Println("ResourceName: " + customer)
}
```

You can also make HTTP calls using [protojson](https://google.golang.org/protobuf/encoding/protojson), though it isn't recommended.

```go
req := services.ListAccessibleCustomersRequest{}
requestBody, err := protojson.Marshal(&req)
if err != nil {
  panic(err)
}
request, err := http.NewRequest("GET", "https://googleads.googleapis.com/v14/customers:listAccessibleCustomers", bytes.NewBuffer(requestBody))
if err != nil {
  panic(err)
}
header := make(http.Header)
header.Set("content-type", "application/json")
header.Set("authorization", os.Getenv("ACCESS_TOKEN"))
header.Set("developer-token", os.Getenv("DEVELOPER_TOKEN"))
request.Header = header
client := &http.Client{}
response, err := client.Do(request)
if err != nil {
  panic(err)
}
defer response.Body.Close()
var responseBody []byte
if responseBody, err = io.ReadAll(response.Body); err != nil {
  panic(err)
}
listAccessibleCustomersResponse := new(services.ListAccessibleCustomersResponse)
if err := protojson.Unmarshal(response, listAccessibleCustomersResponse); err != nil {
  panic(err)
}
for _, customer := range listAccessibleCustomersResponse.ResourceNames {
  fmt.Println("ResourceName: " + customer)
}
```

## Examples

See [clients/internal/snippets](https://github.com/Optable/google-ads-pb/tree/main/clients/internal/snippets).

## Note for the updating process

1. Find the version number in [this folder](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads).
2. Run `git checkout -b bump-to-v15` (change the version accordingly).
3. Run `make generate` to generate the code.
4. It will prompt you to enter the version number. Enter the version number you found in step 1, without the v. Example: For `v15`, enter `15`.
5. Once the process is done, it should have modified some files in `/clients` and `/protogen`
6. Modify this readme.md to add the version number to the table with the next tag (ref: https://developers.google.com/google-ads/api/docs/sunset-dates)
7. Create a PR with the changes against https://github.com/Optable/google-ads-pb/tree/main.
8. Create a tag with the PR merged in step 7. `git tag -a v1.1.0 -m "update to v15"`
9. Mark the pushed tag as `Latest release` in GitHub: https://github.com/Optable/google-ads-pb/releases/edit/v1.1.0
10. Publish the release on Golang `GOPROXY=proxy.golang.org go list -m github.com/Optable/google-ads-pb@v1.1.0`

## Related

Here are some related projects

- [Google APIs](https://github.com/googleapis/googleapis)
- [Google Ads API Client Library for PHP](https://github.com/googleads/google-ads-php)

## Contributing

As this project is fully protoc-generated, we aren't accepting pull requests. However, we value your feedback and suggestions, so feel free to open an issue.
