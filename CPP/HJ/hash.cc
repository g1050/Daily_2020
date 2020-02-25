#include <iostream>

template <class T>
class hash{
public:
    void method(){
        std::cout << "Hello T" << std::endl;
    }
};

template<>//模板偏特化
class hash<char>{
public:
    void method(){
        std::cout << "Hello char" << std::endl;
    }
};

int main()
{
    hash<char> h;
    h.method();
    return 0;
}

