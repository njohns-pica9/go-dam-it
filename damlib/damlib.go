package damlib

import (
    "fmt"
    "github.com/njohns-pica9/godamit/dam"
)

type Dam struct {
    Client *dam.DamConnection
}

func NewDam(client *dam.DamConnection) *Dam {
    return &Dam{Client: client}
}

func (d *Dam) SeedContainers(times int, prefix string) {
    fmt.Printf("Seeding %d Containers Prefixed with %s \n", times, prefix)

    for i := 0; i < times; i++ {
        name := fmt.Sprintf("%s_%d", prefix, i)
        fmt.Println("Seeding: ", name)

        resp := d.Client.NewContainer(name)

        if resp {
            fmt.Print(".")
        } else {
            fmt.Print("E")
        }
    }
}
