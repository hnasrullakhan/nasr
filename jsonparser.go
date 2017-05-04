package main

import (
	"github.com/buger/jsonparser"
	"fmt"
	"io/ioutil"
)
func main() {
	/*
	data := []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)
*/
	filePath := "./swagger_webapp.json";
	fmt.Printf( "// reading file %s\n", filePath )
	file, err1 := ioutil.ReadFile(filePath )
	//fmt.Print(string(file))
	v,_,_,_ :=jsonparser.Get(file, "paths","/about","get","responses","200","schema","$ref")
	fmt.Printf("this is get %s \n",string(v))

	if err1 != nil {
		fmt.Print("this is inside if")
		v,_,_,_ :=jsonparser.Get(file, "basePath")
		fmt.Printf("this is get %s",string(v))
	}
        /*
	// You can specify key path by providing arguments to Get function
	v,_,_,_ :=jsonparser.Get(data, "person", "name", "fullName")
	fmt.Printf("this is value %s\n",string(v))

	// There is `GetInt` and `GetBoolean` helpers if you exactly know key data type
	jsonparser.GetInt(data, "person", "github", "followers")

	// When you try to get object, it will return you []byte slice pointer to data containing it
	// In `company` it will be `{"name": "Acme"}`
	jsonparser.Get(data, "company")

	// If the key doesn't exist it will throw an error
	//var size int64
	//if value, _, err := jsonparser.GetInt(data, "company", "size"); err == nil {
	//	size = value
	//}

	// You can use `ArrayEach` helper to iterate items [item1, item2 .... itemN]
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Println(jsonparser.Get(value, "url"))
	}, "person", "avatars")

	// Or use can access fields by index!
	//jsonparser.GetInt("person", "avatars", "[0]", "url")

	// You can use `ObjectEach` helper to iterate objects { "key1":object1, "key2":object2, .... "keyN":objectN }
	jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		fmt.Printf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
		return nil
	}, "person", "name")

	// The most efficient way to extract multiple keys is `EachKey`

	paths := [][]string{
		[]string{"person", "name", "fullName"},
		[]string{"person", "avatars", "[0]", "url"},
		[]string{"company", "url"},
	}
	jsonparser.EachKey(data, func(idx int, value []byte, vt jsonparser.ValueType, err error){
		switch idx {
		case 0: // []string{"person", "name", "fullName"}
			fmt.Print([]string{"person", "name", "fullName"})
		case 1: // []string{"person", "avatars", "[0]", "url"}
			fmt.Print([]string{"person", "avatars", "[0]", "url"})
		case 2: // []string{"company", "url"},
			fmt.Print([]string{"company", "url"})

		}
	}, paths...)

	// For more information see docs below
	*/
}
