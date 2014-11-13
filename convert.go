package main

import (
	"log"
	"os"
	"unicode/utf8"

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
			Usage: "field delimiter, set to ',' by default",
		},
		cli.StringFlag{
			Name:  "comment, c",
			Value: "",
			Usage: "comment character for start of line",
		},
		cli.IntFlag{
			Name:  "fields, f",
			Value: 0,
			Usage: "number of expected fields per record",
		},
		cli.BoolFlag{
			Name:  "lazy, l",
			Usage: "allow lazy quotes",
		},
		cli.BoolFlag{
			Name:  "trim, t",
			Usage: "trim leading space",
		},
	},
}

func convertAction(c *cli.Context) {
	if len(c.Args()) == 0 {
		log.Fatalln("No input given")
	}
	iFile := c.Args()[0]
	csv, err := os.Open(iFile)
	if err != nil {
		log.Fatalf("Could not open file: %s", iFile)
	}

	oFile := c.String("output")
	if oFile == "" {
		log.Fatalln("No output file")
	}

	converter := NewConverter(csv)

	if d := c.String("delimiter"); d != "," {
		db := []byte(d)
		dr, _ := utf8.DecodeRune(db)
		converter.SetDelimiter(dr)
	}

	if c := c.String("comment"); c != "" {
		cb := []byte(c)
		cr, _ := utf8.DecodeRune(cb)
		converter.SetComment(cr)
	}

	if f := c.Int("fields"); f != 0 {
		converter.SetFieldsPerRecord(f)
	}

	if l := c.Bool("lazy"); l {
		converter.SetLazyQuotes(l)
	}

	if t := c.Bool("trim"); t {
		converter.SetTrim(t)
	}

	converter.Run()
}
