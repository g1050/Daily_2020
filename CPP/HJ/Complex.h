#ifndef _COMPLEX_
#define _COMPLEX_
#include <iostream>


/* template <typename T>//模板 */
class Complex
{
public:
    Complex(double r = 0,double i = 0);//参数默认值
    ~Complex() {}

public:
    inline double real() const;//没有改动值,所以是const
    inline double imag() const;
    Complex& operator+= (const Complex&);
    void f();

private:
    /* T re,im;//实部和虚部 */
    double re,im;

    friend Complex& __doapl(Complex *t,const Complex& r);
    
};
std::ostream& operator<< (std::ostream& os,const Complex& x);
Complex  operator+  (const Complex& x,const Complex& y);

#endif

