package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/jbdalido/go-marathon"
	"net/http"
	"os"
)

const (
	VERSION = "0.1.0"
	AUTHOR  = "Steven Borrelli"
	EMAIL   = "steve@aster.is"
)

func LsApps(c *cli.Context) {
	m := MarathonClient(c.GlobalString("host"))

	if len(c.Args()) > 0 {
		for _, app := range c.Args() {
			r, err := m.GetApp(app)
			if err != nil {
				log.Error("App not found: ", app)
			} else {
				PrettyJson(r)
			}
		}
	} else {
		r, err := m.ListApps()

		if err != nil {
			log.Fatal("Error listing apps ", err)
		}
		PrettyJson(r)
	}
}

func Ping(c *cli.Context) {
	url := fmt.Sprintf("http://%s%s/ping", c.GlobalString("host"), gomarathon.APIVersion)

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error contacting marathon url: %s", err)
	}

	log.Info("Ping successful: ", url)
}

func run_cli() {
	app := cli.NewApp()
	app.Name = "marathon-cli"
	app.Usage = "Mange marathon apps and groups"
	app.Version = VERSION
	app.Author = AUTHOR
	app.Email = EMAIL

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Value:  "localhost:8080",
			Usage:  "Marathon host (default localhost:8080)",
			EnvVar: "MARATHON_HOST",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "lsapp",
			Usage:  "List all or one running app",
			Action: LsApps,
		},

		{
			Name:   "ping",
			Usage:  "Test Marathon connection",
			Action: Ping,
		},
	}

	app.Run(os.Args)
}

func main() {
	run_cli()
}
