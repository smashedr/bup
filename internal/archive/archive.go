package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateZipArchive(excludes []string, source, zipFilename string) error {
	// Create the zip file
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer func() { _ = zipFile.Close() }()

	// Create zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer func() { _ = zipWriter.Close() }()

	// Get the base name of the source for proper relative paths
	sourceBase := filepath.Base(source)

	// Walk the source directory
	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		// Handle errors during walk
		if err != nil {
			return err
		}
		//fmt.Printf("%s: %s\n", info.Name(), path)

		//isExcluded := excludeMap[info.Name()]
		isExcluded := false
		for _, pattern := range excludes {
			matched, err := filepath.Match(pattern, info.Name())
			if err != nil {
				continue
			}
			if matched {
				isExcluded = true
				break
			}
		}

		if info.IsDir() {
			if isExcluded {
				fmt.Printf("SkipDir: %s\n", path)
				return filepath.SkipDir
			}
			// Don't add directory entries to zip, they're created automatically
			return nil
		}
		if isExcluded {
			fmt.Printf("Exclude: %s\n", path)
			return nil
		}

		// Get relative path for zip entry
		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Create a proper zip path (use forward slashes on all platforms)
		zipPath := filepath.ToSlash(filepath.Join(sourceBase, relPath))

		// Create header from file info
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return fmt.Errorf("failed to create header for %s: %w", path, err)
		}

		// Set the header name to our zip path
		header.Name = zipPath

		// Set compression method
		header.Method = zip.Deflate

		// Create writer for this file in the zip
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("failed to create zip entry for %s: %w", path, err)
		}

		// Open the file to copy
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", path, err)
		}
		defer func() { _ = file.Close() }()

		// Copy file content to zip
		_, err = io.Copy(writer, file)
		if err != nil {
			return fmt.Errorf("failed to write file %s to zip: %w", path, err)
		}

		return nil
	})

	if err != nil {
		// Clean up the incomplete zip file on error
		_ = zipFile.Close()
		_ = os.Remove(zipFilename)
		return fmt.Errorf("failed to walk directory: %w", err)
	}

	//_ = zipWriter.Close()
	//fileInfo, err := zipFile.Stat()
	//if err != nil {
	//	fmt.Printf("Error getting archive info: %v", err)
	//} else {
	//	fmt.Printf("Archive Size: %s\n", formatBytes(fileInfo.Size()))
	//}
	//fmt.Printf("Archive: %s\n", zipFilename)

	return nil
}

//func validateDirectory(path, name string) error {
//	info, err := os.Stat(path)
//	if err != nil {
//		if os.IsNotExist(err) {
//			return fmt.Errorf("%s directory does not exist: %s", name, path)
//		}
//		return fmt.Errorf("error accessing %s: %w", name, err)
//	}
//	if !info.IsDir() {
//		return fmt.Errorf("%s is not a directory: %s", name, path)
//	}
//	fmt.Printf("%s: %s\n", name, path)
//	return nil
//}
