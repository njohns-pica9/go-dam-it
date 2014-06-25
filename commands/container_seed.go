package commands

import (
    "fmt"
    "github.com/njohns-pica9/godamit/dam"
    "github.com/njohns-pica9/godamit/damlib"
    "github.com/spf13/cobra"
    "time"
)

var sTimes int
var sPrefix string

var container_seed = &cobra.Command{
    Use:   "container_seed",
    Short: "Seed containers",
    Long:  "Seed containers with random data until a time specified in the times flag",
}

func init() {
    container_seed.Flags().IntVarP(&sTimes, "times", "", 1000, "how many containers to create?")
    container_seed.Flags().StringVarP(&sPrefix, "prefix", "", "seeded_", "Prefix containers")
    container_seed.Run = seed
}

func seed(cmd *cobra.Command, args []string) {
    InitializeConfig()
    http_client := dam.NewConnection(Cfg.API.Url, Cfg.API.Key)
    dam_client := damlib.NewDam(http_client)

    now := time.Now().Format(time.RFC3339)
    prefix := fmt.Sprintf("%s%s", sPrefix, now)
    dam_client.SeedContainers(sTimes, prefix)
}
