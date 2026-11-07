---
name: guide-integration-test-md
description: Write or review Markdown specifications for integration and unit tests. Connect fixtures, case identifiers and assertions to the business contract and the chosen test boundary.
---

# Integration and unit test specifications

Locate the feature contract, existing test specifications, the component under test and its test harness. Follow the project’s naming, case numbering and location conventions. Preserve case identifiers when revising a specification and map each case to its executable test when one exists.

Identify the boundary under test. For HTTP integration tests, exercise a public route and assert on the response and externally observable effects. For library integration tests, call the library’s public API. For unit tests, identify the unit and its collaborators.

State the coverage goal and the rule each case proves. Specify fixture setup, action and exact expected outcome. Include invalid input, boundaries and relevant failure paths. Distinguish fixture traffic from the operation under test. Use deterministic dependencies where the scenario requires control of time, data or external failures.

When test code exists, verify names, fixtures and assertions against the approved contract. Report mismatches without rewriting the contract to match an implementation defect. Identify planned tests separately from observed test results. For a documentation request, edit only the requested specification.
