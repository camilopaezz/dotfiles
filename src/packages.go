package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// PackageManager handles package installation via pacman and AUR
type PackageManager struct {
	dryRun bool
}

// NewPackageManager creates a new PackageManager instance
func NewPackageManager(dryRun bool) *PackageManager {
	return &PackageManager{dryRun: dryRun}
}

// IsYayInstalled checks if yay is installed on the system
func (pm *PackageManager) IsYayInstalled() bool {
	_, err := exec.LookPath("yay")
	return err == nil
}

// InstallYay installs yay if it's not already installed
func (pm *PackageManager) InstallYay() error {
	if pm.dryRun {
		fmt.Println(Blue + "DRY RUN: Would install yay" + Reset)
		return nil
	}

	if pm.IsYayInstalled() {
		fmt.Println("yay is already installed")
		return nil
	}

	fmt.Println("Installing yay...")

	// Install required dependencies
	if err := pm.runCommand("sudo", "pacman", "-S", "--needed", "git", "base-devel"); err != nil {
		return fmt.Errorf("failed to install dependencies: %w", err)
	}

	// Clone and build yay
	yayDir := "/tmp/yay"
	if err := pm.runCommand("rm", "-rf", yayDir); err != nil {
		return fmt.Errorf("failed to clean yay directory: %w", err)
	}

	if err := pm.runCommand("git", "clone", "https://aur.archlinux.org/yay.git", yayDir); err != nil {
		return fmt.Errorf("failed to clone yay: %w", err)
	}

	if err := pm.runCommand("cd", yayDir, "&&", "makepkg", "-si"); err != nil {
		return fmt.Errorf("failed to build and install yay: %w", err)
	}

	fmt.Println(Green + "yay installed successfully" + Reset)
	return nil
}

// InstallPackages installs packages using pacman and yay
func (pm *PackageManager) InstallPackages(packages []string, officialPackages []string) error {
	if len(packages) == 0 {
		fmt.Println("No packages to install")
		return nil
	}

	// Separate official packages (pacman) from AUR packages (yay)
	var officialPackagesToInstall []string
	var aurPackages []string

	for _, pkg := range packages {
		if pm.isOfficialPackage(pkg, officialPackages) {
			officialPackagesToInstall = append(officialPackagesToInstall, pkg)
		} else {
			aurPackages = append(aurPackages, pkg)
		}
	}

	// Install official packages with pacman
	if len(officialPackagesToInstall) > 0 {
		if err := pm.installWithPacman(officialPackagesToInstall); err != nil {
			return fmt.Errorf("failed to install official packages: %w", err)
		}
	}

	// Install AUR packages with yay
	if len(aurPackages) > 0 {
		if err := pm.installWithYay(aurPackages); err != nil {
			return fmt.Errorf("failed to install AUR packages: %w", err)
		}
	}

	return nil
}

// isOfficialPackage checks if a package is in official repositories
func (pm *PackageManager) isOfficialPackage(pkg string, officialPackages []string) bool {
	// First check if package is in the official list from config
	for _, official := range officialPackages {
		if pkg == official {
			return true
		}
	}

	if pm.dryRun {
		// In dry run mode, we trust the config classification
		// If it's not in official list, assume it's AUR
		return false
	}

	// Check if package exists in official repos
	err := pm.runCommand("pacman", "-Si", pkg)
	return err == nil
}

// installWithPacman installs packages using pacman
func (pm *PackageManager) installWithPacman(packages []string) error {
	if pm.dryRun {
		fmt.Printf(Blue+"DRY RUN: Would install with pacman: %s\n"+Reset, strings.Join(packages, " "))
		return nil
	}

	fmt.Printf("Installing official packages with pacman: %s\n", strings.Join(packages, " "))
	return pm.runCommand("sudo", append([]string{"pacman", "-S", "--needed"}, packages...)...)
}

// installWithYay installs packages using yay
func (pm *PackageManager) installWithYay(packages []string) error {
	if pm.dryRun {
		fmt.Printf(Blue+"DRY RUN: Would install with yay: %s\n"+Reset, strings.Join(packages, " "))
		return nil
	}

	fmt.Printf("Installing AUR packages with yay: %s\n", strings.Join(packages, " "))
	return pm.runCommand("yay", append([]string{"-S", "--needed"}, packages...)...)
}

// runCommand executes a command and returns any error
func (pm *PackageManager) runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command failed: %s %s - %w", name, strings.Join(args, " "), err)
	}

	return nil
}