# Modules

## Creating a go module..
```shell
go mod init example.com/username/repo # module init
go test # Run tests in go module
go list -m all # List all modules

```
## Having dependencies
```go
import "rsc.io/quote" # importing a mdule
quoteText := quote.Hello() # Invoking a method of a 

```