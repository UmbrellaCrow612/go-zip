package runner

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/go-zip/cli/shared"
	"github.com/UmbrellaCrow612/go-zip/cli/utils"
)

func RunUnZipCmd(options *shared.Options) error {
	if options == nil {
		return errors.New("options cannot be nil")
	}

	if options.Path == "" {
		return errors.New("no .zip file provided")
	}

	zipPath := options.Path

	stat, err := os.Stat(zipPath)
	if err != nil {
		return errors.New("zip file does not exist: " + err.Error())
	}

	if stat.IsDir() {
		return errors.New("path must be a .zip file, not a directory")
	}

	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return errors.New("failed to open zip file: " + err.Error())
	}
	defer reader.Close()

	outputDir := options.OutPath
	if outputDir == "" {
		base := filepath.Base(zipPath)
		outputDir = filepath.Join(filepath.Dir(zipPath), base[:len(base)-len(filepath.Ext(base))])
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return errors.New("failed to create output directory: " + err.Error())
		}
	} else {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return errors.New("failed to create output directory: " + err.Error())
		}
	}

	utils.PrintStdout("Extracting to: " + outputDir)

	for _, f := range reader.File {
		err = extractZipEntry(f, outputDir, options)
		if err != nil {
			utils.PrintStderr("error extracting " + f.Name + ": " + err.Error())
		}
	}

	if options.Flatten {
		if err := flattenSingleDir(outputDir); err != nil {
			utils.PrintStderr("error flattening directory: " + err.Error())
		}
	}

	if err := removeEmptyDirs(outputDir); err != nil {
		utils.PrintStderr("error removing empty directories: " + err.Error())
	}

	utils.PrintStdout("Extraction complete!")
	return nil
}

func flattenSingleDir(root string) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	if len(entries) == 1 && entries[0].IsDir() {
		singleDir := filepath.Join(root, entries[0].Name())

		subEntries, err := os.ReadDir(singleDir)
		if err != nil {
			return err
		}

		for _, sub := range subEntries {
			srcPath := filepath.Join(singleDir, sub.Name())
			dstPath := filepath.Join(root, sub.Name())
			if err := os.Rename(srcPath, dstPath); err != nil {
				return err
			}
		}

		return os.Remove(singleDir)
	}

	return nil
}

func removeEmptyDirs(root string) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(root, entry.Name())
			if err := removeEmptyDirs(subDir); err != nil {
				return err
			}
		}
	}

	// After subdirs cleaned, remove current dir if empty
	entries, err = os.ReadDir(root)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		return os.Remove(root)
	}

	return nil
}
func extractZipEntry(f *zip.File, basePath string, options *shared.Options) error {
	fileName := filepath.Base(f.Name)
	dirName := filepath.Dir(f.Name)

	if f.FileInfo().IsDir() && options.IncludeFolders != nil && !options.IncludeFolders.MatchString(dirName) {
		utils.PrintStdout("Skipped folder: " + f.Name)
		return nil
	}

	if !f.FileInfo().IsDir() && options.IncludeFiles != nil && !options.IncludeFiles.MatchString(fileName) {
		utils.PrintStdout("Skipped file: " + f.Name)
		return nil
	}

	targetPath := filepath.Join(basePath, f.Name)

	// Ensure no path traversal
	if !strings.HasPrefix(targetPath, filepath.Clean(basePath)+string(os.PathSeparator)) {
		return errors.New("illegal file path: " + f.Name)
	}

	if f.FileInfo().IsDir() {
		return os.MkdirAll(targetPath, 0755)
	}

	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}

	src, err := f.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
