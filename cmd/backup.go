package cmd

import (
	"archive/zip"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//var (
//	port int
//)

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

func createZipArchive(excludes []string, source, destination string) error {
	// Create timestamp filename
	timestamp := time.Now().Format("06-01-02-15-04-05") // YY-MM-DD-HH-MM-SS
	zipFilename := filepath.Join(destination, timestamp+".zip")

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

	fmt.Printf("Archive: %s\n", zipFilename)
	return nil
}

var backupCmd = &cobra.Command{
	Use:     "backup [source] [destination]",
	Aliases: []string{"b", "bu", "bup"},
	Short:   "Backup source to destination as zip.",
	Long:    "Creates a zip archive of the source in the destination with a timestamp filename.",
	Args:    cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var source, destination string
		if len(args) == 2 {
			source = args[0]
			destination = args[1]
		} else if len(args) == 1 {
			source = args[0]
			destination = viper.GetString("destination")
		} else {
			source = "."
			destination = viper.GetString("destination")
		}
		//fmt.Printf("source: %s\n", source)
		//fmt.Printf("destination: %s\n", destination)

		excludes := viper.GetStringSlice("excludes")
		exclude, _ := cmd.Flags().GetStringSlice("exclude")
		excludes = append(excludes, exclude...)
		//exclude, _ := cmd.Flags().GetString("exclude")
		//if exclude != "" {
		//	parts := strings.Split(exclude, ",")
		//	for _, part := range parts {
		//		excludes = append(excludes, strings.TrimSpace(part))
		//	}
		//}
		fmt.Printf("Excludes: %s\n", excludes)

		if destination == "" {
			fmt.Print("Enter Destination Path: ")
			var response string
			_, _ = fmt.Scanln(&response)
			//fmt.Printf("response: \"%s\"\n", response)
			responseInfo, err := os.Stat(response)
			if err != nil || !responseInfo.IsDir() {
				fmt.Printf("Error: inalid destination: %s\n", response)
				return
			}
			destination = response
		}

		sourceInfo, err := os.Stat(source)
		if err != nil || !sourceInfo.IsDir() {
			fmt.Printf("Error: inalid source: %s\n", source)
			return
		}
		destInfo, err := os.Stat(destination)
		if err != nil || !destInfo.IsDir() {
			fmt.Printf("Error: inalid destination: %s\n", destination)
			return
		}

		//if err := validateDirectory(source, "Source"); err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//if err := validateDirectory(destination, "Destination"); err != nil {
		//	fmt.Println(err)
		//	return
		//}

		sourcePath, _ := filepath.Abs(source)
		destPath, _ := filepath.Abs(destination)
		sourceName := filepath.Base(sourcePath)

		//viper.SetDefault("destination", destination)
		//viper.WriteConfig()
		if viper.GetString("destination") == "" {
			viper.Set("destination", destPath)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Printf("Error Saving Config: %s\n", err)
			} else {
				fmt.Printf("Saved Default Destination: %s\n", destPath)
			}
		}

		noConfirm, _ := cmd.Flags().GetBool("yes")
		fmt.Printf("Skip Confirmation: %v\n", noConfirm)

		fmt.Printf("Source: %s\n", sourcePath)
		fmt.Printf("Destination: %s\n", destPath)
		fmt.Printf("Name: %s\n", sourceName)

		if !noConfirm {
			fmt.Print("Proceed? (y/N): ")
			var response string
			_, _ = fmt.Scanln(&response)
			//fmt.Printf("response: \"%s\"\n", response)
			if len(response) < 1 || strings.ToUpper(response[:1]) != "Y" {
				fmt.Println("Operation cancelled")
				os.Exit(0)
			}
		}

		fullDestPath := filepath.Join(destPath, sourceName)
		if err := os.MkdirAll(fullDestPath, 0755); err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Directory: %s\n", fullDestPath)

		if err := createZipArchive(excludes, sourcePath, fullDestPath); err != nil {
			fmt.Printf("Error creating archive: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Success!\n")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// https://cobra.dev/docs/tutorials/customizing-cli/
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")
	// backupCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
	backupCmd.PersistentFlags().BoolP("yes", "y", false, "answer yes to confirmations")

	backupCmd.Flags().StringSliceP("exclude", "e", []string{}, "pattern to exclude")
	//backupCmd.Flags().StringP("exclude", "e", "", "comma-separated patterns to exclude")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
