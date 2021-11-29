package parser

import (
	"fmt"
	"strings"
)

var traceLevel = 0

const traceIdentPlaceholder string = "\t..."

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel++ }

func decIdent() { traceLevel-- }

func trace(msg string) string {
	incIdent()
	tracePrint("END " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
