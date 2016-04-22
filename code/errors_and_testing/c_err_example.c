#include <stdio.h>
#include <errno.h>
#include <string.h>

extern int errno;

int main ()
{
    FILE* pf = fopen("unexist.txt", "rb");
    if (pf == NULL) {
        fprintf(stderr, "Value of errno: %d\n", errno);
        perror("Error printed by perror");
        fprintf(stderr, "Error opening file: %s\n", strerror(errno));
    }
    else {
        fclose(pf);
    }
    return 0;
}
