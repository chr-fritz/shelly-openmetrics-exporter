# shelly-openmetrics-exporter

This is an openmetrics exporter written in Go.
It's `/probe` and `/metrics` endpoint report in OpenTracing format.

## Features

Use this tool to fetch power readings from your Shellys via Prometheus (or any OpenTracing compatible agent).
It supports multiple power lines (like they are available on `Shelly 3EM` or `Shelly Plus 2PM`).

This exporter is compatible with the _First Generation Shelly Devices API_ and the _Second Generation Shelly Devices API_.
Authentication is not supported.

## Usage

Run this program and point prometheus to it.
Use `target` to define which shelly to query.
You may use an IP or a domain name.
You may configure a port like this `name:port`.

If your Shelly is password protected, use basic auth to send the username and password.
For backward compatibility, the query parameters `username` and `password` will also work.

## TODO

- [x] detect shelly generation
- [x] temperature
- [x] goreleaser
- [ ] tests
- [x] rename shelly-openmetrics-exporter
- [x] authentication

## Release

Push a new tag.
A GitHub Action will be triggered.
Goreleaser will do the rest.

## License

This software is licensed [under MIT license](/LICENSE).
