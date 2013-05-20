package qorss

import (
	"testing"
	"qorss"
	)

// TODO: make them fail if they dont do what they're supposed to do... ;)


func TestStrip(t *testing.T) {
	var txt string
	txt = "    Some text: that has some; __init__ -a -b --help random punctuation like . or , (including non latin characters) in it. Does it get all removed?？¿!！------------- Some text:  We   have  3   apples 45g each. That makes how many gram?    Some text: We have 3 apples 45g each. That makes how many gram?  "
	t.Logf("Testing string: %s", txt)
	stripped := qorss.Strip(txt, qorss.STRIP_ALL)
	t.Logf("Got: %s", stripped)
	t.Logf("Done.\n")
}


func TestStripMultipleWS(t *testing.T) {
	var txt string
	txt = "  Some text:  We   have  3   apples 45g each. That makes how many gram?  "
	t.Logf("Testing string: %s", txt)
	stripped := qorss.StripMultipleWS(txt)
	t.Logf("Got: %s", stripped)
	t.Logf("Done.\n")
}

func TestStripPunctuation(t *testing.T) {
	var txt string
	txt = "  Some text: that has some; __init__ -a -b --help random punctuation like . or , (including non latin characters) in it. Does it get all removed?？¿!！ "
	t.Logf("Testing string: %s", txt)
	stripped := qorss.StripPunctuation(txt)
	t.Logf("Got: %s", stripped)
	t.Logf("Done.\n")
}

func TestStripNumbers(t *testing.T) {
	var txt string
	txt = "  Some text: We have 3 apples 45g each. That makes how many gram?"
	t.Logf("Testing string: %s", txt)
	stripped := qorss.StripNumbers(txt)
	t.Logf("Got: %s", stripped)
	t.Logf("Done.\n")
}
