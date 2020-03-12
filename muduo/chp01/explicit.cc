#include <iostream>
using namespace std;

class Test1{
public:
    Test1(int number){
        (void)number;
    }
};

class Test2{
public:
    explicit Test2(int number){
        (void)number;
    }
};

int main()
{
    Test1 t1 = 2020;
    /* Test2 t2 = 2020; */
    Test2 t3(2020);
    return 0;
}

