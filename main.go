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
	d, err := ioutil.ReadFile("./variables.tfvars")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(d))

	variables := Variables{}
	hcl.Unmarshal(d, &variables)

	log.Println("--- Operate the elements ---")

	log.Printf("foo: %s\n", variables.Foo)
	log.Printf("xyz: %s\n", variables.Xyz)
	log.Printf("somelist: %v\n", variables.Somelist)
	log.Printf("somemap: %v\n", variables.Somemap)

	log.Println("variables.Foo = \"baz\"")
	variables.Foo = "baz"

	log.Println("variables.Xyz = \"def\"")
	variables.Xyz = "def"

	log.Println("variables.Somelist = append(variables.Somelist, \"three\")")
	variables.Somelist = append(variables.Somelist, "three")

	log.Println("variables.Somemap[\"quux\"] = \"corge\"")
	variables.Somemap["quux"] = "corge"

	log.Println("--- Generate HCL ---")
	b, err := json.Marshal(variables)
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
