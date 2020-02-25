#ifndef THREADPOOL_H
#define THREADPOOL_H

#include <list>
#include <cstdlib>
#include <cstdio>
#include <exception>
#include <pthread.h>
#include <semaphore.h>

#include "locker.h"
template <class T>
class threadpool{
public:
    threadpool(int threadnum = 8,int max_requests = 10000);
    ~threadpool();
    bool append(T* request);
private:
    static void* worker(void *arg);
    void run();
private:
    int m_pthread_number;
    int m_max_request;
    pthread_t* m_threads;
    std::list<T*> m_workerqueue;
    locker m_queuelocker;
    sem m_queuestat;
    bool m_stop;
};

template< class T>
threadpool<T>::threadpool(int thread_number,int max_requests):
    m_pthread_number(thread_number),
    m_max_request(max_requests),
    m_stop(false),
    m_threads(NULL){
        if((thread_number <= 0) || (max_requests <= 0)){
            throw std::exception();
        }
        m_threads = new pthread_t[m_pthread_number];
        if(!m_threads){
            throw std::exception();
        }

        //Loop create threads.
        for(int i = 0;i<thread_number;i++){
            printf("create the %dth thread\n",i+1);
            //0 means sucess,-1 means fail.Pointer this was passed to the worker.Worker is static.  
            if(pthread_create(m_threads +  i, NULL , worker,this) != 0){
                delete []m_threads;
                throw std::exception();
            }

        //Make all the threads detached.
            if(pthread_detach(m_threads[i])){
                delete []m_threads;
                throw std::exception();
            }
        }
    }

template <class T>
threadpool<T>::~threadpool(){
    delete []m_threads;
    m_stop = true;
}

//Add what you read to the workerqueue for parsing it.
template <class T>
bool threadpool<T>::append(T *request){
    m_queuelocker.lock();//request in worker queue
    //The queue is full,so it will be error.
    if(m_workerqueue.size() > m_max_request){
        m_queuelocker.unlock();
        return false;
    }

    m_workerqueue.push_back(request);
    m_queuelocker.unlock();
    m_queuestat.post();//Count the number of task?
    this->worker(this);
    return true;
}


template <class T>
void* threadpool<T>::worker(void *arg){
    std::cout << "Worker is working" << std::endl;
    threadpool* pool = (threadpool*)arg;
    pool->run();
    return pool;
}

template <class T>
void threadpool<T>::run(){
    while(!m_stop){
        std::cout << "thread pool is run" << std::endl;
        m_queuestat.wait();
        m_queuelocker.lock();
        if(m_workerqueue.empty()){
            std::cout << "Queue is empty!" << std::endl;
            m_queuelocker.unlock();
            continue;
        }

        T* request = m_workerqueue.front();
        m_workerqueue.pop_front();
        m_queuelocker.unlock();
        if(!request){
            continue;
        }
        request->process();
    }
}

#endif
