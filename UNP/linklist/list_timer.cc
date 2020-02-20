#include "list_timer.h"
#include <sys/epoll.h>
#include <signal.h>
#define FD_LIMIT 65535
#define MAX_EVENT_NUMBER 1024
#define TIME_SLOT 5
using namespace std;

static int pipefd[2];
static sort_timer_list timer_list;
static int epollfd = 0;

void timer_handler()
{
    /* 定时处理任务，实际上就是调用tick函数 */
    timer_list.tick();
    /* 因为一次alarm调用只会引起一次SIGALARM信号，所以我们要重新定义,以不断触发SIGALARM信号 */
    alarm(TIME_SLOT);
    return ;
}
/* 删除非活动连接socket上的注册事件，并且关闭他 */
void cb_func(client_data * user_data)
{
    epoll_ctl(epollfd,EPOLL_CTL_DEL,user_data->sockfd,0);
    assert(user_data);
    close(user_data->sockfd);
    cout << "[长时间未请求]Close fd:" << user_data->sockfd << std::endl;
}

void sig_handler(int sig)
{

    int save_errno = errno;
    int msg = sig;
    cout << "处理信号 " << sig << endl; 
    //将信号类型写入管道
    send(pipefd[1],(char*)&msg,1,0);
    errno = save_errno;
    return ;

}

void addsig(int sig)
{
    struct sigaction sa;
    //设置信号相关联的处理动作
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
    int optval = 1;
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
    cout << "Server socket " << listenfd << " create sucessful" << endl;
    if (setsockopt(listenfd, SOL_SOCKET, SO_REUSEADDR,
		   (const void *)&optval , sizeof(int)) < 0)
	    return -1;
    int ret = bind(listenfd,(struct sockaddr*)&address,sizeof(address));
    assert(ret != -1);
    ret = listen(listenfd,5);
    assert(ret != -1);

    epoll_event events[MAX_EVENT_NUMBER];
    int epollfd = epoll_create(5);
    assert(epollfd != -1);
    addfd(epollfd,listenfd);

    ret = socketpair(PF_UNIX,SOCK_STREAM,0,pipefd);//创建管道
    assert(ret != -1);
    setnoblocking(pipefd[1]);
    addfd(epollfd,pipefd[0]);
    //将读端设为非阻塞,加入epoll感兴趣事件

    /* 设置信号处理函数 */
    addsig(SIGALRM);
    addsig(SIGTERM);

    bool stop_server = false;
    bool timeout = false;
    client_data* users = new client_data[FD_LIMIT];
    //users是一条链表用来存储,用来存储连接的套接字的信息
    alarm(TIME_SLOT);
    //闹钟函数 TIME_SLOT时间到后，发送SIGALARM信号,TIME_SLOT单位是秒

    while(!stop_server){
        int number = epoll_wait(epollfd,events,MAX_EVENT_NUMBER,-1);
        if((number < 0) &&  (errno != EINTR)){
            std::cout << "epoll failure" << std::endl;
            break;
        }
        
        for(int i = 0;i < number; i++){
            int sockfd = events[i].data.fd;

            /* 处理连接请求 */
            if(sockfd == listenfd){
                cout << "接受到连接请求" << endl;
                struct sockaddr_in client_address;
                socklen_t client_addrlength = sizeof(client_address);
                int connfd = accept(listenfd,(struct sockaddr*)&client_address,&client_addrlength);
                cout << "接收到新连接的客户端" << connfd << endl;
                addfd(epollfd,connfd);
                users[connfd].address = client_address;
                users[connfd].sockfd = connfd;

                /* 创建定时器，设置其回调函数和超时时间,然后绑定定时器和用户数据，最后添加到链表中 */
                util_timer* timer = new util_timer;
                timer->user_data = &users[connfd];
                timer->cb_func = cb_func;
                time_t cur = time(NULL);
                timer->expire = cur + 3 * TIME_SLOT;//3倍数时间之后就断开
                users[connfd].time = timer;
                timer_list.add_timer(timer);
            }
            /* 处理信号 */
            else if((sockfd == pipefd[0]) && (events[i].events & EPOLLIN)){
                //0是读端，从读端读数据
                int sig;
                char signals[1024];
                ret = recv(pipefd[0],signals,sizeof(signals),0);
                if(ret ==  -1){
                    /* handle the error */
                    continue;
                }
                else if(ret == 0){
                    continue;
                }else{
                    for(int i = 0;i<ret;i++){
                        switch(signals[i]){
                        case SIGALRM:{
                                         /* 用timeout标记有定时任务需要处理，但不立即处理，这是因为任务优先级不是很高 */
                                         timeout = true;
                                         break;
                                     }
                        case SIGTERM:{
                                         stop_server = true;
                                         break;
                                     }
                        }
                    }
                }
            }
            /* 处理客户连接收到的数据 */
            else if(events[i].events & EPOLLIN){
                memset(users[sockfd].buf,'\0',BUFFER_SIZE);
                ret = recv(sockfd,users[sockfd].buf,BUFFER_SIZE-1,0);
                std::cout << "Get " << ret << " bytes of client data " << users[sockfd].buf << "\nfrom " << sockfd << std::endl;

                util_timer* timer = users[sockfd].time;
                if(ret < 0){
                    /* 如果发生错误，则关闭连接，并移除其对应的定时器 */
                    if(errno != EAGAIN){
                        cb_func(&users[sockfd]);
                        if(timer){
                            timer_list.del_timer(timer);
                        }
                    }
                }else if(ret == 0){
                    /* 如果对方已经关闭连接,则我们也关闭连接,并关闭相应的定时器 */
                    cb_func(&users[sockfd]);
                    if(timer){
                        timer_list.del_timer(timer);
                    }
                }else{
                    /* 如果某个客户连接上有数据可读，则我们要调整该连接对应的定时器，以移除该连接被关闭的时间 */
                    if(timer){
                        time_t cur = time(NULL);
                        timer->expire = cur + 3 * TIME_SLOT;
                        std::cout << "Adjust timer once" << std::endl;
                        timer_list.adjust_timer(timer);
                    }
                }
            }else{
                //other
            }
        }

        /* 最后处理定时事件，因为I/O事件有更高的优先级，当然，这样做可能导致定时任务不能精确地按照预期的时间执行 */
        if(timeout){
            timer_handler();
            timeout = false;
        }
    }
    close(listenfd);
    close(pipefd[0]);
    close(pipefd[1]);
    delete []users;
    return 0;
}

