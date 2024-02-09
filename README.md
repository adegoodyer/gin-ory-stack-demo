# gin-ory-stack-demo
- [gin-ory-stack-demo](#gin-ory-stack-demo)
  - [Overview](#overview)
    - [Essential Docs](#essential-docs)
  - [Project](#project)
    - [Features](#features)
    - [Goals](#goals)
    - [Future Iterations](#future-iterations)
  - [Commands](#commands)

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
- [ ] containerise app (Docker, docker-compose or k8s)
- [x] export into Gitlab
- [ ] pipeline - container building (Gitlab action)

### Future Iterations
- [ ] MFA
- [ ] Admin panel
- [ ] Cobra CLI (instead of single flag processing)

## Commands
```bash
task              # list tasks
task app:build:   # build app
task app:run      # run app
```