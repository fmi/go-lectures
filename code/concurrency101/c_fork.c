#include <stdio.h>

int main()
{
    printf("before\n");
    if (fork())
        printf("parent\n");
    else
        printf("child\n");
    printf("both\n");
}
