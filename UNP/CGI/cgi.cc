#include <iostream>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <assert.h>
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

int main()
{
    
    const char *ip = "127.0.0.1";
    const int port = 8080;

    struct sockaddr_in address;
    bzero(&address,sizeof(address));
    address.sin_family = AF_INET;
    inet_pton(AF_INET,ip,&address.sin_addr);
    address.sin_port = htons(port);
    
    int sock = socket(PF_INET,SOCK_STREAM,0);
    int ret = bind(sock,(struct sockaddr*)&address,sizeof(address));

    ret = listen(sock,5);

    struct sockaddr_in client;
    socklen_t length = sizeof(client);
    int connfd = accept(sock,(struct sockaddr*)&client,&length);

    std::cout << "Accept connection!" << std::endl;
    close(STDOUT_FILENO);
    dup(connfd);
    printf("Hello world!\n");

    close(connfd);


    return 0;
}

