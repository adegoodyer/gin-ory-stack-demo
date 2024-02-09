# gin-ory-stack-demo
- [gin-ory-stack-demo](#gin-ory-stack-demo)
  - [Overview](#overview)
    - [Essential Docs](#essential-docs)
  - [Project](#project)
    - [Features](#features)
    - [Goals](#goals)
    - [Future Iterations](#future-iterations)
  - [Launch App](#launch-app)
  - [Additional Commands](#additional-commands)

## Overview
- deep dive into Ory stack
  - Kratos      - headless authentication and user management
  - Hydra       - OAuth 2.0 and OpenID Connect server
  - Oathkeeper  - identity and access proxy (IAP)
  - Keto        - permission and authorization server
- [homepage](https://www.ory.sh/open-source/)
- [introduction](https://www.ory.sh/docs/kratos/ory-kratos-intro)
- [quick-start](https://www.ory.sh/docs/kratos/quickstart)

### Essential Docs
- [self-service flows](https://www.ory.sh/docs/kratos/self-service#browser-flows-for-server-side-apps-nodejs-php-java-)

## Project

### Features
- Ory stack
- Gin router/web framework
- viper configuration
- [Taskfiles](https://taskfile.dev/)

### Goals
- [x] deploy kratos into k8s cluster
- [x] implement basic Bootstrap 5 app to use for testing
- [ ] implement flows (excluding MFA for now) for browser server-side app using Gin
  - [ ] sign-in
  - [ ] sign-out
  - [ ] registration
  - [ ] account-verification
  - [ ] account-recovery
  - [ ] error
  - [ ] update-user-settings
- [x] containerise app (Docker, docker-compose or k8s)
- [x] export into Gitlab
- [ ] pipeline - container building (Gitlab action)

### Future Iterations
- [ ] MFA
- [ ] Admin panel
- [ ] Cobra CLI (instead of single flag processing)

## Launch App
```bash
# create config/development.yaml and config/production.yaml (see examples in dir)

# build container
task container:build

# run container
task container:run

# http://127.0.0.1:8080
```

## Additional Commands
```bash
task # list tasks and descriptions
task app:build
task app:run
task container:build
task container:delete
task container:run
```