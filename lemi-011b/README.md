[![Go Report Card](https://goreportcard.com/badge/github.com/sss-eda/lemi-011b)](https://goreportcard.com/report/github.com/sss-eda/lemi-011b)

# LEMI-011B
This repository contains acquisition software for the LEMI-011B magnetometer.

The software consists of a server and client side. The idea is that the client service runs on the logging cluster. It reads data from the serial port and forwards it to the server API. The server, in turn, runs remotely and listens on the API for new data. The data is then persisted by the server side software.

## Docker
### Build
Instructions for building the Docker container images are show below:

- To build the **server** container image:
```bash
$ docker build -t lemi011b-server:latest -f build/docker/server/Dockerfile .
```
- To build the **client** container image:
```bash
$ docker build -t lemi011b-client:latest -f build/docker/client/Dockerfile .
```

### Run
Example instructions for running the Docker containers are show below:

- To run a **server** container with dependencies:
```bash
# Start a timescaledb instance.
$ docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb:latest-pg12
# Run the server
$ docker run -d -p 8080:8080 -e LEMI011B_SERVER_TIMESCALEDB_URL="postgres://postgres:password@192.168.0.1:5432/lemi011b" lemi011b-server
```
- To run a **client** container:
```bash
# Run the client and mount the serial port into the container.
$ docker run --privileged -d -e LEMI011B_CLIENT_REST_URL="http://192.168.0.1:8080" -v /dev/ttyUSB0:/dev/ttyUSB0 lemi011b-client
```

## TODO
- [ ] Implement tests for all packages
- [ ] Fix problem with new devices
- [ ] Add device ID functionality
