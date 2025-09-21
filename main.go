package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

func main() {
	dryRun := flag.Bool("dry-run", false, "Perform a dry run without copying files")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, Red+"Error: Please provide a base path\n"+Reset)
		fmt.Fprintf(os.Stderr, "Usage: %s [--dry-run] <base_path>\n", os.Args[0])
		os.Exit(1)
	}

	basePath := args[0]

	if err := os.MkdirAll(basePath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error creating directory %s: %v\n"+Reset, basePath, err)
		os.Exit(1)
	}

	configData, err := os.ReadFile("dotfiles.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error reading config file: %v\n"+Reset, err)
		os.Exit(1)
	}

	var dotfiles map[string]string
	if err := json.Unmarshal(configData, &dotfiles); err != nil {
		fmt.Fprintf(os.Stderr, Red+"Error parsing config file: %v\n"+Reset, err)
		os.Exit(1)
	}

	if *dryRun {
		fmt.Printf(Blue+"DRY RUN: Would copy dotfiles to %s...\n"+Reset, basePath)
	} else {
		fmt.Printf("Copying dotfiles to %s...\n", basePath)
	}

	for relSrc, dest := range dotfiles {
		src := filepath.Join("files", relSrc)
		if _, err := os.Stat(src); os.IsNotExist(err) {
			fmt.Printf(Yellow+"Warning: %s not found, skipping\n"+Reset, src)
			continue
		}

		if *dryRun {
			fmt.Printf(Blue+"DRY RUN: Would copy %s to %s/%s\n"+Reset, src, basePath, dest)
		} else {
			fmt.Printf("Copying %s to %s\n", src, dest)
		}

		if !*dryRun {
			destPath := filepath.Join(basePath, dest)
			if err := copyFile(src, destPath); err != nil {
				fmt.Fprintf(os.Stderr, Red+"Error copying %s to %s: %v\n"+Reset, src, destPath, err)
				os.Exit(1)
			}
		}
	}

	if *dryRun {
		fmt.Println(Blue+"DRY RUN: Dotfiles would be copied to", basePath+Reset)
	} else {
		fmt.Println(Green+"Dotfiles copied successfully to", basePath+Reset)
	}
}

func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
