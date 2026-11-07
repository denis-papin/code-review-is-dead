---
name: guide-spec-definition-md
description: Write or review a Markdown specification for a feature, fix, or bug before implementation.
---

# Feature and bug specifications

Use the location requested by the user. Otherwise inspect nearby specifications, repository guidance and links from the project documentation to find the established location and format. Preserve existing identifiers, metadata, lifecycle terms and links. A project without an identifier scheme can refer to a spec by title.

For a feature or fix, describe the user’s goal, the gap in current behavior, accepted inputs, observable outputs, business rules and failure outcomes. Give worked examples with concrete inputs and expected results, including a boundary case. Link rules to acceptance scenarios using the project’s conventions. Distinguish proposed scope from decisions that already have an owner.

For a bug, record reproduction conditions, observed and expected results, impact and evidence. Label a suspected cause as a hypothesis until verified. Identify the regression scenario that would distinguish the fix from current behavior.

When revising a spec, retain established rules and case identifiers unless the requested change supersedes them. Check examples against the rules and links against the actual files. Keep implementation notes separate from required behavior. Deliver the specification within the requested scope.
