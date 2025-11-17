package runner

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/go-zip/cli/shared"
	"github.com/UmbrellaCrow612/go-zip/cli/utils"
)

func RunZipCmd(options *shared.Options) error {
	if options.Path == "" || options.OutPath == "" {
		utils.PrintStderr("both Path and OutPath must be provided")
		return fmt.Errorf("both Path and OutPath must be provided")
	}

	absInputPath, err := filepath.Abs(options.Path)
	if err != nil {
		utils.PrintStderr(fmt.Sprintf("failed to get absolute input path: %v", err))
		return err
	}
	utils.PrintStdout(fmt.Sprintf("Input path: %s", absInputPath))

	absOutPath, err := filepath.Abs(options.OutPath)
	if err != nil {
		utils.PrintStderr(fmt.Sprintf("failed to get absolute output path: %v", err))
		return err
	}
	utils.PrintStdout(fmt.Sprintf("Output path: %s", absOutPath))

	outFile, err := os.Create(absOutPath)
	if err != nil {
		utils.PrintStderr(fmt.Sprintf("failed to create zip file: %v", err))
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	utils.PrintStdout("Starting to zip files...")

	err = filepath.Walk(absInputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			utils.PrintStderr(fmt.Sprintf("error accessing path %s: %v", path, err))
			return err
		}

		relPath, err := filepath.Rel(absInputPath, path)
		if err != nil {
			return err
		}

		if relPath == "." {
			return nil
		}

		// Folder filter
		if info.IsDir() && options.IncludeFolders != nil && !options.IncludeFolders.MatchString(info.Name()) {
			utils.PrintStdout(fmt.Sprintf("Skipping folder: %s", info.Name()))
			return filepath.SkipDir
		}

		// File filter
		if !info.IsDir() && options.IncludeFiles != nil && !options.IncludeFiles.MatchString(info.Name()) {
			utils.PrintStdout(fmt.Sprintf("Skipping file: %s", info.Name()))
			return nil
		}

		// Flatten option placeholder (can be enhanced)
		if options.Flatten && info.IsDir() {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		utils.PrintStdout(fmt.Sprintf("Adding file: %s", relPath))
		return addFileToZip(zipWriter, path, relPath)
	})

	if err != nil {
		utils.PrintStderr(fmt.Sprintf("error walking folder: %v", err))
		return err
	}

	utils.PrintStdout(fmt.Sprintf("Created zip successfully: %s", absOutPath))
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filePath, zipPath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		utils.PrintStderr(fmt.Sprintf("failed to open file %s: %v", filePath, err))
		return err
	}
	defer file.Close()

	w, err := zipWriter.Create(zipPath)
	if err != nil {
		utils.PrintStderr(fmt.Sprintf("failed to add file to zip %s: %v", zipPath, err))
		return err
	}

	_, err = io.Copy(w, file)
	if err != nil {
		utils.PrintStderr(fmt.Sprintf("failed to write file %s to zip: %v", zipPath, err))
		return err
	}

	return nil
}
