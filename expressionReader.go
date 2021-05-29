package accessdot

import "strings"

type ExpressionReader struct {
	especialChars     string
	eof               string
	especialCharEvent func(c string, buff string) string
}

func NewExpressionReader(especialChars, eof string) *ExpressionReader {
	return &ExpressionReader{
		especialChars: especialChars,
		eof:           eof,
	}
}

func (reader *ExpressionReader) OnEspecialChar(event func(c string, buff string) string) {
	reader.especialCharEvent = event
}

func (reader *ExpressionReader) Read(exp string) {
	buff := ""

	for _, v := range exp {
		char := string(v)

		if strings.Contains(reader.especialChars, char) {
			buff = reader.especialCharEvent(char, buff)
			continue
		}

		buff += char
	}

	reader.especialCharEvent(reader.eof, buff)
}
