# hcl_parse_sample
Sample scripts to parse and generate HashiCorp Configuration Language (HCL).

```console
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure
$ go run main.go
2017/09/08 11:11:39 --- Parsing HCL ---
foo = "bar"
xyz = "abc"

somelist = [
  "one",
  "two",
]

somemap = {
  foo = "bar"
  baz = "qux"
}
2017/09/08 11:11:39 --- Operate the elements ---
2017/09/08 11:11:39 foo: bar
2017/09/08 11:11:39 xyz: abc
2017/09/08 11:11:39 somelist: [one two]
2017/09/08 11:11:39 somemap: map[foo:bar baz:qux]
2017/09/08 11:11:39 variables.Foo = "baz"
2017/09/08 11:11:39 variables.Xyz = "def"
2017/09/08 11:11:39 variables.Somelist = append(variables.Somelist, "three")
2017/09/08 11:11:39 variables.Somemap["quux"] = "corge"
2017/09/08 11:11:39 --- Generate HCL ---
"foo" = "baz"

"xyz" = "def"

"somelist" = ["one", "two", "three"]

"somemap" = {
  "baz" = "qux"

  "foo" = "bar"

  "quux" = "corge"
}
```
