package converter

import (
	"encoding/csv"
	"io"
)

type Converter struct {
	Input  io.Reader
	reader *csv.Reader
}

func NewConverter(csvInput io.Reader) *Converter {
	converter := &Converter{}
	converter.reader = csv.NewReader(csvInput)
	return converter
}

func (c *Converter) SetDelimiter(d rune) {
	c.reader.Comma = d
}

func (c *Converter) SetComment(d rune) {
	c.reader.Comment = d
}

func (c *Converter) SetFieldsPerRecord(f int) {
	c.reader.FieldsPerRecord = f
}

func (c *Converter) SetLazyQuotes(l bool) {
	c.reader.LazyQuotes = l
}

func (c *Converter) SetTrim(t bool) {
	c.reader.TrimLeadingSpace = t
}

func (c *Converter) Run() {

}
