package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
)

type Variables struct {
	Foo      string            `hcl:"foo" json:"foo"`
	Xyz      string            `hcl:"xyz" json:"xyz"`
	Somelist []string          `hcl:"somelist" json:"somelist"`
	Somemap  map[string]string `hcl:"somemap" json:"somemap"`
}

func main() {
	log.Println("--- Parsing HCL ---")
	tfvars, err := ioutil.ReadFile("./variables.tfvars")
	if err != nil {
		log.Fatal(err)
	}
	varables := Variables{}
	hcl.Unmarshal(tfvars, &varables)

	log.Printf("foo: %s\n", varables.Foo)
	log.Printf("xyz: %s\n", varables.Xyz)
	log.Printf("somelist: %v\n", varables.Somelist)
	log.Printf("somemap: %v\n", varables.Somemap)

	log.Println("--- Generate HCL ---")
	varables.Foo = "baz"
	varables.Xyz = "def"
	varables.Somelist = append(varables.Somelist, "three")
	varables.Somemap["quux"] = "corge"

	b, err := json.Marshal(varables)
	if err != nil {
		log.Fatal(err.Error())
	}

	ast, err := hcl.ParseBytes(b)
	if err != nil {
		log.Fatal(err.Error())
	}

	out := new(bytes.Buffer)
	if err := printer.Fprint(out, ast); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(out)
}
