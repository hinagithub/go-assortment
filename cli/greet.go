package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    "config",
                Aliases: []string{"c"},
                Usage:   "Load configuration from `FILE`",
            },
            &cli.StringFlag{
                Name:    "test",
                Aliases: []string{"t"},
                Usage:   "just test ",
            },
            &cli.StringFlag{
                Name:    "lang",
                Aliases: []string{"l"},
                Value: "english",
                Usage:   "just test ",
            },
        },
        Action :func(cCtx *cli.Context)error {
            name:= "Nefertiti"
            if cCtx.NArg()>0{
                name = cCtx.Args().Get(0)
            }
            if cCtx.String("lang")=="jp"{
                fmt.Println("こんにちは, ", name)
            }else{
                fmt.Println("Hello, ", name)
            }
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}