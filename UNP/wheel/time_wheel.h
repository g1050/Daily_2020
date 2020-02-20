#ifndef _TIME_WHEEL_H_
#define _TIME_WHEEL_H_ 
#include <iostream>
#include <time.h>
#include <netinet/in.h>
#include <stdio.h>

#define BUFFER_SIZE 63
class tw_timer;

/* 链表节点中的一个属性，表示用户的数据 */
struct client_data{

    sockaddr_in address;
    int sockfd;
    char buf[BUFFER_SIZE];
    tw_timer *timer;

};

/* 定时器类对应时间轮中链表的一个个节点 */
class tw_timer{
public:
    tw_timer(int rot,int ts):
        next(NULL),
        prev(NULL),
        rotation(rot),
        time_slot(ts){}
public:
    /* 记录定时器在时间轮转多少圈后开始生效 */
    int rotation;
    /* 记录定时器属于时间轮的哪个槽 */
    int time_slot; 
    void (*cb_func)(client_data *);//定时器回调函数
    client_data* user_data;//客户数据
public:
    tw_timer* next;
    tw_timer* prev;
};

class time_wheel{
public:
    time_wheel():cur_slot(0){
        for(int i = 0;i<N;i++){
            slots[i] = NULL;//初始化每个槽的头结点
        }
    }
    ~time_wheel(){
        /* 遍历每一个槽并且销毁其中的定时器 */
        for(int i = 0;i<N;i++){
            tw_timer *tmp = slots[i];
            while(tmp){
                slots[i] = tmp->next;
                delete tmp;
                tmp = slots[i];
            }
        }
    }
public:
    /* 根据定时值timeout创建一个定时器,并且加入时间轮合适的位置中 */
    /* 添加的时候不初始化用户数据，而只是返回一个指针 */
    /* 而用户数据是public的所以，返回的指针可以直接指向用户数据 */
    tw_timer* add_timer(int timeout){
        if(timeout < 0)
            return NULL;

        int ticks = 0;
        /* 算出设定的timeout需要tick多少次 */
        if(timeout < SI)//不足一个tick按照一个tick来处理
            ticks = 1;
        else
            ticks = timeout/SI;

        /* 计算出设定的时间多少圈后被触发 */
        int rotation = ticks/N;
        /* 计算新的定时器插入哪一个槽中 */
        int ts = (cur_slot + (ticks%N))%N;
        tw_timer *timer = new tw_timer(rotation,ts);
        if(!slots[ts]){
            std::cout << "Add timer,rotation is " << rotation << std::endl; 
            slots[ts] = timer;
        }else{//头插发将新定时器插入链表之中
            timer->next = slots[ts];
            slots[ts]->prev = timer;
            slots[ts] = timer;
        }
        return timer;
    }

    /* 删除目标定时器 */
    void del_timer(tw_timer* timer){
        if(!timer)
            return ;
        int ts = timer->time_slot;

        if(timer == slots[ts]){//目标定时器是头结点
            slots[ts] = slots[ts]->next;
            if(slots[ts]){
                slots[ts]->prev = NULL;
            }
            delete timer;
        }else{//不是头结点
            timer->prev->next = timer->next;
            if(timer->next){
                timer->next->prev = timer->prev;
            }
            delete timer;
        }
    }

    void tick(){
        /* 获得时间轮上当前槽的头结点 */
        tw_timer* tmp = slots[cur_slot];
        std::cout << "current slot is " << cur_slot << std::endl;
        while(tmp){
            std::cout << "tick the timer once " << std::endl;
            /* 轮数大于0，说明这轮没有超时 */
            if(tmp->rotation > 0){
                tmp->rotation--;
                tmp = tmp->next;//判断下一个节点
            }else{
                tmp->cb_func(tmp->user_data);
                if(tmp == slots[cur_slot]){//头结点到时候
                    std::cout << "delete header in cur_slot\n" << std::endl;
                    slots[cur_slot] = tmp->next;
                    delete tmp;
                    if(slots[cur_slot]){
                        slots[cur_slot]->prev = NULL;
                    }
                    tmp = slots[cur_slot];
                }else{
                    tmp->prev->next = tmp->next;
                    if(tmp->next){
                        tmp->next->prev = tmp->prev;
                    }
                    tw_timer* tmp2 = tmp->prev;
                    delete tmp;
                    tmp = tmp2;//向后转移
                }
            }
            /* 时间轮转动 */
            cur_slot = ++cur_slot%N;
        }
    }


private:
    static const int N = 60;//时间轮上槽的数目
    static const int SI = 1;//时间轮的槽每一秒时间轮转动一格
    tw_timer* slots[N];//时间轮的一圈是用数组来模拟的
    int cur_slot;//时间轮当前槽
};
#endif
