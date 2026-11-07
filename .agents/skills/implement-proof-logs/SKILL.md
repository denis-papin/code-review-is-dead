---
name: implement-proof-logs
description: Implement or audit structured logging that demonstrates execution of a specified feature. Trace decisions, outcomes and failures using the project's logging conventions.
---

# Proof of execution through logs

Read the feature contract, associated acceptance scenarios and the repository’s logging guidance if available. Inspect nearby logging calls to determine the logger, levels, event fields and correlation mechanism. Preserve the existing encoding of these fields.

Map each relevant execution path to the following placements. For each placement, identify the spec rule and code site, or explain why that placement does not apply.

| Placement | Evidence |
| --- | --- |
| Entry | The operation started under a specific execution context. |
| Exit | The operation completed with a success or failure outcome. |
| External call | An outbound request and its observed result. |
| Business decision | The selected branch and reason. |
| State change | The confirmed change after persistence succeeds. |
| Rule check | The rule evaluated and its result, including success. |
| Milestone | An observable business step completed. |
| Error | The failure category and context at detection or handling. |

Use the existing operation and request identifiers to distinguish concurrent calls, and carry an entity identifier once available. Use established severity levels for expected events, handled failures and unexpected errors. Keep secrets out of events. Logging arguments must have no side effects. Limit edits to logging and preserve the behavior under test.

Run the applicable build and tests, then exercise representative success and failure scenarios. Group observed events by execution and verify their order, reached placements and exactly one entry and exit per completed invocation. Report sample evidence, uncovered placements and any correlation limits. Keep specs and tests read-only during logging changes.
