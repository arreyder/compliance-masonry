package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/arreyder/compliance-masonry/commands/docs"
	"github.com/arreyder/compliance-masonry/commands/docs/gitbook"
	"github.com/arreyder/compliance-masonry/commands/get"
	"github.com/arreyder/compliance-masonry/tools/constants"
	"github.com/arreyder/compliance-masonry/tools/fs"
	"github.com/codegangsta/cli"
)

var exportPath, markdownPath, opencontrolDir string

// NewCLIApp creates a new instances of the CLI
func NewCLIApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Compliance Masonry"
	app.Usage = "Open Control CLI Tool"
	app.Version = "1.1.2"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Indicates whether to run the command with verbosity.",
		},
	}
	app.Before = func(c *cli.Context) error {
		// Resets the log to output to nothing
		log.SetOutput(ioutil.Discard)
		if c.Bool("verbose") {
			log.SetOutput(os.Stderr)
			log.Println("Running with verbosity")
		}
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Install compliance dependencies",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dest",
					Value: constants.DefaultDestination,
					Usage: "Location to download the repos.",
				},
				cli.StringFlag{
					Name:  "config",
					Value: constants.DefaultConfigYaml,
					Usage: "Location of system yaml",
				},
			},
			Action: func(c *cli.Context) error {
				f := fs.OSUtil{}
				config := c.String("config")
				configBytes, err := f.OpenAndReadFile(config)
				if err != nil {
					fmt.Fprintf(app.Writer, "%v\n", err.Error())
					os.Exit(1)
				}
				wd, err := os.Getwd()
				if err != nil {
					fmt.Fprintf(app.Writer, "%v\n", err.Error())
					os.Exit(1)
				}
				destination := filepath.Join(wd, c.String("dest"))
				err = get.Get(destination, configBytes)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				fmt.Fprintf(app.Writer, "%v\n", "Compliance Dependencies Installed")
				return nil
			},
		},
		{
			Name:    "docs",
			Aliases: []string{"d"},
			Usage:   "Create Documentation",
			Subcommands: []cli.Command{
				{
					Name:    "gitbook",
					Aliases: []string{"g"},
					Usage:   "Create Gitbook Documentation",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "opencontrols, o",
							Value:       "opencontrols",
							Usage:       "Set opencontrols directory",
							Destination: &opencontrolDir,
						},
						cli.StringFlag{
							Name:        "exports, e",
							Value:       "exports",
							Usage:       "Sets the export directory",
							Destination: &exportPath,
						},
						cli.StringFlag{
							Name:        "markdowns, m",
							Value:       "markdowns",
							Usage:       "Sets the markdowns directory",
							Destination: &markdownPath,
						},
					},
					Action: func(c *cli.Context) error {
						config := gitbook.Config{
							Certification:  c.Args().First(),
							OpencontrolDir: opencontrolDir,
							ExportPath:     exportPath,
							MarkdownPath:   markdownPath,
						}
						warning, errMessages := docs.MakeGitbook(config)
						if warning != "" {
							fmt.Fprintf(app.Writer, "%v\n", warning)
						}
						if errMessages != nil && len(errMessages) > 0 {
							err := cli.NewMultiError(errMessages...)
							return cli.NewExitError(err.Error(), 1)
						}
						fmt.Fprintf(app.Writer, "%v\n", "New Gitbook Documentation Created")
						return nil
					},
				},
			},
		},
		diffCommand,
		todoCommand,
	}
	return app
}

func main() {
	app := NewCLIApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
