package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyPrint(v Value) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		fmt.Println()
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		fmt.Println()
		return
	}

	fmt.Println(out.String())

}

