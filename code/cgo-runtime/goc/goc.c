#include "goc.h"
#include <stdio.h>

int main() {
  printf("Hi, I am a C program.\n");
  GoString name = {"Doycho", 4};
  GreetFromGo(name);
  return 0;
}
