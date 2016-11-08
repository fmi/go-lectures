#include <stdio.h>
#include <unistd.h>

int main()
{
    pid_t pid = fork();
    if (pid == 0) {
        printf("child sleeping...\n");
        execl("/bin/sleep", "/bin/sleep", "2", (char *) 0);
    } else {
        waitpid(pid, NULL, 0);
    }
    printf("done!\n");
    return 0;
}
