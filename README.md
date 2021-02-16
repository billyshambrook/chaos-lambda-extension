# Chaos proxy for AWS Lambda

POC using Lambda Extensions to run a "Forward Proxy" to inject failure into AWS Lambda functions.

## Usage

First build the extension and functions:

```bash
$ make build
```

You can then run the functions locally.

The GoFunction will inject 1000ms of latency to HTTP requests.

```bash
ENABLE_LAMBDA_EXTENSIONS_PREVIEW=1 sam local invoke -e - GoFunction
```

The PyFunction will inject 2000ms of latency to HTTP requests.

```bash
ENABLE_LAMBDA_EXTENSIONS_PREVIEW=1 sam local invoke -e - PyFunction
```

## Available failures

### Latency

Inject latency into requests by setting `CHAOS_LATENCY_MS`.
