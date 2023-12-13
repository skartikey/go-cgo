package main

/*
#include "person.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	// create a new Person struct
	person := C.Person{}

	// populate the fields of the struct
	name := C.CString("Arya")
	defer C.free(unsafe.Pointer(name))

	// Call C function to copy the name
	C.copyName((*C.char)(unsafe.Pointer(&person.name[0])), name)

	person.age = 3

	// call the C function with a pointer to the struct
	C.printPerson((*C.Person)(unsafe.Pointer(&person)))

	// Output:
	// Name: Alice
	// Age: 30
}
