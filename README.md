# hcl_parse_sample
Sample scripts to parse and generate HashiCorp Configuration Language (HCL).

```console
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure
$ go run main.go
2017/09/08 10:49:48 --- Parsing HCL ---
2017/09/08 10:49:48 foo: bar
2017/09/08 10:49:48 xyz: abc
2017/09/08 10:49:48 somelist: [one two]
2017/09/08 10:49:48 somemap: map[foo:bar baz:qux]
2017/09/08 10:49:48 --- Generate HCL ---
"foo" = "baz"

"xyz" = "def"

"somelist" = ["one", "two", "three"]

"somemap" = {
  "baz" = "qux"

  "foo" = "bar"

  "quux" = "corge"
}
```
