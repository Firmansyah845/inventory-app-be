package main

import (
	"fmt"
	"os"

	"awesomeProjectSamb/cmd"
	"awesomeProjectSamb/cmd/app"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var (
	version     = "1.0.0"
	shortCommit = ""
)

func main() {

	appVersion := fmt.Sprintf("%s-%s", version, shortCommit)
	viper.Set("APP_VERSION", appVersion)
	app.Init()
	defer app.Shutdown()

	cliApp := cli.NewApp()
	cliApp.Name = "awesomeProjectSamb"
	cliApp.Version = appVersion

	cliApp.Commands = cli.Commands{
		{
			Name:  "server",
			Usage: "Start server",
			Action: func(c *cli.Context) error {
				cmd.StartServer()
				return nil
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
