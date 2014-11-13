package converter

type Records struct {
	Fields []string
}

func NewRecords(fields []string) *Records {
	r := &Records{
		Fields: fields,
	}
	return r
}

func (r *Records) Convert(line []string) map[string]interface{} {
	data := make(map[string]interface{}, len(r.Fields))
	for idx, f := range r.Fields {
		data[f] = line[idx]
	}
	return data
}
