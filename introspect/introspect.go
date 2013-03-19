package introspect

import (
	"fmt"
	"reflect"
)

func TypeDump(i interface{}) {
	t := reflect.TypeOf(i)
	// pointer? follow until we find the actual thing
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fmt.Println("TYPE:", t)
	fmt.Printf("      %-15.15s | %-20.20s | %-7.7s | %-10.10s\n", "Name", "Type", "Index", "Offset")
	for y := 0; y < t.NumField(); y++ {
		p := t.Field(y)
		if !p.Anonymous {
			fmt.Printf("  ==> %-15.15s | %-20.20s | %-1.5v | %-10.10v\n", p.Name, p.Type, p.Index, p.Offset)
		}
	}
	for m := 0; m < t.NumMethod(); m++ {
		p := t.Method(m)
		fmt.Printf("  M=> %-15.15s | %-20.20s | %-1.5v\n", p.Name, p.Type, p.Index)
	}

}