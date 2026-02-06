package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	cfgFile string
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "answer yes to confirmations")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file [default: ~/.config/bup.yaml]")
	rootCmd.Flags().BoolP("version", "V", false, "version for bup")
}

func initConfig() {
	//fmt.Printf("initConfig: cfgFile: %s\n", cfgFile)
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
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("Unable to read config file: %s\n\n", viper.ConfigFileUsed())
			os.Exit(1)
		}
		fmt.Printf("Config File: %s\n", cfgFile)
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
		fmt.Printf("homeDir: %s\n", homeDir)
		configPath := filepath.Join(homeDir, ".config")
		//fmt.Printf("configPath: %s\n", configPath)
		_ = os.MkdirAll(configPath, 0755)
		configFile := filepath.Join(configPath, "bup.yaml")
		//fmt.Printf("configFile: %s\n", configFile)
		viper.SetConfigFile(configFile)
		_ = viper.SafeWriteConfigAs(configFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config: %s\nUsing Default Config!", configFile)
		}
		fmt.Printf("Config File: %s\n", configFile)
	} else {
		fmt.Printf("Config File: %s\n", viper.ConfigFileUsed())
	}
}
