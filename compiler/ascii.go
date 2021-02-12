package compiler

const (
	HORIZONTAL_TAB    = 9   // \t
	LINE_FEED         = 10  // \n
	VERTICAL_TAB      = 11  //
	FORM_FEED         = 12  // \f
	CARRIAGE_RETURN   = 13  // \r
	SPACE             = 32  // space
	EXCLAMATION_MARK  = 33  // !
	DOUBLE_QUOTE      = 34  // "
	NUMBER_SIGN       = 35  // #
	DOLLAR_SIGN       = 36  // $
	PERCENT           = 37  // %
	AMPERSAND         = 38  // &
	SINGLE_QUOTE      = 39  // '
	LEFT_PARENTHESIS  = 40  // (
	RIGHT_PARENTHESIS = 41  // )
	ASTERISK          = 42  // *
	PLUS              = 43  // +
	COMMA             = 44  // ,
	MINUS             = 45  // -
	DOT               = 46  // .
	FORWARD_SLASH     = 47  // /
	COLON             = 58  // :
	SEMI_COLON        = 59  // ;
	LESS_THAN         = 60  // <
	EQUAL_SIGN        = 61  // =
	GREATER_THAN      = 62  // >
	QUESTION_MARK     = 63  // ?
	AT_SYMBOL         = 64  // @
	LEFT_BRACKET      = 91  // [
	BACK_SLASH        = 92  // \
	RIGHT_BRACKET     = 93  // ]
	CARET             = 94  // ^
	UNDERSCORE        = 95  // _
	OPEN_APOSTROPHE   = 96  // `
	LEFT_BRACE        = 123 // {
	VERTICAL_BAR      = 124 // |
	RIGHT_BRACE       = 125 // }
	TILDE             = 126 // ~
)

func IsLetter(char rune) bool {
	return IsLowerLetter(char) || IsUpperLetter(char)
}

func IsLowerLetter(char rune) bool {
	return char >= 97 && char <= 122
}

func IsUpperLetter(char rune) bool {
	return char >= 65 && char <= 90
}

func IsNumber(char rune) bool {
	return char >= 48 && char <= 57
}

func IsBlank(char rune) bool {
	return char == SPACE ||
		char == LINE_FEED ||
		char == FORM_FEED ||
		char == CARRIAGE_RETURN ||
		char == HORIZONTAL_TAB
}
