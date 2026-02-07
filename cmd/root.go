package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
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
	Use:   "bup",
	Short: "Easily backup directories to destination with excludes.",
	Long:  "Easily create a timestamped archive of the current directory or [source] to a [destination] or saved destination with excludes.",
	// Consider moving backup to the top level command
	// Run: func(cmd *cobra.Command, args []string) { },
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
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "answer yes to confirmations")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file [default: ~/.config/bup.yaml]")
	rootCmd.PersistentFlags().CountVarP(&verbose, "verbose", "v", "verbose output (-vvv debug)")
	rootCmd.Flags().BoolP("version", "V", false, "version for bup")
}

func onInitialize() {
	initLogger(verbose)
	log.Info("Log Level", "verbose", verbose)

	//viper.SetEnvPrefix("bup")
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

	// Provided Config
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
	viper.SetConfigType("yaml")
	viper.SetConfigName("bup")

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("$HOME/AppData/Local")
	viper.AddConfigPath("$HOME/AppData/Roaming")
	viper.AddConfigPath("$HOME/Library/Application Support")

	if err := viper.ReadInConfig(); err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			homeDir = "."
		}
		log.Debugf("homeDir: %v", homeDir)
		configPath := filepath.Join(homeDir, ".config")
		log.Debugf("configPath: %v", configPath)
		_ = os.MkdirAll(configPath, 0755)
		configFile := filepath.Join(configPath, "bup.yaml")
		log.Debugf("configFile: %v", configFile)
		viper.SetConfigFile(configFile)
		_ = viper.SafeWriteConfigAs(configFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config: %s\nUsing Default Config!", configFile)
		}
		log.Infof("Config File: %v", configFile)
	} else {
		log.Infof("Config File: %v", viper.ConfigFileUsed())
	}
}

func initLogger(verbosity int) {
	log.SetReportCaller(verbosity >= 3)
	log.SetReportTimestamp(verbosity >= 3)
	log.SetTimeFormat("15:04:05")
	//log.SetPrefix("bup")

	switch verbosity {
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
