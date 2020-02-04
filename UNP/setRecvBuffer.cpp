#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <cassert>
#include <unistd.h>
#include <cstdlib>
#include <cerrno>
#include <cstring>
using namespace std;

const int BUFFER_SIZE = 1024;


int main(int argc,char *argv[])
{
    if(argc <= 3){//参数个数不足
        cout << "Usage: " << argv[0] << "ip_address port_number recvbufsize" << endl;
        return 1;
    }
    
    //分别取出
    const char *ip = argv[1];
    cout << ip << endl;
    int port = atoi(argv[2]);
    cout << port << endl;
    int recvBuf = atoi(argv[3]);
    cout << recvBuf << endl;
    cout << "==========================\n";

    //开始创建套接字
    struct sockaddr_in address;
    bzero(&address,sizeof(address));
    address.sin_family = AF_INET;
    inet_pton(AF_INET,ip,&address.sin_addr);
    address.sin_port = port;

    int sock = socket(PF_INET,SOCK_STREAM,0);
    assert(sock >= 0);//此处处理其实不好，没有提示错误原因，不方便排查错误
    int len = sizeof(recvBuf);
    
    //先设置TCP缓冲区的大小然后立即读取他,在bind之前设置buffer size
    setsockopt(sock,SOL_SOCKET,SO_RCVBUF,&recvBuf,sizeof(recvBuf));
    getsockopt(sock,SOL_SOCKET,SO_RCVBUF,&recvBuf,(socklen_t*)&len);
    cout << "The TCP socket receive buffer size after setting is " << recvBuf << endl;

    int ret = bind(sock,(struct sockaddr*)&address,sizeof(address));
    assert(ret != 1);

    ret = listen(sock,5);
    assert(ret != -1);

    struct sockaddr_in client;
    socklen_t clientAddLen = sizeof(client);
    int connfd = accept(sock,(struct sockaddr*)&client,&clientAddLen);
    if(connfd < 0){
        /* cout << "error is: " << perror << endl; */
    }else{
        char buffer[BUFFER_SIZE];
        bzero(buffer,BUFFER_SIZE);
        while(recv(connfd,buffer,BUFFER_SIZE-1,0) > 0);
        close(connfd);
    }
    close(sock);
    return 0;
}

