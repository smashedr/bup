package cmd

import (
	"fmt"
	"github.com/smashedr/bup/internal/archive"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.design/x/clipboard"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var backupCmd = &cobra.Command{
	Use:     "backup [source] [destination]",
	Aliases: []string{"b", "bu", "bup"},
	Short:   "Backup source to destination as zip",
	Long:    "Creates a zip archive of the source in the destination with a timestamp filename.",
	Args:    cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("--------------------\n")
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

		// Create timestamp filename
		timestamp := time.Now().Format("06-01-02-15-04-05") // YY-MM-DD-HH-MM-SS
		zipFilename := filepath.Join(fullDestPath, timestamp+".zip")
		fmt.Printf("zipFilename: %s\n", zipFilename)

		if err := archive.CreateZipArchive(excludes, sourcePath, zipFilename); err != nil {
			fmt.Printf("Error creating archive: %v\n", err)
			os.Exit(1)
		}

		copyToClipboard := viper.GetBool("clipboard")
		fmt.Printf("copyToClipboard: %v\n", copyToClipboard)
		if copyToClipboard {
			if err := clipboard.Init(); err != nil {
				fmt.Printf("%v\n", err)
			} else {
				clipboard.Write(clipboard.FmtText, []byte(zipFilename))
			}
		}

		fileInfo, err := os.Stat(zipFilename)
		if err != nil {
			fmt.Printf("Error getting archive info: %v", err)
		} else {
			fmt.Printf("Archive Size: %s\n", formatBytes(fileInfo.Size()))
		}

		fmt.Printf("Success!\n")
	},
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

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.PersistentFlags().BoolP("yes", "y", false, "answer yes to confirmations")
	backupCmd.Flags().StringSliceP("exclude", "e", []string{}, "inline pattern to exclude")
}
