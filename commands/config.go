package commands

import (
    "code.google.com/p/gcfg"
)

type Config struct {
    API struct {
        Url string
        Key string
    }
}

var CfgFile string

func LoadConfig() error {

    err := gcfg.ReadFileInto(&Cfg, CfgFile)

    return err
}
