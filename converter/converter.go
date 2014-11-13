/*
The converter package handles conversion between csv and json
*/
package converter

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
)

type Converter struct {
	reader          *csv.Reader
	writer          *json.Encoder
	input           io.Reader
	output          io.Writer
	useTypeGuessing bool
}

// Returns a new Converter for the given input and output
func NewConverter(csvInput io.Reader, jsonOutput io.Writer, utg bool) *Converter {
	converter := &Converter{
		input:           csvInput,
		output:          jsonOutput,
		reader:          csv.NewReader(csvInput),
		writer:          json.NewEncoder(bufio.NewWriter(jsonOutput)),
		useTypeGuessing: utg,
	}
	return converter
}

// Delimiter symbol
func (c *Converter) SetDelimiter(d rune) {
	c.reader.Comma = d
}

// Comment character
func (c *Converter) SetComment(d rune) {
	c.reader.Comment = d
}

// Expected fields per line
func (c *Converter) SetFieldsPerRecord(f int) {
	c.reader.FieldsPerRecord = f
}

// Turns on lazy quotes
func (c *Converter) SetLazyQuotes(l bool) {
	c.reader.LazyQuotes = l
}

// Sets Trim Leading Space flag
func (c *Converter) SetTrim(t bool) {
	c.reader.TrimLeadingSpace = t
}

// Processes the input and writes converted objects onto the output
func (c *Converter) Run() {
	f, err := c.reader.Read()
	if err != nil {
		log.Fatalf("Could not read input: %s", err)
	}
	r := NewRecords(f)
	c.output.Write([]byte("["))
	for {
		line, err := c.reader.Read()
		if err == io.EOF {
			break
		}
		c.writer.Encode(r.Convert(line))
		c.output.Write([]byte(","))
	}
	c.output.Write([]byte("]"))
}
