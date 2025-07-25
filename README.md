# StackPath Go SDK
## Go Client APIs for StackPath
## Authored by Kris Raney

This repo provides Go client libraries for StackPath’s CDN and Edge Compute APIs. It was created to facilitate the development of ultra-low-latency audio collaboration software during the COVID era, requiring tight control over container orchestration and geographically aware routing. This in turn was intended to support student music ensemble practice during at-home schooling.

This project showcases:
* Real-world use of OpenAPI codegen and multi-package Go SDKs
* Platform-level integration with CDN, DNS, SSL, WAF, Edge Compute, and Object Storage
* A reusable OAuth2 token source for developer convenience and testing (GoVCR-ready)
* The repo demonstrates fluency in cloud-native API integration, platform SDK design, and pragmatic tooling for external infrastructure control.

These APIs are generated automatically from StackPath's OpenAPI specs for their
APIs. (The one exception is their Monitoring (Legacy) API, which I skipped. I
didn't think there was much justification for writing new client code in a new
language but targeting a legacy interface.)

To recreate them or update them you will need [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

You can run `make clean` then `make`. There's no build of the libraries, since
`go get` will just pull source anyway.

Since StackPath uses a collection of separate APIs, these are generated as
separate packages.

Use these in your go project by running

```
go get gitlab.com/kraney/stackpath
```

or by getting an individual package, like

```
go get gitlab.com/kraney/stackpath/pkg/cdn
```

Detailed instructions for the APIs are in the generated output:

- [Accounts and Users](pkg/accounts_and_users/README.md)
- [Sites](pkg/sites/README.md)
- [SDN](pkg/cdn/README.md)
- [WAF](pkg/waf/README.md)
- [DNS](pkg/dns/README.md)
- [Edge Compute](pkg/edge_compute/README.md)
- [Edge Compute Networking](pkg/edge_compute_networking/README.md)
- [SSL](pkg/ssl/README.md)
- [Monitoring](pkg/monitoring/README.md)
- [Object Storage](pkg/object_storage/README.md)

## Authentication

Note that the automatically generated documentation doesn't really mention authorization on the individual APIs, but for most of the APIs you will need to set the request header `Authorization: Bearer <token>` as described in [StackPath's documentation](https://stackpath.dev/docs/stackpath-api-authentication).

The generated documentation on how to do this is a bit misleading. A
golang.org/x/oauth2.TokenSource implementation has been provided to make it
easier. This will automatically refresh the token when it expires.

Here's an example usage:

```golang
package main

import (
	"context"
	"flag"
	"log"

	spauth "github.com/kraney/stackpath/pkg/oauth2"
	"github.com/kraney/stackpath/pkg/stacks"
	"golang.org/x/oauth2"
)

func main() {
	clientid := flag.String("clientid", "", "client ID.")
	clientsecret := flag.String("secret", "", "client secret.")
	flag.Parse()

	var ts oauth2.TokenSource = spauth.NewTokenSource(*clientid, *clientsecret)
	auth := context.WithValue(context.Background(), stacks.ContextOAuth2, ts)
	stackclient := stacks.NewAPIClient(stacks.NewConfiguration())
    stacksresp, _, err := stackclient.StacksApi.GetStacks(auth).Execute()
	if err != nil {
		log.Fatalf("Error getting stacks list: %v", err)
	}
	log.Printf("%v", stacksresp.Results)
}
```

This TokenSource is also designed to facilitate use of [GoVCR](https://github.com/seborama/govcr) for testing.

```golang
    httpclient := govcr.NewVCR(testname, &govcr.VCRConfig{
		CassettePath:     "./testdata/recordings",
	}).Client
	var ts *spauth.TokenSource = spauth.NewTokenSource(*clientid, *clientsecret, spauth.HTTPClientOption(httpclient))
	auth := context.WithValue(context.Background(), stacks.ContextOAuth2, ts)
```
