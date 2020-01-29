#include "Complex.h"

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
