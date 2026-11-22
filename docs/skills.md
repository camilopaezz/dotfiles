# Shared skills

`files/claude/skills/` is the source of truth for personal skills in Claude Code, Codex, and Grok Build CLI. Each skill folder is symlinked into `~/.claude/skills/`, `~/.agents/skills/`, and `~/.grok/skills/`. Codex's built-in skills stay in `~/.codex/skills/.system/`.

Deploy only skills, without touching other dotfiles:

```sh
python deploy.py skills --dry-run
python deploy.py skills
```

The `link` and `all` commands also deploy skills. Existing skill folders move to the corresponding agent's `skill-backups/` directory before replacement. Repeated deployments preserve earlier backups. Backups stay outside skill discovery directories.

Edit skills in the repo. Folder links share changes immediately, including newly added references and metadata. Run the deploy command when adding a new skill folder. Start a new agent session if the current session has cached its skill list.

Claude and Grok use `/skill-name`; Codex uses `$skill-name`. A manual-only shared skill needs both `disable-model-invocation: true` in `SKILL.md` and this Codex sidecar in `agents/openai.yaml`:

```yaml
policy:
  allow_implicit_invocation: false
```

Deployment replaces matching skill names but leaves unrelated installed skills alone. The initial synchronization removed the old Codex-only skills from discovery, with backups in `~/.agents/skill-backups/`. Install new shared skills into this repo rather than maintaining a separate Codex set.

## Frontend design

`frontend-design` is vendored from [Anthropic's skill](https://github.com/anthropics/skills/tree/main/skills/frontend-design), with its instructions and Apache 2.0 license unchanged. It guides visual design for new interfaces and redesigns, including typography, layout, interface copy, and visual review.

Use `$frontend-design` in Codex or `/frontend-design` in Claude Code and Grok. The skill also allows automatic invocation for relevant UI work. All three agents use the same repo folder through the shared deployment above.

## Grok Build CLI

After deployment, run `grok inspect --json` to check discovery or open `/skills` in an interactive session. Grok also discovers `~/.agents/skills/` and Claude Code skills automatically. The explicit `~/.grok/skills/` links keep deployment available through Grok's own skill directory.

These links supply skills only. They do not install the CLI or register MCP servers referenced by a skill. Start a fresh Grok session after deployment.

Reference: [Grok skills and compatibility](https://docs.x.ai/build/features/skills-plugins-marketplaces).

## Hypruse in Codex

The skill supplies instructions. The MCP server supplies the desktop tools. Its registration lives in `files/codex/config.toml`, which is linked to `~/.codex/config.toml`.

Launch Codex inside the Hyprland session so it has the desktop environment variables. Open `/mcp` to check the connection, then try:

```text
$hypruse inspect my desktop without clicking or typing
```

This config uses inline MCP screenshots. The shared skill also handles file-based screenshots for other hosts. A standalone `hypruse` executable is not required when the MCP server runs through `uvx`.

Check local readiness without downloading packages:

```sh
uvx --offline hypruse doctor
codex mcp get hypruse
```

The offline command needs the existing cached Hypruse install. Open a fresh Codex session after changing MCP settings. Restart a persistent app-server only when its active sessions can be interrupted.

References: [Hypruse](https://github.com/IlyasKhallouki/hypruse), [Codex MCP configuration](https://learn.chatgpt.com/docs/extend/mcp?surface=cli).
