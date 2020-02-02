#include "MyString.h"

int main()
{
    MyString s1();
    MyString s2("Hello World!");
    MyString s3(s1);//也是一个拷贝动作,拷贝构造函数

    std::cout << s3 << std::endl;
    s3 = s2;//s3已经出现了，此时要对s3赋值,拷贝动作,拷贝赋值函数
    std::cout << s2 << std::endl;

    return 0;
}

