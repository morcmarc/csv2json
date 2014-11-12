package main

import (
	"log"

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
			Value: "",
			Usage: "path to file where result should be written to",
		},
		cli.StringFlag{
			Name:  "delimiter, d",
			Value: ",",
		},
		cli.StringFlag{
			Name:  "quote, q",
			Value: "\"",
		},
		cli.StringFlag{
			Name:  "escape, e",
			Value: "\\",
		},
	},
}

func convertAction(c *cli.Context) {
	if len(c.Args()) == 0 {
		log.Fatalln("No input given")
	}

	oFile := c.String("output")
	if oFile == "" {
		log.Fatalln("No output file")
	}

	converter := &Converter{
		Delimiter: c.String("delimiter"),
		Quote:     c.String("quote"),
		Escape:    c.String("escape"),
	}

	converter.Run()
}
