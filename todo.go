package main

import (
	"fmt"
	"github.com/arreyder/compliance-masonry/commands/diff"
	"github.com/codegangsta/cli"
	"github.com/tg/gosortmap"
)

const (
	todoCommandName  = "todo"
	todoCommandUsage = "Compute Gap Analysis for remaining work. (TODOs)"
)

var (
	todoCommandAliases = []string{"t"}
	todoCommandFlags   = []cli.Flag{
		cli.StringFlag{
			Name:        "opencontrols, o",
			Value:       "opencontrols",
			Usage:       "Set opencontrols directory",
			Destination: &opencontrolDir,
		},
	}
	todoCommand = cli.Command{
		Name:    todoCommandName,
		Aliases: todoCommandAliases,
		Usage:   todoCommandUsage,
		Flags:   todoCommandFlags,
		Action:  todoCommandAction,
	}
)

func todoCommandAction(c *cli.Context) error {
	config := diff.Config{
		Certification:  c.Args().First(),
		OpencontrolDir: opencontrolDir,
	}
	inventory, errs := diff.ComputeGapAnalysisTODO(config)
	if errs != nil && len(errs) > 0 {
		return cli.NewExitError(cli.NewMultiError(errs...).Error(), 1)
	}
	fmt.Fprintf(c.App.Writer, "\nNumber of noncomplete but documented controls: %d\n", len(inventory.MissingControlList))
	for _, standardAndControl := range sortmap.ByKey(inventory.MissingControlList) {
		fmt.Fprintf(c.App.Writer, "%s\n", standardAndControl.Key)
	}
	return nil
}
