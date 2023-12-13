#include <stdio.h>
#include <stdlib.h>
#include "person.h"

void printPerson(Person *p) {
    printf("Name: %s\nAge: %d\n", p->name, p->age);
}

void copyName(char *dest, const char *src) {
    while ((*dest++ = *src++)) {}
}

void freeName(char *name) {
    free(name);
}
