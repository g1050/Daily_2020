#include <iostream>

//迭代器和智能指针
//智能指针有++的操作
template <class T>
class shared_ptr{
public:
    shared_ptr(T*p)
        :px(p){}
public:
    T& operator*() const{
        return *px;
    }
    T* operator->() const{
        return px;
    }
private:
    T* px;
    long* pn; 
};

struct Foo{
public:
    void method(){
        std::cout << "Hello World!" << std::endl;
    }
};

int main()
{
    shared_ptr<Foo> sp(new Foo);
    sp->method();

    Foo f = *sp;
    f.method();
    return 0;
}

