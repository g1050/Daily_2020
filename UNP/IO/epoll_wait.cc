#include <iostream>
#define TIMEOUT 5000

int timeout = TIMEOUT;
time_t start = time(NULL);
time_t end = time(NULL);

while(1){
    std::cout << "the timeout now is " << timeout << std::endl;
    start = time(NULL);
    int number = epoll_wait(epollfd,events,MAX_EVENT_NUMBER,timeout);
    if((number < 0) && (errno != EINTR)){
        std::cout << "epoll failure" std::endl;
        break;
    }

    /* epoll_wait返回0,说明超时时间到，此时便可处理定时任务，并重置定时时间 */
    if(number == 0){
        timeout = TIMEOUT;
        continue;
    }

    end = time(NULL);

}
