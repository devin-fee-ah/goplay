# goplay

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

## Usage
Run `make` and you'll see the commands.

As a quick start, `make run`.

## Etc.
Worthwhile to review, but not approached:
- vendoring dependencies
- custom codegen (via go generate)
- swapping the `gin` web-framework for `fastapi`

## Benchmarks

### Write performance
To test, I sent a POST request to create users (env=production).
These users were written to disk (demonstrating consistent I/O) by sqlite.

You can run these on your own machine:

`make run &`
`wrk -t12 -c400 -d30s -s bench.lua http://localhost:8080/api/v1/users/`

```
Running 30s test @ http://localhost:8080/api/v1/users/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   105.75ms  110.13ms 883.36ms   78.91%
    Req/Sec   397.20    132.42     0.98k    69.81%
  142372 requests in 30.08s, 28.04MB read
  Socket errors: connect 0, read 251, write 0, timeout 0
  Non-2xx or 3xx responses: 121873
Requests/sec:   4733.48
Transfer/sec:      0.93MB
```

### Read performance
To test, I sent a POST request to create users (env=production).
This reads users written to disk (demonstrating consistent I/O) by sqlite.
(Note: if there are no users, you'll just be getting an empty list)

You can run these on your own machine:

`make run &`
`wrk -t12 -c400 -d30s http://localhost:8080/api/v1/users/`

```
Running 30s test @ http://localhost:8080/api/v1/users/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    61.11ms   73.27ms 677.54ms   79.59%
    Req/Sec     1.19k   319.61     2.50k    69.28%
  426295 requests in 30.07s, 54.48MB read
  Socket errors: connect 0, read 253, write 0, timeout 0
Requests/sec:  14175.24
Transfer/sec:      1.81MB
`