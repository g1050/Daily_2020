#include "MyString.h"


inline MyString::MyString(const char * s)//默认值是0
{
    if(s){//s 中含有内容
        data = new char[strlen(s)+1];//结尾有一个'\0'
        strcpy(data,s);
    }else{
        data = new char[1];
        *data = '\0';//空串
    }
}

inline MyString::MyString(const MyString &s)//拷贝构造函数
{
    data = new char[strlen(s.data) + 1];
    strcpy(data,s.data);
}

MyString::~MyString()//delete constructor
{
    delete []data;//释放时候调用多次析构函数
}

inline MyString& MyString::operator= (const MyString& s)//不是local object就可以return by reference
{//删除、申请、复制
    if(this == &s)
        return *this;

    if(data){
        delete []data;
    }

    data = new char[strlen(s.data) + 1];
    strcpy(data,s.data);
    return *this;//支持链式赋值
}

