package lexer

import (
	"testing"

	"github.com/tkancf/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
	
let add = fn(x, y) {
	x + y;
};
	
let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},       //tests[0]
		{token.IDENT, "five"},    //tests[1]
		{token.ASSIGN, "="},      //tests[2]
		{token.INT, "5"},         //tests[3]
		{token.SEMICOLON, ";"},   //tests[4]
		{token.LET, "let"},       //tests[5]
		{token.IDENT, "ten"},     //tests[6]
		{token.ASSIGN, "="},      //tests[7]
		{token.INT, "10"},        //tests[8]
		{token.SEMICOLON, ";"},   //tests[9]
		{token.LET, "let"},       //tests[10]
		{token.IDENT, "add"},     //tests[11]
		{token.ASSIGN, "="},      //tests[12]
		{token.FUNCTION, "fn"},   //tests[13]
		{token.LPAREN, "("},      //tests[14]
		{token.IDENT, "x"},       //tests[15]
		{token.COMMA, ","},       //tests[16]
		{token.IDENT, "y"},       //tests[17]
		{token.RPAREN, ")"},      //tests[18]
		{token.LBRACE, "{"},      //tests[19]
		{token.IDENT, "x"},       //tests[20]
		{token.PLUS, "+"},        //tests[21]
		{token.IDENT, "y"},       //tests[22]
		{token.SEMICOLON, ";"},   //tests[23]
		{token.RBRACE, "}"},      //tests[24]
		{token.SEMICOLON, ";"},   //tests[25]
		{token.LET, "let"},       //tests[26]
		{token.IDENT, "result"},  //tests[27]
		{token.ASSIGN, "="},      //tests[28]
		{token.IDENT, "add"},     //tests[29]
		{token.LPAREN, "("},      //tests[30]
		{token.IDENT, "five"},    //tests[31]
		{token.COMMA, ","},       //tests[32]
		{token.IDENT, "ten"},     //tests[33]
		{token.RPAREN, ")"},      //tests[34]
		{token.SEMICOLON, ";"},   //tests[35]
		{token.BANG, "!"},        //tests[36]
		{token.MINUS, "-"},       //tests[37]
		{token.SLASH, "/"},       //tests[38]
		{token.ASTERISK, "*"},    //tests[39]
		{token.INT, "5"},         //tests[40]
		{token.SEMICOLON, ";"},   //tests[41]
		{token.INT, "5"},         //tests[42]
		{token.LT, "<"},          //tests[43]
		{token.INT, "10"},        //tests[44]
		{token.GT, ">"},          //tests[45]
		{token.INT, "5"},         //tests[46]
		{token.SEMICOLON, ";"},   //tests[47]
		{token.IF, "if"},         //tests[48]
		{token.LPAREN, "("},      //tests[49]
		{token.INT, "5"},         //tests[50]
		{token.LT, "<"},          //tests[51]
		{token.INT, "10"},        //tests[52]
		{token.RPAREN, ")"},      //tests[53]
		{token.LBRACE, "{"},      //tests[54]
		{token.RETURN, "return"}, //tests[55]
		{token.TRUE, "true"},     //tests[56]
		{token.SEMICOLON, ";"},   //tests[57]
		{token.RBRACE, "}"},      //tests[58]
		{token.ELSE, "else"},     //tests[59]
		{token.LBRACE, "{"},      //tests[60]
		{token.RETURN, "return"}, //tests[61]
		{token.FALSE, "false"},   //tests[62]
		{token.SEMICOLON, ";"},   //tests[63]
		{token.RBRACE, "}"},      //tests[64]
		{token.INT, "10"},        //tests[65]
		{token.EQ, "=="},         //tests[66]
		{token.INT, "10"},        //tests[67]
		{token.SEMICOLON, ";"},   //tests[68]
		{token.INT, "10"},        //tests[69]
		{token.NOT_EQ, "!="},     //tests[70]
		{token.INT, "9"},         //tests[71]
		{token.SEMICOLON, ";"},   //tests[72]
		{token.EOF, ""},          //tests[73]
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
