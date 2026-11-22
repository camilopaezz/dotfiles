# Matt Pocock skills review

Compared on 2026-10-06 against [mattpocock/skills](https://github.com/mattpocock/skills/tree/6fd947921b935b7e1e69293a200400f0fdd5c15f), version 1.3.1, commit `6fd947921b935b7e1e69293a200400f0fdd5c15f`.

Added upstream `skills/engineering/retro` and `skills/engineering/pr` verbatim to `files/claude/skills/`, including Codex metadata and PR credits. Deployment links them into Claude Code, Codex, and Grok. `retro` is manual-only; `pr` can run automatically when writing a PR body. Both are available on the next turn, though a host that caches its skill list may need a fresh session.

## Existing skills

Ten existing folders match upstream skill names, and all ten have differences. The [full diff](mattpocock-skills-upstream.diff) runs from the installed copy to upstream and includes added and removed supporting files. Existing skills were left as installed.

- `codebase-design`: domain-document references change to `GLOSSARY.md`; other edits mostly replace em dashes and rephrase prose.
- `diagnosing-bugs`: reads `GLOSSARY.md`; removes the post-mortem architecture hand-off from phase 6, which now covers cleanup. Also revises punctuation in the skill and human-feedback script.
- `domain-modeling`: changes `CONTEXT.md` and `CONTEXT-MAP.md` to `GLOSSARY.md` and `GLOSSARY-MAP.md`, and renames `CONTEXT-FORMAT.md` to `GLOSSARY-FORMAT.md`. ADR references follow the new convention.
- `grill-me`: explicitly calls the Skill tool for `grilling`.
- `grill-with-docs`: explicitly calls the Skill tool for `grilling` and `domain-modeling`.
- `grilling`: illustrates a whole question round with two questions and a separator; also revises punctuation.
- `improve-codebase-architecture`: adopts glossary filenames and explicit Skill-tool calls, plus prose and HTML-report punctuation edits.
- `wait-what`: follows `GLOSSARY-MAP.md` when several glossaries exist; also revises its description.
- `wizard`: mostly punctuation edits in the skill and shell template.
- `writing-for-agents`: mostly prose and punctuation edits, including the invocation reference.

Upstream supplies `agents/openai.yaml` for all ten. Six folders lack that file locally; the four existing manual-only sidecars gain display names and short descriptions upstream, preserving their implicit-invocation policy.

The new `pr` skill uses `GLOSSARY.md`, while the installed domain-modeling skills still use `CONTEXT.md`. An eventual update should migrate those conventions together. The explicit Skill-tool wording is Claude-specific, so review it for Codex and Grok before replacing the shared copies.

`cli-tools`, `computer-use`, `hypruse`, and `unslop` have no same-name counterpart in this upstream checkout. They were excluded from the upstream comparison.

## Other upstream skills not installed

These names are absent locally, which does not mean every one was newly released.

- Engineering: `ask-matt`, `code-review`, `implement`, `implement-spec`, `prototype`, `research`, `setup-matt-pocock-skills`, `tdd`, `to-spec`, `to-tickets`, `triage`, `wayfinder`.
- Productivity: `handoff`, `teach`, `to-questionnaire`.
- Miscellaneous: `git-guardrails-claude-code`, `migrate-to-shoehorn`, `scaffold-exercises`, `setup-pre-commit`.
- In progress: `chief-of-staff`, `claude-handoff`, `loop-me`, `setup-ts-deep-modules`, `writing-beats`, `writing-fragments`, `writing-shape`.

The [1.3 changelog](https://github.com/mattpocock/skills/blob/6fd947921b935b7e1e69293a200400f0fdd5c15f/CHANGELOG.md) promotes `implement-spec`, `pr`, and `retro` into Engineering. Only the requested `pr` and `retro` were added.
