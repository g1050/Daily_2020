#include <iostream>

//仿函数,class重载()小括号就是仿函数
template<class T>
struct identity{//认为同一个东西
    const T& operator()(const T&x) const{
        return x;
    }
};

template <class Pair>
struct select1st{
    const typename Pair::first_type& operator()(const Pair& x)const {
        return x.first;
    }
};

int main()
{
    return 0;
}

