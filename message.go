package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Cmd string `json:"-"`
	Data      json.RawMessage
}

type CreateMessage struct {
	Conf map[string]interface{}
	Info map[string]int `json:"info"`
}

func main() {
	data := []byte(`{"cmd":"create","data":{"conf":{"a":1},"info":{"b":2}}}`)
	var m Message
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Print(err)
	}
	var cm CreateMessage
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
	}
}
