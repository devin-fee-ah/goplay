# api

Proof of concept demonstrating:
- aws (`aws-sdk`)
- di (`fx`)
- env (`dotenv`)
- logging (`zap`)
- orm (`ent`)
- swagger (`swaggo`)
- web-framework (`gin`)

Also:
- uses Go 1.15
- is a Go module
- bundles into a Docker image (see `make`)

### Usage
Run `make` and you'll see the commands.

As a quick start, `make run`.

### Etc.
Worthwhile to review, but not approached:
- vendoring dependencies
- custom codegen (via go generate)
- swapping the `gin` web-framework for `fastapi`
