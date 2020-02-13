#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <cassert>
#include <cstdlib>
#include <unistd.h>
#include <cerrno>
#include <cstring>
#include <fcntl.h>
using namespace std;
int timeout_connect(const char *ip,int port,int time)
{
    int ret = 0;
    struct sockaddr_in address;
    bzero( &address,sizeof(address));
    address.sin_family = AF_INET;
    inet_pton(AF_INET,ip,&address.sin_addr);
    address.sin_port = htons(port);

    int sockfd = socket(PF_INET,SOCK_STREAM,0);
    assert(sockfd >= 0);
    struct timeval timeout;
    timeout.tv_usec = 0;
    timeout.tv_sec = time;
    socklen_t len = sizeof(timeout);
    ret = setsockopt(sockfd,SOL_SOCKET,SO_SNDTIMEO,&timeout,len);
    assert(ret != -1);

    ret = connect(sockfd,(struct sockaddr*)&address,sizeof(address));
    if(ret == -1){//返回-1说明connect出错
        if(errno = EINPROGRESS){
            cout << "Time out when connecting" << endl;
        }
        cout << "Sorry,something wrong happen" << endl;
    }   
    else  cout << "Sucessfuly connect!" << endl;
    return sockfd;
}

int main(int argc,char *argv[])
{
    if(argc <= 2){
        cout << "Usage:" << argv[0] << " ip_adress port_number\n";
    }

    const char *ip = argv[1];
    int port = atoi(argv[2]);

    int sockfd = timeout_connect(ip,port,100000000);
    if(sockfd < 0){
        return 1;
    }
    return 0;
}

