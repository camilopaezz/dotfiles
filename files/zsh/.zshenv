# ~/.zshenv — loaded by every zsh shell (interactive, non-interactive, scripts).
# Keep this file minimal and fast: PATH + env only. No plugins, prompts, or aliases.

# Load shared environment variables / secrets if present.
[ -f "$HOME/.config/secrets/ai.env" ] && source "$HOME/.config/secrets/ai.env"

# Android SDK
export ANDROID_HOME="$HOME/Android/Sdk"

# Universal PATH additions (prepend user tools; go/cli-tools, grok, bun, local, Android, CUDA).
export PATH="$HOME/go/bin:$HOME/.grok/bin:$HOME/.bun/bin:$HOME/.local/bin:$ANDROID_HOME/emulator:$ANDROID_HOME/platform-tools:/opt/cuda/bin:$PATH"

export OPENCODE_ENABLE_EXA=true

# >>> Claude Code via CLIProxyAPI (xAI Grok + Claude/GPT + custom models) >>>
# ANTHROPIC_BASE_URL + ANTHROPIC_AUTH_TOKEN live in ~/.config/secrets/ai.env
# Bare model IDs — effort is controlled via Claude Code / T3 effort UI, not
# the CLIProxyAPI model(effort) suffix.
export ANTHROPIC_DEFAULT_OPUS_MODEL=grok-4.6
export ANTHROPIC_DEFAULT_SONNET_MODEL=grok-4.5
export ANTHROPIC_DEFAULT_HAIKU_MODEL=grok-composer-2.5-fast
# Tell Claude Code these gateway IDs support effort (otherwise the effort
# control is hidden for unrecognized/non-claude model names).
_CAPS='effort,xhigh_effort,max_effort,thinking,adaptive_thinking,interleaved_thinking'
export ANTHROPIC_DEFAULT_OPUS_MODEL_SUPPORTED_CAPABILITIES="$_CAPS"
export ANTHROPIC_DEFAULT_SONNET_MODEL_SUPPORTED_CAPABILITIES="$_CAPS"
export ANTHROPIC_DEFAULT_HAIKU_MODEL_SUPPORTED_CAPABILITIES="$_CAPS"
# Per-model capability map for new Claude + GPT hub IDs (and Grok/Kiro).
# Format: modelOrPrefix*=cap1,cap2;other=cap1  (leading -cap disables)
export CLAUDE_CODE_MODEL_CAPABILITIES="grok-*=$_CAPS;claude-*=$_CAPS;gpt-*=$_CAPS;kiro-*=$_CAPS;union-alpha=$_CAPS"
# Discover /v1/models from ANTHROPIC_BASE_URL (CLIProxyAPI) so gateway IDs show up.
export CLAUDE_CODE_ENABLE_GATEWAY_MODEL_DISCOVERY=1
# Blunt hammer so every unrecognized gateway ID still exposes effort in the UI.
# Safe here: CLIProxyAPI clamps unsupported levels.
export CLAUDE_CODE_ALWAYS_ENABLE_EFFORT=1
unset _CAPS
# Claude Code hardcodes 200k for non-claude-* model IDs and does not use
# CLIProxyAPI's max_input_tokens. Grok 4.5/4.6 are 500k upstream.
export CLAUDE_CODE_MAX_CONTEXT_TOKENS=500000
# <<< Claude Code via CLIProxyAPI (xAI Grok + Claude/GPT + custom models) <<<

# Vite+ bin (https://viteplus.dev)
. "$HOME/.vite-plus/env"
