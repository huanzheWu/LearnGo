package main

import "fmt"

type IntA int32
type IntB int32

type XXX interface {
	write()
	read()
}

func main() {
	var vars []interface{} = make([]interface{}, 5)
	vars[0] = "one"
	vars[1] = "two"
	vars[2] = "three"
	vars[3] = 10
	vars[4] = []byte{'a', 'b', 'c'}

	for index, element := range vars {
		switch value := element.(type) {
		case int:
			fmt.Printf("vars[%d] type is int,value is %d \n", index, value)
		case string:
			fmt.Printf("vars[%d] type is string,value is %s \n", index, value)
		case []byte:
			fmt.Printf("vars[%d] type is []byte,value is %s \n", index, value)
		}
	}
}
