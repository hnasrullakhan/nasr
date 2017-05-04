package main
import (
	"fmt"
	"io/ioutil"
	"os"

	"encoding/json"
	"github.com/buger/jsonparser"
)


type Swagger struct{
	Swagger string   `json:"swagger"`
	Paths json.RawMessage
	//Paths  interface{} `json:"paths"`
	//Definitions interface{} `json:"definitions"`
	Definitions json.RawMessage
}

type PathWrapper struct {
	DynamicPaths 	interface{} `json:"paths"`
}


type definitions struct {
	DynamicDefs 	map[string]interface{} `json:"definitions"`

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
	var path *PathWrapper
	err3 := json.Unmarshal(swag.Paths, &path)
	fmt.Printf( "The paths %s \n", swag.Paths);
	fmt.Println("================================")
	var v json.RawMessage
	v,_,_,_ =jsonparser.Get(swag.Paths,"/about","get","responses","200","schema")
	fmt.Printf("%s\n",(v))
	fmt.Println("================================")

	//fmt.Printf( "The defintion %s \n", (swag.Definitions));
	fmt.Print(err2,err3)

}

