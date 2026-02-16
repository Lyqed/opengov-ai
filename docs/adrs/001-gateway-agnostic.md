# ADR-001: Gateway-Agnostic Architecture

## Status: Accepted

## Context

The MCP gateway market is saturated with well-funded competitors (RunLayer $11M, Backslash $27M, Cloudflare free tier). Competing on gateway features is a losing strategy.

## Decision

OpenGov AI will NOT build its own MCP gateway/proxy. Instead, we will build an intelligence layer that integrates with existing gateways via adapters.

## Consequences

- We avoid direct competition with RunLayer, Cloudflare, MintMCP
- We can partner with gateway vendors instead of competing
- We need adapter development for each gateway (ongoing cost)
- Our value prop is differentiated: risk intelligence, not infrastructure
- Customers can switch gateways without losing OpenGov AI
