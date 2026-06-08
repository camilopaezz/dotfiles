#!/usr/bin/env bash
set -euo pipefail

sudo -v

if command -v paru &>/dev/null; then
    echo "paru already installed"
    exit 0
fi

echo "==> Installing base-devel..."
sudo pacman -S --needed --noconfirm base-devel

echo "==> Installing rustup..."
sudo pacman -S --needed --noconfirm rustup

echo "==> Setting up Rust toolchain..."
rustup default stable

echo "==> Cloning and building paru..."
workdir=$(mktemp -d)
git clone https://aur.archlinux.org/paru.git "$workdir/paru"
cd "$workdir/paru"
makepkg -si --noconfirm
cd /
rm -rf "$workdir"

echo "==> paru installed successfully"
