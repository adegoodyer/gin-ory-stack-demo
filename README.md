# gin-ory-stack-demo

## Overview
- deep dive into Ory stack
  - Kratos      - headless authentication and user management
  - Hydra       - OAuth 2.0 and OpenID Connect server
  - Oathkeeper  - identity and access proxy (IAP)
  - Keto        - permission and authorization server
- [homepage](https://www.ory.sh/open-source/)
- [introduction](https://www.ory.sh/docs/kratos/ory-kratos-intro)
- [quick-start](https://www.ory.sh/docs/kratos/quickstart)

## Goals
- [x] deploy kratos into k8s cluster
- [x] implement basic Bootstrap 5 app to use for testing
- [ ] implement flows (excluding MFA for now) for browser server-side app using Gin

## Essential Docs
- [self-service flows](https://www.ory.sh/docs/kratos/self-service#browser-flows-for-server-side-apps-nodejs-php-java-)

## Project Setup
```bash
# init project
go mod init gin-ory-stack-demo

# get dependencies
go get -u \
github.com/gin-gonic/gin \
github.com/spf13/viper \
github.com/prometheus/client_golang/prometheus \
github.com/gin-contrib/sessions

# run locally
go run cmd/self-service-user-flows-browser-server-side/main.go

# http://127.0.0.1:8080/dashboard
```

## Development
```bash
# download modules
go mod download

# run locally
go run cmd/self-service-user-flows-browser-server-side/main.go

# http://127.0.0.1:8080/dashboard
```