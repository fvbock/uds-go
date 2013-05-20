package introspect

import (
	"fmt"
	"reflect"
)

func TypeDump(i interface{}) (typedump string) {
	t := reflect.TypeOf(i)
	// pointer? follow until we find the actual thing
	// fmt.Println(t.Kind())
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	typedump += fmt.Sprintf("TYPE: %s\n", t)
	typedump += fmt.Sprintf("     %-15.15s | %-20.20s | %-7.7s | %-10.10s\n", "Name", "Type", "Index", "Offset")
	typedump += fmt.Sprintf("FIELDS               |                      |         |\n")
	if t.NumField() > 0 {
		for y := 0; y < t.NumField(); y++ {
			p := t.Field(y)
			if !p.Anonymous {
				typedump += fmt.Sprintf(" F=> %-15.15s | %-20.20s | %-1.5v | %-10.10v\n", p.Name, p.Type, p.Index, p.Offset)
			}
		}
	}
	if t.NumMethod() > 0 {
		typedump += fmt.Sprintf("METHODS              |                      |         |\n")
		for m := 0; m < t.NumMethod(); m++ {
			p := t.Method(m)
			typedump += fmt.Sprintf("  M=> %-15.15s | %-20.20s | %-1.5v\n", p.Name, p.Type, p.Index)
		}
	}
	return
}

func PrintTypeDump(i interface{}) {
	fmt.Printf("%s\n", TypeDump(i))
}
