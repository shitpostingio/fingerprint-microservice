# Fingerprint Microservice

REST-powered microservice for analyzing media and returning feature-based information.
The repository also contains a client to perform requests to the service.

It is important to notice that this service provides no support for authentication or caching. It is also completely stateless, making it ideal to be used in the backend. A possible "frontend" implementation can be found in [Analysis API](https://gitlab.com/shitposting/analysis-api).

## Endpoints

- Image endpoint: `<bind-address>/fingerprinting/image`
- Video endpoint: `<bind-address>/fingerprinting/video`
- Health check: `<bind-address>/healthy`

## Returned data

The data returned by the server is in the form:

```go
type Analysis struct {
    Fingerprint            FingerprintResponse
    NSFW                   NSFWResponse
    FingerprintErrorString string
    NSFWErrorString        string
}

```

The client trims off the unnecessary data and returns:

```go
type FingerprintResponse struct {
    PHash     string
    Histogram []float64
}
```

## Environment options

- Service bind address and port: `FP_BIND_ADDRESS` (defaults to `localhost:10000`).
- Max size for image files: `FP_MAX_IMAGE_SIZE` (defaults to `10 << 20`, 10 MB).
- Max size for video files: `FP_MAX_VIDEO_SIZE` (defaults to `20 << 20`, 20 MB).
