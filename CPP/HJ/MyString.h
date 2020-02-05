#ifndef _MYSTRING_
#define _MYSTRING_
#include <iostream>
#include <cstdlib>
#include <cstring>

class MyString
{
public:
    MyString(const char * s = 0);//默认值是0
    MyString(const MyString &s);//拷贝构造函数
    char *getCStr() const{  return data;};
    ~MyString();

public:
    MyString& operator= (const MyString &s);//不是local object就可以return by reference

private:
    char *data;//字符串对象本身只有指针大小
};

#endif

