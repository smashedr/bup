package cmd

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/smashedr/bup/internal/archive"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.design/x/clipboard"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func backupCmd(cmd *cobra.Command, args []string) {
	log.Debug("backupCmd:", "args", args)

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
	log.Infof("Excludes: %v", excludes)

	if destination == "" {
		destination = promptForDestination()
	}

	sourceInfo, err := os.Stat(source)
	if err != nil || !sourceInfo.IsDir() {
		log.Errorf("Error: inalid source: %v", source)
		return
	}
	destInfo, err := os.Stat(destination)
	if err != nil || !destInfo.IsDir() {
		log.Errorf("Error: inalid destination: %v", destination)
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
			log.Warnf("Error Saving Config: %v", err)
		} else {
			fmt.Printf("Set Default Destination: %v\n", destPath)
		}
	}

	noConfirm, _ := cmd.Flags().GetBool("yes")
	log.Infof("Skip Confirmation: %v", noConfirm)

	fmt.Printf("Source: %s\n", sourcePath)
	fmt.Printf("Destination: %s\n", destPath)
	fmt.Printf("Name: %s\n", sourceName)

	if !noConfirm {
		var confirm = true
		form := huh.NewConfirm().
			Title("Proceed?").
			Affirmative("Yes.").
			Negative("No!").
			Value(&confirm).
			WithTheme(huh.ThemeDracula())
		err := form.Run()
		if err != nil {
			fmt.Printf("prompt error: %v\n", err)
			os.Exit(1)
		}
		if !confirm {
			fmt.Printf("Operation cancelled\n")
			os.Exit(0)
		}
	}

	fullDestPath := filepath.Join(destPath, sourceName)
	if err := os.MkdirAll(fullDestPath, 0755); err != nil {
		log.Errorf("Error creating directory: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Directory: %s\n", fullDestPath)

	// Create timestamp filename
	timestamp := time.Now().Format("06-01-02-15-04-05") // YY-MM-DD-HH-MM-SS
	zipFileName := timestamp + ".zip"
	fmt.Printf("File Name: %s\n", zipFileName)
	zipFilePath := filepath.Join(fullDestPath, zipFileName)
	fmt.Printf("File Path: %s\n", zipFilePath)

	if err := archive.CreateZipArchive(excludes, sourcePath, zipFilePath); err != nil {
		log.Fatalf("Error creating archive: %v", err)
	}

	copyToClipboard := viper.GetBool("clipboard")
	log.Infof("copyToClipboard: %v", copyToClipboard)
	if copyToClipboard {
		if err := clipboard.Init(); err != nil {
			log.Warnf("Clipboard not available.")
		} else {
			clipboard.Write(clipboard.FmtText, []byte(zipFilePath))
		}
	}

	fileInfo, err := os.Stat(zipFilePath)
	if err != nil {
		log.Warnf("Error getting archive info: %v", err)
	} else {
		fmt.Printf("Archive Size: %s\n", formatBytes(fileInfo.Size()))
	}

	fmt.Printf("Archive File: %s\nSuccess!\n", zipFilePath)
}

func promptForDestination() string {
	validate := func(input string) error {
		if strings.HasPrefix(input, "~") {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("cannot determine home directory")
			}
			input = filepath.Join(homeDir, input[1:])
		}
		info, err := os.Stat(input)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("directory does not exist")
			}
			return err
		}
		if !info.IsDir() {
			return fmt.Errorf("path is not a directory")
		}
		return nil
	}

	var result string
	form := huh.NewInput().
		Title("Enter full path to backup directory.").
		Prompt("> ").
		Validate(validate).
		Value(&result).
		WithTheme(huh.ThemeDracula())
	err := form.Run()
	if err != nil {
		log.Errorf("Prompt failed %v", err)
		return ""
	}
	log.Infof("Result: %q", result)
	absPath, err := filepath.Abs(result)
	if err != nil {
		return ""
	}
	log.Infof("absPath %q", absPath)
	return absPath
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
