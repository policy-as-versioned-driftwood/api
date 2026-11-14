# api

The api team's app. Real-estate epic, ticket 08 — the good citizen: current Go dependencies (a
single small real router, `go-chi/chi`, latest at time of writing — see `go.mod`), policy version
`2.2.0`. The contrast with ledger's Log4Shell-era log4j is the point: same design, same gates,
very different dependency hygiene, both visible on the estate dashboard.

`k8s/` carries this team's workload manifest: `mycompany.com/policy-version: "2.2.0"` and its
`department` label, this team's own adoption decision.

## Release

Push a `vX.Y.Z` tag; `.github/workflows/release.yml` builds and publishes
`ghcr.io/policy-as-versioned-flux/api:vX.Y.Z` and prints the digest in the run summary.

## Correction, 2026-10-03 (eco-system ticket 154)

The historical description "good citizen" describes policy declarations, not vulnerability health. The 2026-09-25 scan found eight HIGH CVEs in Go 1.26.5 and an end-of-support Alpine 3.20.10 base. New inventory records grade the served image digest; the label does not assert a clean scan.
