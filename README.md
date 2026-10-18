# okt-api : Open Korean Text API server

## Building and Running Locally

Build image:
```sh
make build_image IMAGE_TAG=okt-api:test BUILD_ARCHS=linux/amd64 BUILD_FLAGS=
```
> Change `BUILD_ARCHS` to `linux/arm64` for compatibility with Apple M series processors

Run the API server (token via env):
```
docker run -it --rm -p 8080:80 -e OKT_APIKEY=your-secret-apikey okt-api:test
```

Or with docker compose:
```
OKT_APIKEY=your-secret-apikey docker compose up --build
```

Browse API documents:
- http://localhost:8080/v1/docs

## References
- [Open Korean Text](https://github.com/open-korean-text): GitHub repositories for the Open Korean Text project