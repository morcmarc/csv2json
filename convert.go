package main

import (
	. "github.com/morcmarc/csv2json/converter"

	"github.com/codegangsta/cli"
)

var convertCommand = cli.Command{
	Name:   "convert",
	Usage:  "Convert",
	Action: convertAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "file",
			Usage: "path to file where result should be written to",
		},
		cli.StringFlag{
			Name:  "delimiter, d",
			Value: ",",
		},
	},
}

func convertAction(c *cli.Context) {
	converter := NewConverter()
	converter.Run()
}
