# ~/.zshenv — loaded by every zsh shell (interactive, non-interactive, scripts).
# Keep this file minimal and fast: PATH + env only. No plugins, prompts, or aliases.

# Load shared environment variables / secrets if present.
[ -f "$HOME/.config/secrets/ai.env" ] && source "$HOME/.config/secrets/ai.env"

# ai.env points Claude Code at CLIProxyAPI. Keep those coordinates under names
# Claude does not read. Bare `claude` stays first-party so /login works later.
# `kiro` and `hub` put ANTHROPIC_BASE_URL and ANTHROPIC_AUTH_TOKEN back.
if [[ -n ${ANTHROPIC_BASE_URL-} ]]; then
  export CLIPROXY_BASE_URL="$ANTHROPIC_BASE_URL"
fi
if [[ -n ${ANTHROPIC_AUTH_TOKEN-} ]]; then
  export CLIPROXY_AUTH_TOKEN="$ANTHROPIC_AUTH_TOKEN"
fi
unset ANTHROPIC_BASE_URL ANTHROPIC_AUTH_TOKEN \
  ANTHROPIC_DEFAULT_OPUS_MODEL ANTHROPIC_DEFAULT_SONNET_MODEL ANTHROPIC_DEFAULT_HAIKU_MODEL \
  ANTHROPIC_DEFAULT_OPUS_MODEL_SUPPORTED_CAPABILITIES \
  ANTHROPIC_DEFAULT_SONNET_MODEL_SUPPORTED_CAPABILITIES \
  ANTHROPIC_DEFAULT_HAIKU_MODEL_SUPPORTED_CAPABILITIES \
  CLAUDE_CODE_MODEL_CAPABILITIES CLAUDE_CODE_ENABLE_GATEWAY_MODEL_DISCOVERY \
  CLAUDE_CODE_ALWAYS_ENABLE_EFFORT CLAUDE_CODE_MAX_CONTEXT_TOKENS

# Android SDK
export ANDROID_HOME="$HOME/Android/Sdk"

# Universal PATH additions (prepend user tools; go/cli-tools, grok, bun, local, Android, CUDA).
export PATH="$HOME/go/bin:$HOME/.grok/bin:$HOME/.bun/bin:$HOME/.local/bin:$ANDROID_HOME/emulator:$ANDROID_HOME/platform-tools:/opt/cuda/bin:$PATH"

export OPENCODE_ENABLE_EXA=true

# Vite+ bin (https://viteplus.dev)
. "$HOME/.vite-plus/env"
