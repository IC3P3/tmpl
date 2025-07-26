package data

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyDirectory(srcDir, destDir string) error {
	srcInfo, err := os.Stat(srcDir)
	if err != nil {
		return fmt.Errorf("Could not get template directory info for %s: %w", srcDir, err)
	}

	if !srcInfo.IsDir() {
		return fmt.Errorf("source %s is not a directory", srcDir)
	}

	err = os.MkdirAll(destDir, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("Could not create project directory %s: %w", destDir, err)
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("Could not read project directory %s: %w", srcDir, err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(destDir, entry.Name())

		if entry.IsDir() {
			err = CopyDirectory(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Could not open source file %s: %w.", src, err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("Could not create destination file %s: %w.", dest, err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("Could not copy file content from %s to %s: %w.", src, dest, err)
	}

	srcStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("Could not get source file info %s: %w.", src, err)
	}
	if err = os.Chmod(dest, srcStat.Mode()); err != nil {
		return fmt.Errorf("Could not set permissions for destination file %s: %w", dest, err)
	}

	return nil
}
