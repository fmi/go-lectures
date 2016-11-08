#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

void *ticker(void *x_void_ptr) {
    while (42) {
        sleep(1);
        printf("tick\n");
    }

    return NULL;
}

int main() {
    pthread_t ticker_thread;

    if(pthread_create(&ticker_thread, NULL, ticker, NULL)) {
        fprintf(stderr, "Error creating thread\n");
        return 1;
    }

    if(pthread_join(ticker_thread, NULL)) {
        fprintf(stderr, "Error joining thread\n");
        return 2;
    }

    return 0;
}
