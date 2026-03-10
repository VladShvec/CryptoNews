# CryptoNews

CryptoNews is a microservice-based crypto news aggregation and processing platform written in Go.

The project collects crypto news from external sources, parses and normalizes articles, enriches and classifies them, and exposes processed data through an HTTP API.

## Architecture

The system is split into independent microservices with weak coupling.

### Services

- **api** — external HTTP API for clients
- **scanner** — schedules source scans and emits scan events
- **parser** — fetches and parses news from RSS / API / HTML sources
- **worker** — enriches, deduplicates, classifies, and processes parsed articles

### Infrastructure

- **nginx** — API Gateway / reverse proxy
- **kafka** — event bus for inter-service communication
- **postgres** — persistent storage for sources and articles
- **docker compose** — local development orchestration

## High-Level Flow

```text
scanner -> kafka -> parser -> kafka -> worker -> postgres -> api
                        ^
                        |
                     sources