#include <iostream>

class Test{
public:
   Test():times(0){}
   ~Test(){}
   void Output() const{
       std::cout << "Hello World\n" << std::endl;
       times++;
   }
   int getTimes() const{
       return times;
   }
private:
   mutable int times;//mutable变量在const中是可以修改的
};

int main()
{
    Test test;
    test.Output();
    std::cout << test.getTimes() << std::endl;
    return 0;
}

