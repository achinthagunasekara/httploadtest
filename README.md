# HTTP Load Test

Simple load testing tool for web sites written in Go.

## Useage

| Option | Description | Default | Required |
| ------ | ----------- | ------- | -------- |
| -url | URL to request for the load test | N/A | Yes |
| -timeout | wait for before timing out a request | 30s | No |
| -concurrent | number of concurrent tests to run at a given time | 1 | No |
| -duration | duration of the load test | 30s | No |
| -debug | run in debug mode. log everything. | false | No |

## Using a Proxy

Please set the following environment variables to use a proxy.

* http_proxy
* https_proxy
* HTTP_PROXY
* HTTPS_PROXY
