#ifndef _LIST_TIMER_
#define _LIST_TIMER_
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
#include <ctime>

#define BUFFER_SIZE 64

class util_timer;
struct client_data{
    sockaddr_in address;
    int sockfd;
    char buf[BUFFER_SIZE];
    util_timer *time;
};

class util_timer{
public:
    util_timer():prev(NULL),next(NULL){}
public:
    time_t expire;//任务超时限制，有绝对时间和相对时间之别
    void (*cb_func)(client_data *);//任务回调函数
    client_data* user_data;
    util_timer* prev;
    util_timer* next;
};

/* 定义定时器链表是一个生序、双向、带有头结点。 */
class sort_timer_list{
public:
   sort_timer_list():head(NULL),tail(NULL){}
   ~sort_timer_list(){
        util_timer *tmp = head;
        while(tmp){
            head = tmp->next;
            delete tmp;
            tmp = head;
        }
   }
public:
   void adjust_timer(util_timer *timer){
       if(!timer){
           return ;
       }

       util_timer* tmp = timer->next;
       //不需要调整
       if(!tmp || (timer->expire < tmp->expire)){
           return ;
       }

       if(timer == head ){
           head = head->next;
           head->prev = NULL;
           timer->next = NULL;
           add_timer(timer,head);
       }else{
           timer->prev->next = timer->next;
           timer->next->prev = timer->prev;
           add_timer(timer,timer->next);
       }

   }
    
   /* 删除的时候指定了定时器 */
   /* 将目标定时器从链表中删除 */
   void del_timer(util_timer* timer){
       if(!timer){
           return ;
       }
        /* 下面条件成立表示只有一个定时器，删除目标定时器即可 */
       if(timer == head && timer == tail){
           delete timer;
           head = NULL;
           tail = NULL;
           return ;
       }

       /* 目标定时器是头结点 */
       if(timer == head){
           head = head->next;
           head->prev = NULL;
           delete timer;
           return;
       }

       /* 目标定时器是尾节点 */
       if(timer == tail){
           tail = tail->prev;
           tail->next = NULL;
           delete timer;
           return ;
       }

       /* 目标定时器处于普通位置 */
       timer->prev->next = timer->next;
       timer->next->prev = timer->prev;
       delete timer;
   }

   /* 将目标定时器timer添加到链表中 */  
   void add_timer(util_timer*timer){

       if(!timer){
           return ;
       }

       //链表为空的时候直接添加
       if(!head){
           head = tail = timer;
           return ;
       }

       /* 需要保证链表是升序的 */
       if(timer->expire < head->expire){//当前定时器小于头结点定时器,直接添加
           timer->next = head;
           head->prev = timer;
           head = timer;
           return ;
       }
       /* 不是最小的话调用重载函数 */
       add_timer(timer,head);
   }

   void tick(){
       if(!head){
           return ;
       }
       std::cout << "Timer tick![定时器滴答响]" << std::endl;
       time_t cur = time(NULL);//获得当前系统的时间
       util_timer* tmp = head;

       /* 从头结点开始处理每个定时器，直到遇到一个尚未到期的定时器，核心逻辑就在这里 */
       while(tmp){
           if(cur < tmp->expire){
               break;
           }
            
           /* 调用定时器的回调函数 */
           tmp->cb_func(tmp->user_data);

           /* 删除 */
           head = tmp->next;
           if(head){
               head->prev = NULL;
           }
           delete tmp;
           tmp = head;
       }
   }
private:
   void add_timer(util_timer *timer,util_timer *lst_head){
       util_timer *prev = lst_head;
       util_timer *tmp = prev->next;

       //遍历寻找合适的位置插入到链表之中
       while(tmp){
           if(timer->expire < tmp->expire){
               prev->next = timer;
               timer->next = tmp;
               tmp->prev = timer;
               timer->prev = prev;
               break;
           }
           prev = tmp;
           tmp = tmp->next;
       }

       //插入到最后
       if(!tmp){
           prev->next = timer;
           timer->prev = prev;
           timer->next = NULL;
           tail = timer;
       }
   }
   util_timer *head;
   util_timer *tail;
};
#endif
