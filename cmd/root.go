package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/confluentinc/go-editor"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	cfgFile string
	verbose int
)

var rootCmd = &cobra.Command{
	Use:   "bup [source] [destination]",
	Short: "Easily backup directories to destination with excludes.",
	Long:  "Easily create a timestamped archive of the current directory or [source] to a [destination] or saved destination with excludes.",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("rootCmd:", "args", args)
		editFlag, _ := cmd.Flags().GetBool("edit")
		infoFlag, _ := cmd.Flags().GetBool("info")
		listFlag, _ := cmd.Flags().GetBool("list")
		log.Info("Flags:", "edit", editFlag, "info", infoFlag, "list", listFlag)

		switch {
		case infoFlag:
			infoCmd(cmd, args)
		case listFlag:
			listCmd(cmd, args)
		case editFlag:
			file := viper.ConfigFileUsed()
			log.Infof("viper.ConfigFileUsed(): %v", file)
			edit := editor.NewEditor()
			if err := edit.Launch(file); err != nil {
				log.Fatal(err)
			}
		default:
			backupCmd(cmd, args)
		}
	},
}

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s from SHA %s)", version, date, commit)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(onInitialize)
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file [default: ~/.config/bup.yaml]")
	rootCmd.Flags().StringSliceP("exclude", "e", []string{}, "inline pattern to exclude")
	rootCmd.Flags().CountVarP(&verbose, "verbose", "v", "verbose output (-vvv debug)")
	rootCmd.Flags().BoolP("version", "V", false, "version for bup")
	rootCmd.Flags().BoolP("yes", "y", false, "answer yes to confirmations")
	// Flag Commands
	rootCmd.Flags().BoolP("edit", "E", false, "edit config in default editor")
	rootCmd.Flags().BoolP("info", "i", false, "information about bup")
	rootCmd.Flags().BoolP("list", "l", false, "list backups")
}

func onInitialize() {
	initLogger()
	log.Info("Log Level", "verbose", verbose)

	// Default Config
	viper.SetDefault("clipboard", true)
	viper.SetDefault("excludes", []string{
		".*cache",
		".venv",
		"build",
		"dist",
		"node_modules",
		"out",
		"venv",
		"*.exe",
	})

	// User Provided Config
	log.Debug("onInitialize", "cfgFile", cfgFile)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Unable to read config file: %v", viper.ConfigFileUsed())
		}
		log.Infof("Config File: %v", cfgFile)
		return
	}

	// Find Config
	configName := "bup"
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("bup")

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("$HOME/AppData/Local")
	viper.AddConfigPath("$HOME/AppData/Roaming")
	viper.AddConfigPath("$HOME/Library/Application Support")

	if err := viper.ReadInConfig(); err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			homeDir = "." // NOTE: improve fallback method
		}
		log.Debugf("homeDir: %v", homeDir)
		configPath := filepath.Join(homeDir, ".config")
		log.Debugf("configPath: %v", configPath)
		if err := os.MkdirAll(configPath, 0755); err != nil {
			log.Fatalf("Creating config directory: %v: %v", configPath, err)
		}
		configFile := filepath.Join(configPath, configName+".yaml")
		log.Infof("Config File: %v", configFile)
		viper.SetConfigFile(configFile)
		if err := viper.SafeWriteConfigAs(configFile); err != nil {
			// NOTE: This will error if the config file exist
			log.Debugf("SafeWriteConfigAs: %v: %v", configFile, err)
		}
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Reading config: %v: %v", configFile, err)
		}
	} else {
		log.Infof("Config File: %v", viper.ConfigFileUsed())
	}
}

func initLogger() {
	log.Debug("initLogger", "verbose", verbose)
	log.SetReportCaller(verbose >= 3)
	log.SetReportTimestamp(verbose >= 3)
	log.SetTimeFormat("15:04:05")
	//log.SetPrefix("bup")

	switch verbose {
	case 0:
		log.SetLevel(log.WarnLevel) // Default
	case 1:
		log.SetLevel(log.InfoLevel) // -v
	case 2:
		log.SetLevel(log.DebugLevel) // -vv
	default:
		log.SetLevel(log.DebugLevel) // -vvv+
	}
}
