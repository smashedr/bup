package cmd

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/dustin/go-humanize"
	"github.com/smashedr/bup/internal/archive"
	"github.com/smashedr/bup/internal/styles"
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

	// Parse Source/Destination
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
	if destination == "" {
		destination = promptForDestination()
	}
	log.Debug("Args", "source", source, "destination", destination)

	// Validate Source
	sourceInfo, err := os.Stat(source)
	sourcePath, _ := filepath.Abs(source)
	sourceName := filepath.Base(sourcePath)
	if err != nil || !sourceInfo.IsDir() {
		log.Fatalf("Inalid source: %v", source)
	}
	log.Debug("Source", "sourcePath", sourcePath, "sourceName", sourceName)

	// Validate Destination
	destInfo, err := os.Stat(destination)
	destPath, _ := filepath.Abs(destination)
	if err != nil || !destInfo.IsDir() {
		log.Fatalf("Inalid destination: %v", destination)
	}
	log.Debug("Destination", "destPath", destPath)

	// Ensure Default Destination is Set
	if viper.GetString("destination") == "" {
		viper.Set("destination", destPath)
		err := viper.WriteConfig()
		if err != nil {
			log.Warnf("Error Saving Config: %v", err)
		} else {
			styles.PrintKV("Saved Dest", destPath)
		}
	}

	// Process Excludes
	excludes := viper.GetStringSlice("excludes")
	exclude, _ := cmd.Flags().GetStringSlice("exclude")
	excludes = append(excludes, exclude...)
	log.Infof("Excludes: %v", excludes)

	noConfirm, _ := cmd.Flags().GetBool("yes")
	log.Infof("Skip Confirmation: %v", noConfirm)

	styles.PrintKV("Source", sourcePath)
	styles.PrintKV("Destination", destPath)
	styles.PrintKV("Name", sourceName)

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
			log.Fatalf("prompt error: %v", err)
		}
		if !confirm {
			log.Warnf("Operation Cancelled.")
			return
		}
	}

	fullDestPath := filepath.Join(destPath, sourceName)
	if err := os.MkdirAll(fullDestPath, 0755); err != nil {
		log.Errorf("Error creating directory: %v", err)
		os.Exit(1)
	}
	styles.PrintKV("Directory", fullDestPath)

	// Create timestamp filename
	timestamp := time.Now().Format("06-01-02-15-04-05") // YY-MM-DD-HH-MM-SS
	zipFileName := timestamp + ".zip"
	styles.PrintKV("File Name", zipFileName)
	zipFilePath := filepath.Join(fullDestPath, zipFileName)
	styles.PrintKV("File Path", zipFilePath)

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
		styles.PrintKV("File Size", humanize.Bytes(uint64(fileInfo.Size())))
	}

	fmt.Println(styles.Success.Render("Backup Successful"))
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
