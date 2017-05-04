package main

import (
	"encoding/json"
	"fmt"

)

type Message struct {
	Cmd string `json:"cmd"`
	//Data      json.RawMessage
	Data map[string]interface{} `json:data`
}

type Properties struct {
	Name string
	Letter string
	Number float64

}
/*
type CreateMessage struct {
	Conf map[string]interface{}
	Info map[string]int `json:"info"`
}*/

func main() {
	data := []byte(`{"cmd":"create","data":{"conf":{"type":1},"info":{"type":2}}}`)
	var m Message
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Print(err)
	}
	fmt.Println("this is cmd:",m.Cmd)
	fmt.Println("====================")
	fmt.Println("this is Data:",(m.Data))
	fmt.Println("this is conf:",m.Data["conf"])
	var propArray []Properties
	for k, v := range m.Data {

		fmt.Printf("key[%s] value[%s]\n", k, v)
		fmt.Println(v)
		tmp  := v.(map[string]interface{})
		var tname,tletter string
		var tnumber float64
		tname = k
		for key, val := range tmp{
			tletter = string(key)
			tnumber = val.(float64)

		}
		propArray = append(propArray,Properties{tname,tletter,tnumber})

	}
	for i := range propArray{
		fmt.Println(propArray[i].Name)
		fmt.Println(propArray[i].Letter)

		fmt.Println(propArray[i].Number)
	}


	/*var cm CreateMessage
	if err := json.Unmarshal([]byte(m.Data), &cm); err != nil {
		fmt.Print(err)
	}
	fmt.Println(m.Cmd, cm.Conf, cm.Info)
	switch m.Cmd {
	case "create":
		var cm CreateMessage
		if err := json.Unmarshal([]byte(m.Data), &cm); err != nil {
			fmt.Print(err)
		}
		fmt.Println(m.Cmd, cm.Conf, cm.Info)
	default:
		fmt.Print("bad command")

	}*/
}
