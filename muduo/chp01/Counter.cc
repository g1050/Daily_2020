#include <iostream>
#include <boost/noncopyable.hpp>
#include <muduo/base/Mutex.h>

/* A thread-safe counter. */
class Counter :boost::noncopyable{
public:
    Counter() :value_(0){}
    int64_t value() const;
    int64_t getAndIncrease();

private:
    int64_t value_;
    mutable muduo::MutexLock mutex_;
};

int64_t Counter::value() const{
    muduo::MutexLockGuard lock(mutex_);//自动析构,析构前会释放锁子
    return value_;
}

int64_t Counter::getAndIncrease(){
    muduo::MutexLockGuard lock(mutex_);
    int64_t ret = value_++;
    return ret;
}

int main()
{
    Counter counter;
    counter.getAndIncrease();
    counter.getAndIncrease();
    counter.getAndIncrease();
    counter.getAndIncrease();
    std::cout << counter.value() << std::endl;

    return 0;    
}

