package formatter

import (
	"strings"
)

func (f *Formatter) addWhitespace(result *strings.Builder, depth int) {
	indent := strings.Repeat(f.indentChar, f.indentSize*depth)
	result.WriteString(indent)
}
