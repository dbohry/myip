# MyIP
[![CI](https://github.com/dbohry/myip/actions/workflows/build-and-push.yml/badge.svg)](https://github.com/dbohry/myip/actions/workflows/build-and-push.yml)
[![Docker Image](https://img.shields.io/badge/docker-dbohry/myip-2496ED?logo=docker&logoColor=white)](https://hub.docker.com/r/dbohry/myip)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

This Go program sets up a simple HTTP server that responds with the client's public IP address in JSON format.

<img width="469" height="656" alt="image" src="https://github.com/user-attachments/assets/7941f316-c797-4836-ab10-88912cf817f9" />


### Dependencies
- Go 1.22 or higher

### Usage

#### Environment Variables
- `PORT`: (optional) The port where the application will listen. Default is 8080.

#### Run
```
git clone https://github.com/dbohry/myip.git
cd myip
PORT=8081 go run .
```

Then open http://localhost:8081 in a browser.

Because the real client IP comes from proxy headers (`CF-Connecting-IP`, `X-Forwarded-For`, `X-Real-IP`), when running locally behind no proxy the response shows your local address (e.g. `127.0.0.1` or `::1`) — that is expected. To test with a custom IP, send a header:

```
curl -H "X-Real-IP: 203.0.113.7" http://localhost:8081
```

#### Tests
```
go test ./...
```

#### Run with Docker
```
git clone https://github.com/dbohry/myip.git
cd myip
docker build -t myip .
docker run -p 8080:8080 -e PORT=8080 myip
```

Or pull the published image:
```
docker run -p 8080:8080 dbohry/myip:latest
```

### License
[MIT](LICENSE)
