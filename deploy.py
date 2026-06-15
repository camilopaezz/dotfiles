#!/usr/bin/env python3
"""Deploy and manage dotfiles — symlink configs and install packages."""

import argparse
import shutil
import subprocess
import sys
from pathlib import Path

REPO = Path(__file__).resolve().parent
FILES = REPO / "files"
PKG_FILE = REPO / "packages.txt"
BACKUP_SUFFIX = ".bak"

# ── Symlink mappings ────────────────────────────────────────────

FILE_LINKS = {
    ".gitconfig": Path.home() / ".gitconfig",
    "ghostty/config.ghostty": Path.home() / ".config" / "ghostty" / "config",
    "ghostty/themes/noctalia": Path.home() / ".config" / "ghostty" / "themes" / "noctalia",
    "zsh/.zshenv": Path.home() / ".zshenv",
    "zsh/.zshrc": Path.home() / ".zshrc",
    "zsh/.p10k.zsh": Path.home() / ".p10k.zsh",
}

DIR_LINKS = {
    "hypr": Path.home() / ".config" / "hypr",
    "noctalia": Path.home() / ".config" / "noctalia",
}


# ── Helpers ─────────────────────────────────────────────────────

def info(msg):
    print(f"  \033[36m\u2192\033[0m {msg}")


def ok(msg):
    print(f"  \033[32m\u2713\033[0m {msg}")


def warn(msg):
    print(f"  \033[33m~\033[0m {msg}")


def err(msg):
    print(f"  \033[31m\u2717\033[0m {msg}")


def link_file(src, dst, dry_run):
    dst.parent.mkdir(parents=True, exist_ok=True)

    if dst.is_symlink():
        if dst.resolve() == src:
            ok(f"Already linked: {dst.name}")
            return
        if not dry_run:
            dst.unlink()

    if dst.exists() and not dst.is_symlink():
        backup = dst.with_name(dst.name + BACKUP_SUFFIX)
        if dry_run:
            warn(f"Would backup {dst.name} \u2192 {backup.name}")
        else:
            warn(f"Backing up {dst.name} \u2192 {backup.name}")
            shutil.move(str(dst), str(backup))

    if dry_run:
        info(f"Would link: {src} \u2192 {dst}")
    else:
        dst.symlink_to(src)
        ok(f"Linked: {dst.name}")


def link_dir(src_dir, dst_dir, dry_run):
    if not src_dir.is_dir():
        err(f"Source directory not found: {src_dir}")
        return

    for src_path in src_dir.rglob("*"):
        if src_path.is_dir():
            continue
        rel = src_path.relative_to(src_dir)
        dst = dst_dir / rel
        link_file(src_path, dst, dry_run)


# ── Commands ────────────────────────────────────────────────────

def cmd_link(dry_run=False):
    print("\n\033[1mSymlinking dotfiles\u2026\033[0m")

    for rel, dst in FILE_LINKS.items():
        src = FILES / rel
        if not src.exists():
            err(f"Source not found: {src}")
            continue
        link_file(src, dst, dry_run)

    for dir_name, dst in DIR_LINKS.items():
        src = FILES / dir_name
        if not src.is_dir():
            err(f"Source directory not found: {src}")
            continue
        link_dir(src, dst, dry_run)


def cmd_install(aur=False, dry_run=False):
    if not PKG_FILE.exists() or PKG_FILE.stat().st_size == 0:
        warn("packages.txt is empty \u2014 nothing to install")
        return

    pkgs = [
        line.strip()
        for line in PKG_FILE.read_text().splitlines()
        if line.strip() and not line.strip().startswith("#")
    ]
    if not pkgs:
        warn("No packages listed in packages.txt")
        return

    print(f"\n\033[1mInstalling {len(pkgs)} packages\u2026\033[0m")

    if aur:
        for helper in ("yay", "paru"):
            if shutil.which(helper):
                cmd = [helper, "-S", "--noconfirm", *pkgs]
                break
        else:
            err("No AUR helper found (yay/paru)")
            sys.exit(1)
    else:
        cmd = ["sudo", "pacman", "-S", "--noconfirm", *pkgs]

    if dry_run:
        info(f"Would run: {' '.join(cmd)}")
    else:
        subprocess.run(cmd, check=True)
        ok("Packages installed")


# ── CLI ─────────────────────────────────────────────────────────

def main():
    parser = argparse.ArgumentParser(description="Deploy dotfiles")
    parser.add_argument(
        "command",
        choices=["link", "install", "all"],
        help="Action to perform",
    )
    parser.add_argument(
        "--dry-run",
        action="store_true",
        help="Preview changes without applying",
    )
    parser.add_argument(
        "--aur",
        action="store_true",
        help="Install packages using yay/paru (AUR support)",
    )
    args = parser.parse_args()

    if args.command in ("link", "all"):
        cmd_link(dry_run=args.dry_run)
    if args.command in ("install", "all"):
        cmd_install(aur=args.aur, dry_run=args.dry_run)


if __name__ == "__main__":
    main()
