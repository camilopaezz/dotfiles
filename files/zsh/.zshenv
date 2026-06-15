# ~/.zshenv — loaded by every zsh shell.
# Keep this file minimal and fast.

# Load shared environment variables / secrets if present.
[ -f "$HOME/.config/secrets/ai.env" ] && source "$HOME/.config/secrets/ai.env"

# Universal PATH additions.
export PATH="$HOME/.bun/bin:$HOME/.local/bin:$PATH"
