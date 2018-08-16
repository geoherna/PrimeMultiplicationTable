# Prime Number Multiplication Table

This program can be used to generate a visual multiplication table using `n`(initially set to 10) prime numbers. This program performs a manual check to decide of a number is prime or not. A Custom `prime` sub package includes the code needed for needed prime number operations.

##Usage
The program can be run using `go run main.go`

##Program Structure
The program consists of the point of entry file (`main`) and a sub package called `prime`

The `main` function initializes an instance of a `struct` named `PrimeTable`, which has several methods associated with it:

`func (t *PrimeTable) Print()`: Prints the visual graph to the CLI

`func (t *PrimeTable) GenerateRows(max int)`: Generates the content for table rows

##Testing / Benchmarks

In order to run unit tests to ensure proper prime number checks + benchmarking, you may use: `go test -bench . ./...`



