package accessdot

type Parser struct {
	reader *ExpressionReader
}

func NewParser(reader *ExpressionReader) *Parser {
	return &Parser{reader}
}

func (parser *Parser) Parse(exp string) []Access {
	list := []Access{}

	buildAccess := func(c string, buff string) string {
		if buff == "" {
			return ""
		}

		list = append(list, NewAccess(c, buff))

		return ""
	}

	parser.reader.OnEspecialChar(buildAccess)
	parser.reader.Read(exp)

	return list
}
