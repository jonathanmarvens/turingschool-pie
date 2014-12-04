package pieql_test

import (
	"strings"
	"testing"

	"github.com/turingschool-examples/pie/pieql"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok pieql.Token
		lit string
	}{
		// Special Tokens (whitespace)
		{s: `!`, tok: pieql.ILLEGAL, lit: "!"},
		{s: ``, tok: pieql.EOF, lit: ""},
		{s: ` `, tok: pieql.WS, lit: " "},
		{s: ` X`, tok: pieql.WS, lit: " "},
		{s: "\n \t", tok: pieql.WS, lit: "\n \t"},

		// Misc
		{s: `*`, tok: pieql.ASTERISK, lit: `*`},
		{s: `,`, tok: pieql.COMMA, lit: `,`},

		// Identifiers
		{s: `foo`, tok: pieql.IDENT, lit: `foo`},
		{s: `foo_20 `, tok: pieql.IDENT, lit: `foo`},

		// Keywords
		{s: `SELECT`, tok: pieql.SELECT},
		{s: `FROM`, tok: pieql.FROM},
		{s: `from`, tok: pieql.FROM},
	}

	for i, tt := range tests {
		s := pieql.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		}
	}
}