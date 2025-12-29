# plat-caddy

Custom Caddy build with commonly used modules for the plat-* ecosystem.

## Included Modules

- **caddy-dns/cloudflare** - DNS challenge for Let's Encrypt wildcards with Cloudflare

## Quick Start

```bash
# 1. Install the binary
xplat manifest install .

# 2. Configure (optional - for Cloudflare DNS challenge)
cp .env.example .env
# Edit .env with your Cloudflare API token

# 3. Run
caddy run --config Caddyfile
```

## Build from Source

```bash
go build -o caddy ./cmd/caddy
```

## Usage

```bash
# Run with config
caddy run --config Caddyfile

# Reload config
caddy reload --config Caddyfile

# Validate config
caddy validate --config Caddyfile

# Format Caddyfile
caddy fmt --overwrite Caddyfile
```

## Cloudflare DNS Challenge

For wildcard certificates, set up Cloudflare DNS challenge:

1. Create API token at https://dash.cloudflare.com/profile/api-tokens
2. Permissions: Zone:DNS:Edit
3. Add to .env: `CLOUDFLARE_API_TOKEN=your_token`
4. Use in Caddyfile:
   ```
   {
       acme_dns cloudflare {env.CLOUDFLARE_API_TOKEN}
   }

   *.example.com {
       tls {
           dns cloudflare {env.CLOUDFLARE_API_TOKEN}
       }
       ...
   }
   ```

## Core Infrastructure

This is a core infrastructure package required by other plat-* systems.
Install via: `xplat setup` or `xplat manifest install .`
