package main

import (
	"flag"
	"fmt"
	"os"
)


func main() {
	dryRun := flag.Bool("dry-run", false, "Perform a dry run without making changes")
	installPackages := flag.Bool("install-packages", false, "Install packages from configuration")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, Red+"Error: Please provide a base path\n"+Reset)
		fmt.Fprintf(os.Stderr, "Usage: %s [--dry-run] [--install-packages] <base_path>\n", os.Args[0])
		os.Exit(1)
	}

	basePath := args[0]

	if err := os.MkdirAll(basePath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error creating directory %s: %v\n"+Reset, basePath, err)
		os.Exit(1)
	}

	dotfiles, err := LoadConfig("dotfiles.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error loading config: %v\n"+Reset, err)
		os.Exit(1)
	}

	if *dryRun {
		fmt.Printf(Blue+"DRY RUN: Would copy dotfiles to %s...\n"+Reset, basePath)
	} else {
		fmt.Printf("Copying dotfiles to %s...\n", basePath)
	}

	copier := NewFileCopier(*dryRun)
	if err := copier.CopyDotfiles(dotfiles, basePath); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error copying dotfiles: %v\n"+Reset, err)
		os.Exit(1)
	}

	if *dryRun {
		fmt.Println(Blue+"DRY RUN: Dotfiles would be copied to", basePath+Reset)
	} else {
		fmt.Println(Green+"Dotfiles copied successfully to", basePath+Reset)
	}

	// Install packages if requested
	if *installPackages {
		if err := installPackagesFromConfig(*dryRun); err != nil {
			fmt.Fprintf(os.Stderr, Red+"Error installing packages: %v\n"+Reset, err)
			os.Exit(1)
		}

		if *dryRun {
			fmt.Println(Blue + "DRY RUN: Packages would be installed" + Reset)
		} else {
			fmt.Println(Green + "Packages installed successfully" + Reset)
		}
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

	// Install yay if not present
	if err := pm.InstallYay(); err != nil {
		return fmt.Errorf("error installing yay: %w", err)
	}

	// Install packages
	if err := pm.InstallPackages(allPackages); err != nil {
		return fmt.Errorf("error installing packages: %w", err)
	}

	return nil
}

