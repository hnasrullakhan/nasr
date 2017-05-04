package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"encoding/json"
	"github.com/buger/jsonparser"
	"strings"
)


type Swagger struct{
	Swagger string   `json:"swagger"`
	Paths json.RawMessage
	//Paths  map[string]interface{} `json:"paths"`
	//Definitions  map[string]interface{}  `json:"definitions"`
	Definitions json.RawMessage
}


type definitionsprops struct {
	name string
	Properties map[string]interface{} `json:"properties"`
	indprop []property
	//ppty []*property

}

type property struct {
	Name string
	Type string
	Format string
	Items interface{}
	Enum interface{}
	Default bool

}
func main() {
	filePath := "./swagger_webapp.json";
	fmt.Printf( "// reading file %s\n", filePath )
	file, err1 := ioutil.ReadFile( filePath )

	// parsing Yaml to populate structures
	if err1 != nil {
		fmt.Printf( "// error while reading file %s\n", filePath )
		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}
	var swag *Swagger
	err2 := json.Unmarshal(file,&swag)
	fmt.Printf("this is swag value : %s \n",swag.Swagger)
	//var path *PathWrapper
	//err3 := json.Unmarshal(swag.Paths, &path)
	//fmt.Printf( "The paths %s \n", swag.Paths);
	fmt.Println("================================")

	v,_,_,_ :=jsonparser.Get(swag.Paths,"/"+"summary","get","responses","200","schema","$ref")
	fmt.Printf("%s\n",string(v))
	fmt.Println("================================")
	defintion :=strings.SplitAfter(string(v),"#/definitions/")
	fmt.Println(defintion[1])
	//fmt.Printf( "The defintion %s \n", (swag.Definitions));
	def,_,_,_ := jsonparser.Get(swag.Definitions,defintion[1])
	//fmt.Println(string(def))
	var vardef definitionsprops
	_ = json.Unmarshal(def,&vardef)
	vardef.name = defintion[1]
	fmt.Println(vardef)
	//v,_,_,_ = jsonparser.Get(swag.Definitions,defintion[1],"properties")
	fmt.Println("================================")

	for key,val := range vardef.Properties{
		lname := key
		ltype := val.(map[string]interface{})["type"]
		if ltype == nil{
			ltype =""
		}

		lFormat := val.(map[string]interface{})["format"]
		if lFormat == nil{
			lFormat =""
		}
		lItems := val.(map[string]interface{})["items"]
		if lItems == nil{
			lItems =""
		}
		lEnum := val.(map[string]interface{})["enum"]
		if lEnum == nil{
			lEnum =""
		}
		lDefault := val.(map[string]interface{})["default"]
		if lDefault == nil{
			lDefault = false
		}
		fmt.Println("================================")
		fmt.Println("new property")
		fmt.Println("Name:",lname)
		fmt.Println("Type:",ltype)
		fmt.Println("Format:",lFormat)
		fmt.Println("Items:",lItems)
		fmt.Println("Enum:",lEnum)
		fmt.Println("Default:",lDefault)
		fmt.Println("================================")

		tmpProperty :=property{Name:string(lname),Type:ltype.(string),Format:lFormat.(string),Items:lItems,Enum:lEnum,Default:lDefault.(bool)}

		vardef.indprop = append(vardef.indprop,tmpProperty)

	}

	//var propertymap map[string]json.RawMessage

	fmt.Print(err2)



}

