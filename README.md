# MyIP
This Go program sets up a simple HTTP server that responds with the client's public IP address in JSON format.

<img width="469" height="656" alt="image" src="https://github.com/user-attachments/assets/7941f316-c797-4836-ab10-88912cf817f9" />


### Dependencies
- Go 1.22 or higher

### Usage

#### Environment Variables
- `PORT`: (optional) The port where the application will listen. Default is 8080.

#### Run
```
PORT=8081 go run main.go
```

#### Run with Docker
```
docker build -t myip .
docker run -p 8080:8080 -e PORT=8080 myip
```

Or pull the published image:
```
docker run -p 8080:8080 dbohry/myip:latest
```

### License
[MIT](LICENSE)
