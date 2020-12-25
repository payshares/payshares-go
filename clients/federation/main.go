package federation

import (
	"net/http"

	"github.com/payshares/go/clients/horizon"
	"github.com/payshares/go/clients/paysharestoml"
)

// FederationResponseMaxSize is the maximum size of response from a federation server
const FederationResponseMaxSize = 100 * 1024

// DefaultTestNetClient is a default federation client for testnet
var DefaultTestNetClient = &Client{
	HTTP:        http.DefaultClient,
	Horizon:     horizon.DefaultTestNetClient,
	PaysharesTOML: paysharestoml.DefaultClient,
}

// DefaultPublicNetClient is a default federation client for oubnet
var DefaultPublicNetClient = &Client{
	HTTP:        http.DefaultClient,
	Horizon:     horizon.DefaultPublicNetClient,
	PaysharesTOML: paysharestoml.DefaultClient,
}

// Client represents a client that is capable of resolving a Payshares.toml file
// using the internet.
type Client struct {
	PaysharesTOML PaysharesTOML
	HTTP        HTTP
	Horizon     Horizon
	AllowHTTP   bool
}

// Horizon represents a horizon client that can be consulted for data when
// needed as part of the federation protocol
type Horizon interface {
	HomeDomainForAccount(aid string) (string, error)
}

// HTTP represents the http client that a federation client uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// PaysharesTOML represents a client that can resolve a given domain name to
// payshares.toml file.  The response is used to find the federation server that a
// query should be made against.
type PaysharesTOML interface {
	GetPaysharesToml(domain string) (*paysharestoml.Response, error)
}

// confirm interface conformity
var _ PaysharesTOML = paysharestoml.DefaultClient
var _ HTTP = http.DefaultClient
var _ Horizon = horizon.DefaultTestNetClient
