#!/usr/bin/env bash
set -euo pipefail

repo_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)"

sudo pacman -S --needed zsh fzf git github-cli

for filename in .zshenv .zshrc .p10k.zsh; do
    source_file="$repo_dir/files/zsh/$filename"
    target_file="$HOME/$filename"
    zsh -n "$source_file"
    if [[ -L "$target_file" && "$(readlink -f -- "$target_file")" == "$source_file" ]]; then
        continue
    fi
    if [[ -e "$target_file" || -L "$target_file" ]]; then
        backup_file="$target_file.bak"
        index=1
        while [[ -e "$backup_file" || -L "$backup_file" ]]; do
            backup_file="$target_file.bak.$index"
            ((index+=1))
        done
        mv -- "$target_file" "$backup_file"
        printf 'Backed up %s to %s\n' "$target_file" "$backup_file"
    fi
    ln -s -- "$source_file" "$target_file"
done

zsh -ic 'print "Zsh configuration loaded successfully."'
if [[ "$(getent passwd "$(id -un)" | cut -d: -f7)" != /usr/bin/zsh ]]; then
    chsh -s /usr/bin/zsh
fi
printf '\nZsh is ready. Run exec zsh now; log out and back in to update your login shell.\n'
