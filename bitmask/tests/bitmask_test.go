package bitmask

import (
	"testing"
	"github.com/fvbock/uds-go/bitmask"
)

func TestFlagsSet(t *testing.T) {
	var flags int
	flags = 1
	t.Logf("Testing flag: %v - %v\n", flags, bitmask.FlagsSet(flags))
	flags = 3
	t.Logf("Testing flag: %v - %v\n", flags, bitmask.FlagsSet(flags))
	flags = 5
	t.Logf("Testing flag: %v - %v\n", flags, bitmask.FlagsSet(flags))
	flags = 6
	t.Logf("Testing flag: %v - %v\n", flags, bitmask.FlagsSet(flags))
	flags = 8
	t.Logf("Testing flag: %v - %v\n", flags, bitmask.FlagsSet(flags))
	flags = 12
	t.Logf("Testing flag: %v - %v\n", flags, bitmask.FlagsSet(flags))
	flags = 15
	t.Logf("Testing flag: %v - %v\n", flags, bitmask.FlagsSet(flags))
}
