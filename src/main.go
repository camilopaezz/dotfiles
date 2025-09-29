package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func main() {
	for {
		// Show welcome message
		fmt.Println("🎯 Dotfiles Manager")
		fmt.Println("==================")

		// Create the select prompt with promptui (no filter by default)
		prompt := promptui.Select{
			Label: "What would you like to do?",
			Items: []string{
				"📦 Complete Install (copy dotfiles + install packages)",
				"📁 Copy Dotfiles Only",
				"🔧 Install Packages Only",
				"👁️  Dry Run: Show what would be copied",
				"🔍 Dry Run: Show what packages would be installed",
				"❌ Exit",
			},
			Size: 6,
		}

		_, result, err := prompt.Run()
		if err != nil {
			if err == promptui.ErrInterrupt {
				fmt.Println("\n👋 Goodbye!")
				os.Exit(0)
			}
			fmt.Fprintf(os.Stderr, Red+"Error: %v\n"+Reset, err)
			os.Exit(1)
		}

		// Handle the selected action
		switch result {
		case "📦 Complete Install (copy dotfiles + install packages)":
			handleInteractiveCompleteInstall()
		case "📁 Copy Dotfiles Only":
			handleInteractiveCopyDotfiles()
		case "🔧 Install Packages Only":
			handleInteractiveInstallPackages()
		case "👁️  Dry Run: Show what would be copied":
			handleInteractiveDryRunCopy()
		case "🔍 Dry Run: Show what packages would be installed":
			handleInteractiveDryRunInstall()
		case "❌ Exit":
			fmt.Println("👋 Goodbye!")
			os.Exit(0)
		}

		// Ask if user wants to continue
		fmt.Print("\nPress Enter to continue...")
		fmt.Scanln()
	}
}

// handleInteractiveCompleteInstall handles the complete installation with user prompts
func handleInteractiveCompleteInstall() {
	fmt.Println("\n🔧 Complete Installation")
	fmt.Println("======================")

	// Ask for base path
	basePathPrompt := promptui.Prompt{
		Label:   "Enter base path for installation",
		Default: getDefaultHome(),
	}
	basePath, err := basePathPrompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error getting base path: %v\n"+Reset, err)
		return
	}

	// Confirm action
	confirmPrompt := promptui.Prompt{
		Label:     fmt.Sprintf("Install dotfiles and packages to %s? (y/N)", basePath),
		IsConfirm: true,
	}
	_, err = confirmPrompt.Run()
	if err != nil {
		fmt.Println("❌ Installation cancelled")
		return
	}

	// Perform complete installation
	handleCompleteInstall(basePath, false)
}

// handleInteractiveCopyDotfiles handles dotfile copying with user prompts
func handleInteractiveCopyDotfiles() {
	fmt.Println("\n📁 Copy Dotfiles")
	fmt.Println("================")

	// Ask for base path
	basePathPrompt := promptui.Prompt{
		Label:   "Enter base path for dotfiles",
		Default: getDefaultHome(),
	}
	basePath, err := basePathPrompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error getting base path: %v\n"+Reset, err)
		return
	}

	// Confirm action
	confirmPrompt := promptui.Prompt{
		Label:     fmt.Sprintf("Copy dotfiles to %s? (y/N)", basePath),
		IsConfirm: true,
	}
	_, err = confirmPrompt.Run()
	if err != nil {
		fmt.Println("❌ Copy cancelled")
		return
	}

	// Perform dotfile copy
	handleCopyDotfiles(basePath, false)
}

// handleInteractiveInstallPackages handles package installation with user prompts
func handleInteractiveInstallPackages() {
	fmt.Println("\n📦 Install Packages")
	fmt.Println("==================")

	// Confirm action
	confirmPrompt := promptui.Prompt{
		Label:     "Install packages from configuration? (y/N)",
		IsConfirm: true,
	}
	_, err := confirmPrompt.Run()
	if err != nil {
		fmt.Println("❌ Package installation cancelled")
		return
	}

	// Perform package installation
	handleInstallPackages(false)
}

// handleInteractiveDryRunCopy shows what would be copied
func handleInteractiveDryRunCopy() {
	fmt.Println("\n👁️  Dry Run: Copy Dotfiles")
	fmt.Println("=========================")

	// Ask for base path
	basePathPrompt := promptui.Prompt{
		Label:   "Enter base path to show what would be copied",
		Default: getDefaultHome(),
	}
	basePath, err := basePathPrompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error getting base path: %v\n"+Reset, err)
		return
	}

	fmt.Printf("Would copy dotfiles to: %s\n", basePath)
	handleCopyDotfiles(basePath, true)
}

// handleInteractiveDryRunInstall shows what packages would be installed
func handleInteractiveDryRunInstall() {
	fmt.Println("\n👁️  Dry Run: Install Packages")
	fmt.Println("============================")

	fmt.Println("Would install the following packages:")
	handleInstallPackages(true)
}

