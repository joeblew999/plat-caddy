// plat-caddy: Custom Caddy build with commonly used modules
//
// This builds Caddy with selected plugins for the plat-* ecosystem:
// - caddy-dns/cloudflare: DNS challenge for Let's Encrypt with Cloudflare
// - caddy-security: Authentication and authorization
// - caddy-git: Git integration for auto-deploy
// - caddy-cache: Response caching
//
// Build: go build -o caddy ./cmd/caddy
// Run:   ./caddy run --config Caddyfile

package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// Standard Caddy modules
	_ "github.com/caddyserver/caddy/v2/modules/standard"

	// DNS providers for ACME challenges
	_ "github.com/caddy-dns/cloudflare"

	// Security plugins (optional - uncomment as needed)
	// _ "github.com/greenpau/caddy-security"

	// Additional modules can be added here
	// _ "github.com/caddyserver/cache-handler"
)

func main() {
	caddycmd.Main()
}
