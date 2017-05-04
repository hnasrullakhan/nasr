package main

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
type model struct{
	Source_dir string
	Objects	[]*shipObject
}

type shipObject struct {
	Name string
}

func main() {
	filePath := "./hx.yaml";
	fmt.Printf( "// reading file %s\n", filePath )
	file, err1 := ioutil.ReadFile( filePath )
	if err1 != nil {
		fmt.Printf( "// error while reading file %s\n", filePath )
		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}

	fmt.Println( "// defining array of struct shipObject" )
	fmt.Print(string(file))
	var mod *model

	err2 := yaml.Unmarshal(file, &mod)
	if err2 != nil {
		fmt.Println("error:", err2)
		os.Exit(1)
	}

	fmt.Println( "// loop over array of structs of shipObject" )
	fmt.Printf( "The model '%s\n", mod.Source_dir  );
	for k := range mod.Objects {
		fmt.Printf( "The name %s \n", mod.Objects[k].Name);
	}
}