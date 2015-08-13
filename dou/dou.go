package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dou"
	app.Usage = "A Deploy tool for Maodou written in Golang."
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Rebuild & run.",
			Action: func(c *cli.Context) {
				appName := strings.TrimRight(c.Args().Get(len(c.Args())-1), ".go")
				// Build
				cmd := exec.Command("go", "build", c.Args().Get(len(c.Args())-1))
				log.Println(strings.Join(cmd.Args, " "))
				var err_output bytes.Buffer
				cmd.Stdout = os.Stdout
				cmd.Stderr = &err_output

				if err := cmd.Start(); err != nil { //Use start, not run
					log.Println("An error occured: ", err) //replace with logger, or anything you want
				}

				if err := cmd.Wait(); err != nil {
					log.Printf("Build error: %v. %s.\n", err, string(err_output.Bytes()))
					return
				}
				log.Println("Build Success.")

				// Run
				cmd = exec.Command("./" + appName)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				cmd.Run()
			},
		},
		{
			Name:  "pack",
			Usage: "Pack & generate supervisor configuration file",
			Action: func(c *cli.Context) {
				println("pack")
			},
		},
	}

	app.Run(os.Args)
}
