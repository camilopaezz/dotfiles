package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FileCopier handles the copying of dotfiles
type FileCopier struct {
	dryRun bool
}

// NewFileCopier creates a new FileCopier instance
func NewFileCopier(dryRun bool) *FileCopier {
	return &FileCopier{dryRun: dryRun}
}

// CopyFile copies a single file from src to dst
func CopyFile(src, dst string) error {
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

// CopyDotfiles copies all dotfiles from the configuration to the base path
func (fc *FileCopier) CopyDotfiles(dotfiles DotfilesConfig, basePath string) error {
	var filesDir string = "files"

	for relSrc, dest := range dotfiles {
		src := filepath.Join(filesDir, relSrc)
		if _, err := os.Stat(src); os.IsNotExist(err) {
			fmt.Printf(Yellow+"Warning: %s not found, skipping\n"+Reset, src)
			continue
		}

		if fc.dryRun {
			fmt.Printf(Blue+"DRY RUN: Would copy %s to %s/%s\n"+Reset, src, basePath, dest)
		} else {
			fmt.Printf("Copying %s to %s\n", src, dest)
		}

		if !fc.dryRun {
			destPath := filepath.Join(basePath, dest)
			if err := CopyFile(src, destPath); err != nil {
				return fmt.Errorf("error copying %s to %s: %w", src, destPath, err)
			}
		}
	}

	return nil
}