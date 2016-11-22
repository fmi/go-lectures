#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

int main(int argc, char *argv[]) {
    char *const cat_args[] = {"cat", argv[1], NULL};
    char *const wc_args[] = {"wc", "-l", NULL};

    if (fork() == 0) {
        int pipefd[2];
        pipe(pipefd);

        if (fork() == 0) {
            // replace standard input with input part of pipe
            dup2(pipefd[0], 0);
            // close unused hald of pipe
            close(pipefd[1]);
            // execute wc
            execvp("wc", wc_args);
        } else {
            // replace standard output with output part of pipe
            dup2(pipefd[1], 1);
            // close unused unput half of pipe
            close(pipefd[0]);
              // execute cat
            execvp("cat", cat_args);
        }
    } else {
        wait(NULL);
    }
}
