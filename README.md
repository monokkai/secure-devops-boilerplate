# Secure DevOps Boilerplate

A **security-focused DevOps boilerplate** that demonstrates modern CI/CD with automated security checks and containerized deployment.

## Overview

This project is a template repository that combines:

- **NestJS** backend (Node.js)
- **Dockerized environment**
- **Security-first CI/CD pipeline**

It is designed to help developers and DevSecOps engineers quickly bootstrap a secure project with automated pipelines.

---

## Features

### 1. Application

- Simple backend built with **NestJS**
- Configurable `.env` for secrets (never commit real secrets)

### 2. Docker

- Dockerfile for the app
- `docker-compose.yml` for local development

### 3. CI/CD Pipeline (GitHub Actions)

The pipeline includes:

- **Linting**
- **Secret scanning** with [Gitleaks](https://github.com/gitleaks/gitleaks)
- **Dependency scanning** with [Trivy](https://github.com/aquasecurity/trivy)
- **Static Application Security Testing (SAST)** with [Semgrep](https://semgrep.dev/)
- **Deployment to a dev environment** (Docker-based)

### 4. Bash Scripts

- Local CI runner (`scripts/run_ci.sh`)
- Health check script
- Pre-commit checks

---

## Why This Project Is Useful

- Demonstrates **DevSecOps practices**
- Integrates **Security-as-Code** early in the development lifecycle
- Easy to extend into a full **security framework**

---

## Roadmap

- [ ] Add dynamic application security testing (DAST)
- [ ] Add Infrastructure-as-Code scanning (e.g., Checkov)
- [ ] Extend deployment to cloud (Kubernetes)

---

## How to Run Locally

```bash
git clone https://github.com/monokkai/secure-devops-boilerplate.git
cd secure-devops-boilerplate

cp .env.example .env

docker-compose up --build
```
