#include <sys/socket.h>
#include <stdio.h>
#include <unistd.h>


int main()
{
    int listenfd = socket(AF_INET, SOCK_STREAM, 0);
    int dupfd = 0;
    dup2(listenfd, dupfd);

    puts("Type something to exit");
    int exit = -1;
    scanf("%d", &exit);
    puts("Got exit signal");

    return 0;
}
