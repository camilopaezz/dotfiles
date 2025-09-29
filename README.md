# Dotfiles Manager

🎯 **A modern, interactive CLI tool for managing dotfiles and system packages**

Dotfiles Manager is a Go-based tool that provides an intuitive, interactive interface for copying dotfiles to your home directory and installing system packages. It features a clean, inquirer.js-style interface with arrow key navigation and supports both official packages (via pacman) and AUR packages (via yay).

## ✨ Features

- **🎨 Interactive CLI Interface** - Clean, user-friendly interface with arrow key navigation
- **📦 Complete Installation** - Copy dotfiles and install packages in one command
- **🔧 Selective Operations** - Copy only dotfiles or install only packages
- **👁️ Dry Run Mode** - Preview what would be copied/installed without making changes
- **📋 Smart Package Management** - Automatically separates official vs AUR packages
- **⚙️ Configuration-Driven** - JSON-based configuration for dotfiles and packages
- **🛡️ Error Handling** - Robust error handling with clear user feedback

## 🚀 Quick Start

### Prerequisites

- Go 1.25.1 or later
- Arch Linux (or Arch-based distribution)
- Git (for AUR package management)

### Installation

1. **Clone the repository:**

   ```bash
   git clone <your-repo-url>
   cd dotfiles
   ```

2. **Build the application:**

   ```bash
   go build -o dotfiles src/*.go
   ```

3. **Run the interactive interface:**
   ```bash
   ./dotfiles
   ```

## 📖 Usage

### Interactive Mode

Run `./dotfiles` to launch the interactive menu:

```
🎯 Dotfiles Manager
==================
📦 Complete Install (copy dotfiles + install packages)
📁 Copy Dotfiles Only
🔧 Install Packages Only
👁️  Dry Run: Show what would be copied
🔍 Dry Run: Show what packages would be installed
❌ Exit
```

Navigate with ↑↓ arrow keys and select with Enter.

### Configuration

The tool uses a `dotfiles.json` configuration file to specify which dotfiles to copy and which packages to install.

#### Example Configuration

```json
{
  "dotfiles": {
    "alacritty.toml": ".alacritty.toml",
    "gitconfig": ".gitconfig",
    "zshrc": ".zshrc",
    "tmux.conf": ".tmux.conf"
  },
  "packages": {
    "official": ["git", "curl", "wget", "neovim", "tmux", "zsh"],
    "aur": ["visual-studio-code-bin", "google-chrome"]
  }
}
```

#### Configuration Options

- **`dotfiles`** (object): Maps source files to destination paths

  - Key: Source file path (relative to `files/` directory)
  - Value: Destination path (relative to target directory, typically `~`)

- **`packages`** (object):
  - **`official`** (array): Packages available in official repositories (installed via pacman)
  - **`aur`** (array): Packages from AUR (installed via yay)

## 🎮 Menu Options

### 📦 Complete Install

Copies all configured dotfiles and installs all configured packages. Prompts for:

- Base path for dotfile installation (default: `~`)
- Confirmation before proceeding

### 📁 Copy Dotfiles Only

Copies only the configured dotfiles without installing packages. Prompts for:

- Base path for dotfile installation (default: `~`)
- Confirmation before proceeding

### 🔧 Install Packages Only

Installs only the configured packages without copying dotfiles. Prompts for:

- Confirmation before proceeding

### 👁️ Dry Run: Copy

Shows what dotfiles would be copied without actually copying them. Prompts for:

- Base path to simulate (default: `~`)

### 🔍 Dry Run: Install

Shows what packages would be installed without actually installing them.

### ❌ Exit

Exits the application.

## 🏗️ Project Structure

```
dotfiles/
├── src/
│   ├── main.go          # Main application and CLI interface
│   ├── config.go        # Configuration loading and parsing
│   ├── copier.go        # Dotfile copying functionality
│   ├── packages.go      # Package installation management
│   └── colors.go        # Color constants for output
├── files/               # Directory containing dotfiles to copy
│   ├── alacritty.toml
│   ├── gitconfig
│   ├── zshrc
│   ├── tmux.conf
│   └── hyprconf/        # Subdirectories supported
├── dotfiles.json        # Configuration file
├── go.mod              # Go module definition
└── README.md           # This file
```

## 🔧 Development

### Building

```bash
go build -o dotfiles src/*.go
```

### Running

```bash
./dotfiles
```

### Project Structure

- **`src/main.go`** - Entry point with interactive CLI interface
- **`src/config.go`** - Configuration loading and parsing logic
- **`src/copier.go`** - File copying functionality
- **`src/packages.go`** - Package management (pacman/yay integration)
- **`src/colors.go`** - Color constants for terminal output

### Dependencies

- **`github.com/manifoldco/promptui`** - Interactive CLI prompts
- **`github.com/AlecAivazis/survey/v2`** - (Previously used, now replaced with promptui)

## 🎨 Interface Features

- **Arrow Key Navigation** - Navigate menu options with ↑↓ keys
- **No Filter Text** - Clean interface without typing filter
- **Interactive Prompts** - Input prompts for paths and confirmations
- **Color-Coded Output** - Visual feedback with colors
- **Progress Indicators** - Clear status messages during operations
- **Error Handling** - Graceful error handling with user-friendly messages

## 🛠️ Package Management

The tool intelligently manages both official and AUR packages:

### Official Packages (pacman)

- Installed using `sudo pacman -S --needed`
- Automatically detected based on configuration
- Supports package version locking

### AUR Packages (yay)

- Installed using `yay -S --needed`
- Automatically installs yay if not present
- Handles AUR dependencies automatically

### Smart Classification

The tool respects your configuration's package classification:

- Packages listed in `packages.official` → installed via pacman
- Packages listed in `packages.aur` → installed via yay

## 🔍 Dry Run Mode

Both operations support dry-run mode to preview changes:

- **Copy Dry Run**: Shows which files would be copied and where
- **Install Dry Run**: Shows which packages would be installed via which package manager
- **No Changes Made**: Safe way to preview before actual execution

## 🚨 Error Handling

The application includes comprehensive error handling:

- **Configuration Errors**: Clear messages for malformed JSON or missing files
- **File System Errors**: Handles permission issues and missing directories
- **Package Manager Errors**: Reports issues with pacman/yay operations
- **Network Errors**: Handles AUR cloning and package download issues

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is open source. Please check the LICENSE file for specific terms.

## 🙏 Acknowledgments

- Built with Go and the promptui library for interactive CLI
- Inspired by modern CLI tools and inquirer.js patterns
- Designed for Arch Linux and Arch-based distributions

---

**Happy dotfiling! 🎉**
