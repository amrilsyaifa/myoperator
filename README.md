# myoperator

Package myoperator provides simple way to check operator by typing the number.

`go get github.com/amrilsyaifa/myoperator`

The traditional error handling idiom in Go is roughly akin to

```go
number := "08227227xxxx"
result, err := myoperator.GetOperator(number)
if err != nil {
        return fmt.Println("error ", err)
}
```

## Contributing

Because of the Go2 errors changes, this package is not accepting proposals for new functionality. With that said, we welcome pull requests, bug fixes and issue reports.

Before sending a PR, please discuss your change by raising an issue.

## License

BSD-2-Clause
