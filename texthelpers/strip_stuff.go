package texthelpers

import (
	"strings"
	// "log"
	md "github.com/fvbock/blackfriday"
	re "regexp"
)

const (
	STRIP_ALL         = 0
	TRIM              = 1
	STRIP_MWS         = 2
	STRIP_PUNCTUATION = 4
	STRIP_NUMBERS     = 8
	STRIP_MARKDOWN    = 16
	STRIP_HTML        = 32
)

func Strip(s string, flags int) string {
	if flags == STRIP_ALL || (STRIP_HTML&flags == STRIP_HTML) {
		// log.Print("STRIP_HTML is not yet implemented!")
		// s = StripHtml(s)
	}
	if flags == STRIP_ALL || (STRIP_MARKDOWN&flags == STRIP_MARKDOWN) {
		// log.Print("STRIP_MARKDOWN is not yet implemented!")
		// s = StripMarkdown(s)
	}
	if flags == STRIP_ALL || (STRIP_NUMBERS&flags == STRIP_NUMBERS) {
		s = StripNumbers(s)
	}
	if flags == STRIP_ALL || (STRIP_PUNCTUATION&flags == STRIP_PUNCTUATION) {
		s = StripPunctuation(s)
	}
	if flags == STRIP_ALL || (STRIP_MWS&flags == STRIP_MWS) {
		s = StripMultipleWS(s)
	}
	if flags == STRIP_ALL || (TRIM&flags == TRIM) {
		s = strings.TrimSpace(s)
	}
	return s
}

func StripMultipleWS(body string) string {
	rx_s := re.MustCompile(`\s{2,}`)
	return rx_s.ReplaceAllLiteralString(body, " ")
}

func StripPunctuation(body string) string {
	// TODO: detect/log if Compile broke
	// TODO: use unicode table range ?
	rx_p := re.MustCompile(`[\_\-⁞«»"':：「」.。,;<>、?？¿!！¡/／\[\]\(\)\*#&@$\\=\+\^]+`)
	return rx_p.ReplaceAllLiteralString(body, " ")
}

func StripNumbers(body string) string {
	rx_n := re.MustCompile(`\d+`)
	return rx_n.ReplaceAllLiteralString(body, "")
}

func StripMarkdown(body string) string {
	strip_md := md.StrippedRenderer(0)
	md_extensions := 0
	return string(md.Markdown([]byte(body), strip_md, md_extensions))
}
