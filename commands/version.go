package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var version = &cobra.Command{
    Use:   "version",
    Short: "Print the version of godamit",
    Long:  "This is the long text!",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Go DAM IT! version 0.1a")
    },
}
