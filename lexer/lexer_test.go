// lexer/lexer.go


package lexer
import ( 
	"testing"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"
	test := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
}