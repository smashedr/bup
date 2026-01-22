package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	version = "0.0.1"
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "bup",
	Short:   "Easily backup directories to destination with excludes.",
	Long:    "Easily create a timestamped archive of the current directory to a destination with excludes.",
	Version: version,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file [default: ~/.config/bup.yaml]")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("verbose", "v", false, "Cobra is Retarded")
}

func initConfig() {
	//fmt.Printf("initConfig: cfgFile: %s\n", cfgFile)
	//viper.SetEnvPrefix("bup")

	// Set default excluded directories
	viper.SetDefault("excludes", []string{
		".cache",
		".venv",
		"build",
		"dist",
		"node_modules",
		"venv",
	})

	if cfgFile != "" {
		// Provided Config
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("Unable to read config file: %s\n\n", viper.ConfigFileUsed())
			os.Exit(1)
		}
	} else {
		// Find Config
		viper.SetConfigType("yaml")
		viper.SetConfigName("bup")

		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("$HOME/.config")
		viper.AddConfigPath("$HOME/AppData/Local")
		viper.AddConfigPath("$HOME/AppData/Roaming")
		viper.AddConfigPath("$HOME/Library/Application Support")

		//viper.SafeWriteConfig()
		//viper.ReadInConfig()

		if err := viper.ReadInConfig(); err != nil {
			home, _ := os.UserHomeDir()
			//fmt.Printf("home: %s\n", home)
			configPath := filepath.Join(home, ".config")
			//fmt.Printf("configPath: %s\n", configPath)
			os.MkdirAll(configPath, 0755)
			configFile := filepath.Join(configPath, "bup.yaml")
			//fmt.Printf("configFile: %s\n", configFile)

			// Set a specific file to create
			viper.SetConfigFile(configFile)
			viper.SafeWriteConfigAs(configFile)
			viper.ReadInConfig()
		}
	}
	fmt.Printf("Config File: %s\n", viper.ConfigFileUsed())
}
