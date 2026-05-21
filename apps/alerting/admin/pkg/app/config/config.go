package config

import "github.com/grafana/grafana-app-sdk/resource"

// RuntimeConfig carries the in-process dependencies the admin app needs at
// runtime to satisfy its custom routes. The /status aggregate route fans out
// across status kinds via the apiserver client, so it needs a ClientGenerator
// supplied by the host process at registration time.
type RuntimeConfig struct {
	ClientGenerator resource.ClientGenerator
}
