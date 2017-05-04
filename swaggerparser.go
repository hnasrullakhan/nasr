package main
import (
	"fmt"
	"io/ioutil"
	"os"

	"encoding/json"
)


type Swagger struct{
	Swagger string   `json:"swagger"`
	Paths  interface{} `json:"paths"`
	//Paths 	[]*PathWrapper
	Definitions interface{} `json:"definitions"`
	//Definitions interface{} `json:"definitions"`
}
/*
type PathWrapper struct {
	DynamicPaths 	interface{} `json:"paths"`
}


type definitions struct {
	DynamicDefs 	map[string]interface{} `json:"definitions"`
	//DynamicDefs 	interface{} `json:"-"`

}*/
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
	//for k := range swag.Paths {
	fmt.Printf( "The paths %s \n", (swag.Paths));
	fmt.Println("================================")
	fmt.Printf( "The defintion %s \n", (swag.Definitions));
	if err2 != nil {

		//}
	}
	fmt.Print(err2)
}

