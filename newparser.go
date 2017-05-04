package main
import (
	"encoding/json"
	"fmt"

)

type Message struct {
	Cmd string `json:"cmd"`
	//Data      json.RawMessage
	Data map[string]json.RawMessage`json:data`
}


type Properties struct {
	Name string
	values json.RawMessage

}
type val struct{
	Type string `json:"type"`
}
func main() {
	data := []byte(`{"cmd":"create","data":{"conf":{"type":"1"},"info":{"type":"2"}}}`)
	var m Message
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Print(err)
	}
	fmt.Println("this is cmd:",m.Cmd)
	fmt.Println("====================")
	fmt.Println("this is Data:",(m.Data))
	var p Properties
	for key,value := range m.Data{
		p.Name = key
		p.values = value
		fmt.Println(string(p.values))
		var q val
		_ = json.Unmarshal(p.values,&q)
		fmt.Println(q)
	}
	fmt.Println(p)



}
