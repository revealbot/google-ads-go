# google-ads-go

| Google Ads API version 	| [`v0_12_0`](https://developers.google.com/google-ads/api/docs/release-notes#070_2019-01-30) |
|-|:-:|
| Build | [![CircleCI](https://circleci.com/gh/revealbot/google-ads-go.svg?style=shield)](https://circleci.com/gh/revealbot/google-ads-go) |
| Release | ![Release](https://img.shields.io/github/release/revealbot/google-ads-go.svg) |

## Features
- Fully matches the latest [Google Ads API Reference](https://developers.google.com/google-ads/api/reference/rpc/)
- Implemented via [gRPC](https://grpc.io/) with [Protocol Buffers](https://developers.google.com/protocol-buffers/)
- CLI utils and API for generating and refreshing Google Ads credentials

## Installation
To install, simply run:
```bash
$ go get -d github.com/revealbot/google-ads-go
```
Make sure your PATH includes the $GOPATH/bin directory if you want to use the [CLI utils](https://github.com/revealbot/google-ads-go#cli):
```bash
export PATH=$PATH:$GOPATH/bin
````

## Example
```go
package main

import (
  "fmt"

  "github.com/revealbot/google-ads-go/ads"
  "github.com/revealbot/google-ads-go/services"
)

func main() {
  // Create a client from credentials file
  client, err := ads.NewClientFromStorage("google-ads.json")
  if err != nil {
    panic(err)
  }
  
  // Load the "GoogleAds" service
  googleAdsService := services.NewGoogleAdsServiceClient(client.Conn())
  
  // Create a search request
  request := services.SearchGoogleAdsRequest{
    CustomerId: "2984242032",
    Query:      "SELECT campaign.id, campaign.name FROM campaign ORDER BY campaign.id",
  }
  
  // Get the results
  response, err := googleAdsService.Search(client.Context(), &request)
  for _, row := range response.Results {
    campaign := row.Campaign
    fmt.Printf("id: %d, name: %s\n", campaign.Id.Value, campaign.Name.Value)
  }
}
```

When using the `NewGoogleAdsClientFromStorage` method, you must provide a path to a valid `google-ads.json` file (containing your Google Ads API credentials).
```json
{
    "client_id": "YOUR_CLIENT_ID",
    "client_secret": "YOUR_CLIENT_SECRET",
    "refresh_token": "YOUR_REFRESH_TOKEN",
    "developer_token": "YOUR_DEVELOPER_TOKEN"
}

```

## CLI
This library also provides some CLI utilities for generating/refreshing Google OAuth2 credentials. The newly generated token will be printed to stdout, as well as the expiry timestamp.
#### Generate Access Token from Refresh token
```bash
$ gads refresh -client-id CLIENT_ID -client-secret CLIENT_SECRET -refresh-token REFRESH_TOKEN
```
#### Generate Access Token from JSON credentials file
```bash
$ gads refresh -file credentials.json
```
Additionally, you can use the `-help` flag for more information: `gads refresh --help`

## Contributing
### Compiling
All build scripts use `Makefile`

**Build and run**
```bash
$ make run
```

**Run with gRPC debugging output**
```bash
$ make run-debug
```

## Changelog
To see the changes between Google Ads API versions, take a look at the official [Google Ads API Release Notes](https://developers.google.com/google-ads/api/docs/release-notes).

### Manually Building Protos
Building `.pb.go` files from the original `googleads` protos should only be done when updating to a new Google Ads version.

Requirements:
- [protoc](https://github.com/protocolbuffers/protobuf)

Build `.pb.go` protos:
```bash
$ make protos
```

### Generate Google Ads API protos
Make sure that `protoc-gen-go-grpc` and `protoc-gen-go` and available in PATH.
Update previous version to new version in Makefile and shell files, for example, change v18 to v19.
Then do the following:
```bash
make clone-googleapis
make update-googleapis
make protos
```
Script should be executed without errors.
After that rename `resources/experiment_arm.pb.go` to `resources/experiment_arm0.pb.go`, add all new files from common, enums, errors, resources, services. Then commit changes.
