package parse

type Parser struct {
	fields []string
}

func (p *Parser) SetFields(fields []string) {
	p.fields = fields
}

func (p *Parser) Field(i int) string {
	if i < len(p.fields) {
		return p.fields[i]
	}
	return ""
}
