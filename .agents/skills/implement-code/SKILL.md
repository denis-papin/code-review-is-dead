---
name: implement-code
description: Implement a specified feature or bug fix and verify its business rules against the project's tests.
---

# Implement a specified change

Resolve the requested change by its supplied identifier, title or spec path. Read its business contract, linked acceptance cases and applicable repository guidance. Inspect the relevant entry points, test harness, build configuration and nearby implementations. A diagnostic report alone may leave expected behavior undefined. Identify the missing decision when it affects implementation.

Map each rule and failure path to an observable result. Implement the smallest change that satisfies the full contract through the public entry points, including supported existing inputs. For schema changes, follow the repository’s migration and compatibility conventions.

Treat the spec and existing tests as a fixed contract. Preserve assertions, fixtures and test selection. Create missing tests when called for by the specification or the user. Report contradictions between tests and rules with the affected case and evidence.

Run the targeted tests and the applicable repository verification commands. Report which rules were exercised, the actual outcomes, and any gaps. Update implementation status only if the project tracks it and its completion criteria are met.
