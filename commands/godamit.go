package commands

import (
    "github.com/spf13/cobra"
)

var initCommand = &cobra.Command{Use: "godamit"}
var Cfg Config

func AddCommands() {
    initCommand.AddCommand(container_seed)
    initCommand.AddCommand(version)
}

func Execute() {
    AddCommands()
    initCommand.Execute()
}

func InitializeConfig() {
    err := LoadConfig()
    if err != nil {
        panic(err)
    }
}

func init() {
    initCommand.PersistentFlags().StringVar(&CfgFile, "config", "config.gcfg", "config file (default is path/config.gcfg)")
}
