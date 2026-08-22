# ~/.zshenv — loaded by every zsh shell (interactive, non-interactive, scripts).
# Keep this file minimal and fast: PATH + env only. No plugins, prompts, or aliases.

# Load shared environment variables / secrets if present.
[ -f "$HOME/.config/secrets/ai.env" ] && source "$HOME/.config/secrets/ai.env"

# Android SDK
export ANDROID_HOME="$HOME/Android/Sdk"

# Universal PATH additions (prepend user tools; go/cli-tools, grok, bun, local, Android, CUDA).
export PATH="$HOME/go/bin:$HOME/.grok/bin:$HOME/.bun/bin:$HOME/.local/bin:$ANDROID_HOME/emulator:$ANDROID_HOME/platform-tools:/opt/cuda/bin:$PATH"

export OPENCODE_ENABLE_EXA=true

# >>> Claude Code via CLIProxyAPI (xAI Grok, effort by tier) >>>
export ANTHROPIC_BASE_URL=http://192.168.1.91:8317
export ANTHROPIC_AUTH_TOKEN=sk-iKqEhkE7miMYmnkxP
# (effort) suffix sets reasoning effort for that alias.
export ANTHROPIC_DEFAULT_OPUS_MODEL=grok-4.6(high)
export ANTHROPIC_DEFAULT_SONNET_MODEL=grok-4.5(high)
export ANTHROPIC_DEFAULT_HAIKU_MODEL=grok-composer-2.5-fast(high)
# Claude Code hardcodes 200k for non-claude-* model IDs and does not use
# CLIProxyAPI's max_input_tokens. Grok 4.5/4.6 are 500k upstream.
export CLAUDE_CODE_MAX_CONTEXT_TOKENS=500000
# <<< Claude Code via CLIProxyAPI (xAI Grok, effort by tier) <<<

# Vite+ bin (https://viteplus.dev)
. "$HOME/.vite-plus/env"
