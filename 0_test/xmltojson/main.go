package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	xj "github.com/basgys/goxml2json"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

type xmlDecoder struct {
}

func (i *xmlDecoder) Decode(data []byte, v interface{}) error {
	buf, err := xj.Convert(bytes.NewReader(data))
	if err != nil {
		return err
	}
	return json.Unmarshal(buf.Bytes(), &v)
}
func main() {
	xmlData := `
		<person>
			<name>John Doe</name>
			<name>John Doe</name>
			<age>30</age>
			<city>New York</city>
		</person>
	`
	result := gojsonq.New(gojsonq.SetDecoder(&xmlDecoder{})).FromString(xmlData).Find("person")
	fmt.Println(result)
	//var result map[string]interface{}
	//err := xml.Unmarshal([]byte(xmlData), &result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//jsonData, err := json.Marshal(result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(jsonData))
}