// printHelp displays detailed help information
func printHelp() {
	fmt.Println("Dotfiles Manager")
	fmt.Println()
	fmt.Println("Usage: dotfiles <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  complete-install    Copy dotfiles and install packages")
	fmt.Println("  copy-dotfiles       Copy dotfiles only")
	fmt.Println("  install-packages    Install packages only")
	fmt.Println("  dry-run-install     Show what packages would be installed (dry run)")
	fmt.Println("  dry-run-copy        Show what dotfiles would be copied (dry run)")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --base-path <path>  Base path for dotfile installation (default: ~)")
	fmt.Println("  --help              Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  dotfiles complete-install")
	fmt.Println("  dotfiles copy-dotfiles --base-path /custom/path")
	fmt.Println("  dotfiles dry-run-install")
}

// printUsage displays brief usage information
func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <command> [--base-path <path>] [--help]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Use --help for more information\n")
}

// getDefaultHome returns the user's home directory
func getDefaultHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error getting home directory: %v\n"+Reset, err)
		os.Exit(1)
	}
	return home
}

// handleCompleteInstall handles the complete installation (copy + packages)
func handleCompleteInstall(basePath string, dryRun bool) {
	fmt.Printf("Complete installation to %s\n", basePath)

	// Ensure base path exists
	if err := os.MkdirAll(basePath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error creating directory %s: %v\n"+Reset, basePath, err)
		os.Exit(1)
	}

	// Copy dotfiles
	if dryRun {
		fmt.Printf(Blue + "DRY RUN: Would copy dotfiles to %s\n" + Reset, basePath)
	} else {
		fmt.Printf("Copying dotfiles to %s...\n", basePath)
	}

	dotfiles, err := LoadConfig("dotfiles.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error loading config: %v\n"+Reset, err)
		os.Exit(1)
	}

	copier := NewFileCopier(dryRun)
	if err := copier.CopyDotfiles(dotfiles, basePath); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error copying dotfiles: %v\n"+Reset, err)
		os.Exit(1)
	}

	if dryRun {
		fmt.Println(Blue + "DRY RUN: Dotfiles would be copied" + Reset)
	} else {
		fmt.Println(Green + "Dotfiles copied successfully" + Reset)
	}

	// Install packages
	if dryRun {
		fmt.Println(Blue + "DRY RUN: Would install packages" + Reset)
	} else {
		fmt.Println("Installing packages...")
	}

	if err := installPackagesFromConfig(dryRun); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error installing packages: %v\n"+Reset, err)
		os.Exit(1)
	}

	if !dryRun {
		fmt.Println(Green + "Packages installed successfully" + Reset)
		fmt.Println(Green + "Complete installation finished successfully!" + Reset)
	}
}

// handleCopyDotfiles handles dotfile copying only
func handleCopyDotfiles(basePath string, dryRun bool) {
	// Ensure base path exists
	if err := os.MkdirAll(basePath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error creating directory %s: %v\n"+Reset, basePath, err)
		os.Exit(1)
	}

	dotfiles, err := LoadConfig("dotfiles.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error loading config: %v\n"+Reset, err)
		os.Exit(1)
	}

	if dryRun {
		fmt.Printf(Blue+"DRY RUN: Would copy dotfiles to %s\n"+Reset, basePath)
	} else {
		fmt.Printf("Copying dotfiles to %s...\n", basePath)
	}

	copier := NewFileCopier(dryRun)
	if err := copier.CopyDotfiles(dotfiles, basePath); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error copying dotfiles: %v\n"+Reset, err)
		os.Exit(1)
	}

	if dryRun {
		fmt.Println(Blue + "DRY RUN: Dotfiles would be copied to " + basePath + Reset)
	} else {
		fmt.Println(Green + "Dotfiles copied successfully to " + basePath + Reset)
	}
}

// handleInstallPackages handles package installation only
func handleInstallPackages(dryRun bool) {
	if dryRun {
		fmt.Println(Blue + "DRY RUN: Would install packages" + Reset)
	} else {
		fmt.Println("Installing packages...")
	}

	if err := installPackagesFromConfig(dryRun); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error installing packages: %v\n"+Reset, err)
		os.Exit(1)
	}

	if dryRun {
		fmt.Println(Blue + "DRY RUN: Packages would be installed" + Reset)
	} else {
		fmt.Println(Green + "Packages installed successfully" + Reset)
	}
}

// installPackagesFromConfig handles the package installation process
func installPackagesFromConfig(dryRun bool) error {
	// Load complete configuration
	config, err := LoadCompleteConfig("dotfiles.json")
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	// Check if there are packages to install
	allPackages := config.Packages.GetAllPackages()
	if len(allPackages) == 0 {
		fmt.Println("No packages specified in configuration")
		return nil
	}

	// Initialize package manager
	pm := NewPackageManager(dryRun)

	// Only install yay if there are AUR packages to install
	aurPackages := config.Packages.AUR
	if len(aurPackages) > 0 {
		if err := pm.InstallYay(); err != nil {
			return fmt.Errorf("error installing yay: %w", err)
		}
	}

	// Install packages
	if err := pm.InstallPackages(allPackages, config.Packages.Official); err != nil {
		return fmt.Errorf("error installing packages: %w", err)
	}

	return nil
}

