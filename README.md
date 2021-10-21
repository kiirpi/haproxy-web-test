# haproxy test - retries

https://www.haproxy.com/blog/haproxy-layer-7-retries-and-chaos-engineering/

## Project build

```shell
GOOS=linux GOARCH=amd64 go build -o haproxy-test-app
```

## Docker build

```shell
docker build --tag haproxy-test-app .
```


