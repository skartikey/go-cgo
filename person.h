#ifndef PERSON_H
#define PERSON_H

typedef struct {
    char name[100];
    int age;
} Person;

void printPerson(Person *p);
void copyName(char *dest, const char *src);
void freeName(char *name);

#endif
