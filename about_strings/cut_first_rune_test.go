package about_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const longString = "Относительно длинная строка текста, конвертирование коей в руны должно занимать порядочное время"
const testEnglishString = "String"
const maxUtf8CharSeqLen = 5

//go: noinline
func cutFirstRuneSimpleWay(str string) string {
	if len(str) < 1 {
		return str
	}
	return string([]rune(str)[1:])
}

//go: noinline
func cutFirstRuneOverheadWay(str string) string {
	if len(str) < maxUtf8CharSeqLen {
		return cutFirstRuneSimpleWay(str)
	}

	return str[len(string([]rune(str[:maxUtf8CharSeqLen])[0])):]
}

func TestCutFuncEqualResults(t *testing.T) {
	require.Equal(t, cutFirstRuneSimpleWay(longString), cutFirstRuneOverheadWay(longString))
	require.Equal(t, cutFirstRuneSimpleWay(testEnglishString), cutFirstRuneOverheadWay(testEnglishString))
}

var globalStr1 string
var globalStr2 string

func BenchmarkCutFirstRuneSimpleWay(b *testing.B) {
	var str string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str = cutFirstRuneSimpleWay(longString)
	}
	globalStr1 = str
}

func BenchmarkCutFirstRuneOverheadWay(b *testing.B) {
	var str string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str = cutFirstRuneOverheadWay(longString)
	}
	globalStr2 = str
}
