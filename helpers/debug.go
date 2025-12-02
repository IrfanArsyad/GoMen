package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// DD dumps the values and dies (stops execution)
// Similar to Laravel's dd() function
func DD(values ...interface{}) {
	dump(values...)
	os.Exit(1)
}

// Dump prints the values without stopping execution
// Similar to Laravel's dump() function
func Dump(values ...interface{}) {
	dump(values...)
}

func dump(values ...interface{}) {
	// Get caller info
	_, file, line, _ := runtime.Caller(2)

	// Print header
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("\033[36m%s:%d\033[0m\n", file, line)
	fmt.Println(strings.Repeat("-", 60))

	for i, v := range values {
		if len(values) > 1 {
			fmt.Printf("\033[33m[%d]\033[0m ", i)
		}
		prettyPrint(v, 0)
		fmt.Println()
	}

	fmt.Println(strings.Repeat("=", 60))
}

func prettyPrint(v interface{}, indent int) {
	if v == nil {
		fmt.Print("\033[90mnull\033[0m")
		return
	}

	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	// Handle pointers
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			fmt.Print("\033[90mnull\033[0m")
			return
		}
		fmt.Print("&")
		prettyPrint(val.Elem().Interface(), indent)
		return
	}

	switch val.Kind() {
	case reflect.String:
		fmt.Printf("\033[32m\"%s\"\033[0m", val.String())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("\033[34m%d\033[0m", val.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Printf("\033[34m%d\033[0m", val.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Printf("\033[34m%g\033[0m", val.Float())

	case reflect.Bool:
		fmt.Printf("\033[35m%t\033[0m", val.Bool())

	case reflect.Slice, reflect.Array:
		fmt.Printf("\033[33m%s\033[0m [\n", typ.String())
		for i := 0; i < val.Len(); i++ {
			printIndent(indent + 1)
			fmt.Printf("\033[90m%d:\033[0m ", i)
			prettyPrint(val.Index(i).Interface(), indent+1)
			fmt.Println(",")
		}
		printIndent(indent)
		fmt.Print("]")

	case reflect.Map:
		fmt.Printf("\033[33m%s\033[0m {\n", typ.String())
		for _, key := range val.MapKeys() {
			printIndent(indent + 1)
			prettyPrint(key.Interface(), indent+1)
			fmt.Print(": ")
			prettyPrint(val.MapIndex(key).Interface(), indent+1)
			fmt.Println(",")
		}
		printIndent(indent)
		fmt.Print("}")

	case reflect.Struct:
		fmt.Printf("\033[33m%s\033[0m {\n", typ.String())
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			if field.PkgPath != "" { // Skip unexported fields
				continue
			}
			printIndent(indent + 1)
			fmt.Printf("\033[36m%s\033[0m: ", field.Name)
			prettyPrint(val.Field(i).Interface(), indent+1)
			fmt.Println(",")
		}
		printIndent(indent)
		fmt.Print("}")

	default:
		// Fallback to JSON for complex types
		jsonBytes, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			fmt.Printf("%+v", v)
		} else {
			fmt.Print(string(jsonBytes))
		}
	}
}

func printIndent(level int) {
	fmt.Print(strings.Repeat("  ", level))
}
