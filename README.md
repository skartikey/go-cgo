# go-cgo
Exploring CGO

## build and run

    go build -o main
    ./main
This will build the Go program, link it with the C code, and execute the resulting binary.

## Structure

- **main.go**: The main Go program that uses cgo to interact with C code.
- **person.c**: The C file defining a `Person` struct and functions to operate on it.

## Notes

- The `Person` struct is defined in C and used in Go, showcasing the interoperability between Go and C.
- Memory management, including allocation and deallocation, is demonstrated in the code.

Feel free to explore and modify the code to understand more about cgo in Go programs.

---