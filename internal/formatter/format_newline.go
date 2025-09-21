package formatter

import (
	"strings"
)

func (f *Formatter) formatNewline(result *strings.Builder) {
	result.WriteString("\n")
}
