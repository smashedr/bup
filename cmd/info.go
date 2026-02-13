package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func infoCmd(cmd *cobra.Command, args []string) {
	log.Debug("infoCmd:", "args", args)

	fmt.Printf("cfgFile: %s\n", cfgFile)
	fmt.Printf("viper.ConfigFileUsed: %s\n", viper.ConfigFileUsed())
	destination := viper.GetString("destination")
	fmt.Printf("destination: %s\n", destination)
	fmt.Printf("excludes: %s\n", viper.GetStringSlice("excludes"))
}
