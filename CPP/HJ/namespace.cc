#include <iostream>
namespace hh{

void test1(){
    std::cout << "Hello C++!" << std::endl;
} 

void test2(){
    std::cout << "Hi C++!" << std::endl;
}
} 

namespace kk{

void test1(){
    std::cout << "Hello World!" << std::endl;
} 

void test2(){
    std::cout << "Hi World!" << std::endl;
}
} 

int main()
{
    kk::test1();
    kk::test2();
    hh::test1();
    hh::test2();
    return 0;
}

