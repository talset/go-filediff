# go-filediff
take 2 files. remove from file2 the line from file1


## Build

```bash
go build -o gofilediff gofilediff.go
```


## Example

Create a example2.cleaned file container example2 lines without lines existing in example1.

```bash
gofilediff -d example1 example2
```

split a file in X part containing same amount of lines.
```bash
gofilediff -s -num=3 example2
```
