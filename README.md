# Go Client APIs for StackPath

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

You can set this in the context as described in the generated doc, exerpted here.

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```
