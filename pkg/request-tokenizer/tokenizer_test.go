package request_tokenizer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmptyQuery(t *testing.T) {
	r := Tokenize("")
	require.Empty(t, r)
}

func TestOneWorld(t *testing.T) {
	r := Tokenize("abcde")
	require.Equal(t, []string{"abcde"}, r)
}

func TestTwoWorlds(t *testing.T) {
	r := Tokenize("hello   world")
	require.Equal(t, []string{"hello", "world"}, r)
}

func TestTrim(t *testing.T) {
	r := Tokenize("  pasta   pizza cola ")
	require.Equal(t, []string{"pasta", "pizza", "cola"}, r)
}

func TestRusSimple(t *testing.T) {
	r := Tokenize("  Привет всем ")
	require.Equal(t, []string{"привет", "всем"}, r)
}

func TestRusPunk(t *testing.T) {
	r := Tokenize(" Жил \t \n был \tпес, \r\n хороший    \tтакой \n\n\nпес .\" ")
	require.Equal(t, []string{"жил", "был", "пес", "хороший", "такой", "пес"}, r)
}

func TestRusPunkHard(t *testing.T) {
	r := Tokenize(" Съешь| \" еще \n\rэтих .,.326.,.МяГКиХ @?!французских ^^^булок*" +
		", да __ выпей % чаю    .\" ")
	require.Equal(t, []string{"съешь|", "еще", "этих", "326", "мягких", "французских", "^^^булок",
		"да", "выпей", "чаю"}, r)
}
