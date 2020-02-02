#include "Complex.h"

//global function 复数与复数想加
Complex operator+ (const Complex &x , const Complex& y)
{
    return Complex(x.real()+y.real(),x.imag()+y.imag());//临时对象 匿名对象
}

inline Complex& Complex::operator+= (const Complex&r)
{
    //不是local object就可以用reference
    return __doapl(this,r);
}

std::ostream& operator<< (std::ostream& os,const Complex& x)
{
    return os << "(" << x.real() << ","  << x.imag() << ")";
}

inline double Complex::real() const
{
    return re;
}

inline double Complex::imag() const
{
    return im;
}

Complex::Complex(double r ,double i ):re(r),im(i)//参数默认值,只需要声明就可以了，定义时候无需在写
{

}

inline Complex& __doapl(Complex *t,const Complex& r)
{
    t->re += r.re;
    t->im += r.im;
    return *t;
}

