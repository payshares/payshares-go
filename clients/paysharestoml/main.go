package paysharestoml

import "net/http"

// PaysharesTomlMaxSize is the maximum size of payshares.toml file
const PaysharesTomlMaxSize = 5 * 1024

// WellKnownPath represents the url path at which the payshares.toml file should
// exist to conform to the federation protocol.
const WellKnownPath = "/.well-known/payshares.toml"

// DefaultClient is a default client using the default parameters
var DefaultClient = &Client{HTTP: http.DefaultClient}

// Client represents a client that is capable of resolving a Payshares.toml file
// using the internet.
type Client struct {
	// HTTP is the http client used when resolving a Payshares.toml file
	HTTP HTTP

	// UseHTTP forces the client to resolve against servers using plain HTTP.
	// Useful for debugging.
	UseHTTP bool
}

// HTTP represents the http client that a stellertoml resolver uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// Response represents the results of successfully resolving a payshares.toml file
type Response struct {
	AuthServer       string `toml:"AUTH_SERVER"`
	FederationServer string `toml:"FEDERATION_SERVER"`
	EncryptionKey    string `toml:"ENCRYPTION_KEY"`
	SigningKey       string `toml:"SIGNING_KEY"`
}

// GetPaysharesToml returns payshares.toml file for a given domain
func GetPaysharesToml(domain string) (*Response, error) {
	return DefaultClient.GetPaysharesToml(domain)
}

// GetPaysharesTomlByAddress returns payshares.toml file of a domain fetched from a
// given address
func GetPaysharesTomlByAddress(addy string) (*Response, error) {
	return DefaultClient.GetPaysharesTomlByAddress(addy)
}
