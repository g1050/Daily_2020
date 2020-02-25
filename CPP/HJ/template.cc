#include <iostream>

template <class T>
T& min(T &a,T& b){
    return a>b?b:a;
}

class Foo{
public:
    Foo(float data)
        :data(data){}
    bool operator> (const Foo& rhs) const {
        return data < rhs.data;
    }
    friend Foo& operator<< (std::ostream &os,Foo& f);
private:
    float data;
};

Foo& operator<< (std::ostream &os,Foo& f){
    os << f.data;
    return f;
}

int main()
{
    Foo f1 = Foo(1.1);
    Foo f2 = Foo(2.2);

    std::cout << min(f1,f2);

    return 0;
}

