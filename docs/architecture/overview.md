# OpenGov AI Architecture Overview

## Design Principles

1. **Gateway-agnostic**: We integrate with any MCP gateway, not replace it
2. **Intelligence over infrastructure**: Risk scoring and blast radius, not proxy/policy
3. **24-hour time to value**: Deploy and see risk assessment within a day
4. **Board-ready output**: Translate technical risk into business impact (dollars, records, compliance)

## Component Architecture

```
[Gateways] → [Ingestion] → [NATS JetStream] → [Risk Engine] → [Dashboard]
                                    ↓
                              [Discovery] ← [Scanner]
                                    ↓
                              [Response] → [Slack/Jira/SIEM]
```

## Data Flow

1. **Ingestion**: Gateway adapters normalize events into unified schema
2. **Discovery**: Continuous scanning discovers new agents and MCP servers
3. **Event Bus**: NATS JetStream provides reliable event streaming
4. **Risk Engine**: Scores risk, calculates blast radius, maps to frameworks
5. **Response**: Executes playbooks when risk thresholds are breached
6. **Dashboard**: Real-time CISO view of risk posture and compliance

## Technology Choices

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Risk engine language | Rust | Performance-critical scoring at scale, memory safety |
| Service language | Go | Proven for infrastructure services, fast development |
| Event bus | NATS JetStream | Lightweight, reliable, built-in persistence |
| Dashboard | React + TypeScript | Standard enterprise frontend stack |
| Risk policy format | JSON + CEL | Human-readable with expressive conditions |
