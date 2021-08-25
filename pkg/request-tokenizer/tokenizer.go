package request_tokenizer

import "unicode"

func Tokenize(stringQuery string) []string {
	query := []rune(stringQuery)
	var word []rune
	var parsedQuery []string
	for i := 0; i < len(query); i++ {
		if isWhiteSpace(query[i]) {
			if len(word) > 0 {
				parsedQuery = append(parsedQuery, string(word))
				word = []rune{}
			}
		} else {
			word = append(word, unicode.ToLower(query[i]))
		}
	}
	if len(word) > 0 {
		parsedQuery = append(parsedQuery, string(word))
	}

	return parsedQuery
}

func isWhiteSpace(ch rune) bool {
	return unicode.IsSpace(ch) || unicode.IsPunct(ch)
}
