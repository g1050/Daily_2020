#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <cstdio>
#include <unistd.h>
#include <cassert>
using namespace std;

int main(int argc,char *argv[])
{
    if(argc != 2){
        cout << "Usage:" << argv[0] << " hostname" << endl;
        return 1;
    }

    /* 获取目标主机地址信息 */
    struct hostent* hostinfo  = gethostbyname(argv[1]);
    assert(hostinfo);

    /* 获取daytime服务信息 */
    struct servent* servinfo = getservbyname("daytime","tcp");
    assert(servinfo);
    cout << "Daytime port is:" << ntohs(servinfo->s_port) << endl;

    return 0;
}

