package logger

import (
	"encoding/json"
	"fmt"
)

func Pretty(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s \n", p)
}
