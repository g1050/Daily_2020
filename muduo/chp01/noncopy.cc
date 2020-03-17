#include <iostream>
#include <boost/noncopyable.hpp>
using namespace std;

class Test1 : boost::noncopyable{
public:
    Test1(){}
};

class Test2 {

};

int main()
{

    /* Test1 t1 ; */    
    /* Test1 t2(t1); */
    /* Test1 t3 = t1; */

    /* 调用了默认拷贝和默认赋值函数 */
    Test2 t4;
    Test2 t5(t4);
    Test2 t6 = t4;
    return 0;
}

