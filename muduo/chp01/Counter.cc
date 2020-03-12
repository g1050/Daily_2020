#include <iostream>
#include <boost/core/noncopyable.hpp>

class MutexLock {
public:
    MutexLock(){}
};

/* A thread-safe counter. */
class Counter :boost::noncopyable{
public:
    Counter() :value_(0){}
    int64_t value() const;
    int64_t getAndIncrease();

private:
    int64_t value_;
    mutable MutexLock motex_;
};

int64_t Counter::value() const{
    MutexLockGu;
}
int main()
{

    /* std::cout << sizeof(int) << std::endl; */
    /* std::cout << sizeof(int8_t) << std::endl; */
    /* std::cout << sizeof(int16_t) << std::endl; */
    /* std::cout << sizeof(int32_t) << std::endl; */
    /* std::cout << sizeof(int64_t) << std::endl; */

    return 0;
}

