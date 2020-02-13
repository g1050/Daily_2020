#include "list_timer.h"
#include <sys/epoll.h>
#define FD_LIMIT 65535
#define MAX_EVENT_NUMBER 1024
#define TIME_SLOT 5
using namespace std;

static int pipefd[2];
static sort_timer_list timer_list;
static int epolled = 0;

void addsig(int sig)
{
    struct sigaction sa;
    memset(&sa,'\0',sizeof(sa));
    sa.sa_handler = sig_handler;
    sa.sa_flags |= SA_RESTART;
    sigfillset(&sa.sa_mask);
    assert(sigaction(sig,&sa,NULL) != -1);
}
int setnoblocking(int fd)
{
    int old_option = fcntl(fd,F_GETFL);
    int new_option = old_option | O_NONBLOCK;
    fcntl(fd,F_SETFL,new_option);
    return old_option;
}

void addfd(int epollfd,int fd)
{
    epoll_event event;
    event.data.fd = fd;
    event.events = EPOLLIN | EPOLLET;
    epoll_ctl(epollfd,EPOLL_CTL_ADD,fd,&event);
    setnoblocking(fd);

}
int main(int argc,char *argv[])
{
    if(argc <= 2){
        cout << "Usage:" << argv[0] << " ip_address port_number" << endl;
        return 1;
    }

    const char *ip = argv[1];
    const int port = atoi(argv[2]);
    
    struct sockaddr_in address;
    bzero( &address,sizeof(address));
    address.sin_family = AF_INET;
    inet_pton(AF_INET,ip,&address.sin_addr);
    address.sin_port = htons(port);

    int listenfd = socket(PF_INET,SOCK_STREAM,0);
    assert(listenfd >= 0);
    int ret = bind(listenfd,(struct sockaddr*)&address,sizeof(address));
    assert(ret != -1);
    ret = listen(listenfd,5);
    assert(ret != -1);

    epoll_event events[MAX_EVENT_NUMBER];
    int epollfd = epoll_create(5);
    assert(epollfd != -1);
    addfd(epollfd,listenfd);

    ret = socketpair(PF_UNIX,SOCK_STREAM,0,pipefd);
    assert(ret != -1);
    setnoblocking(pipefd[1]);
    addfd(epollfd,pipefd[0]);

    /* 设置信号处理函数 */
    addsig(SIGALRM);
    addsig(SIGTERM);
    return 0;
}

