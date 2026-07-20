# ~/.zshenv — loaded by every zsh shell (interactive, non-interactive, scripts).
# Keep this file minimal and fast: PATH + env only. No plugins, prompts, or aliases.

# Load shared environment variables / secrets if present.
[ -f "$HOME/.config/secrets/ai.env" ] && source "$HOME/.config/secrets/ai.env"

# Android SDK
export ANDROID_HOME="$HOME/Android/Sdk"

# Universal PATH additions (prepend user tools; go/cli-tools, grok, bun, local, Android, CUDA).
export PATH="$HOME/go/bin:$HOME/.grok/bin:$HOME/.bun/bin:$HOME/.local/bin:$ANDROID_HOME/emulator:$ANDROID_HOME/platform-tools:/opt/cuda/bin:$PATH"

export OPENCODE_ENABLE_EXA=true

# >>> Claude Code via CLIProxyAPI (Kimi K3 + xAI Grok) >>>
export ANTHROPIC_DEFAULT_OPUS_MODEL=kimi-k3
export ANTHROPIC_DEFAULT_SONNET_MODEL=grok-4.5(high)
export ANTHROPIC_DEFAULT_HAIKU_MODEL=grok-composer-2.5-fast(high)
# Claude Code hardcodes 200k for non-claude-* model IDs and does not use
# CLIProxyAPI's max_input_tokens. This override applies to all non-claude models
# (e.g. grok-4.5=500k). Composer 2.5 is only 200k upstream — switch this down
# if you primarily use that alias.
export CLAUDE_CODE_MAX_CONTEXT_TOKENS=500000
# <<< Claude Code via CLIProxyAPI (Kimi K3 + xAI Grok) <<<

# Vite+ bin (https://viteplus.dev)
. "$HOME/.vite-plus/env"
