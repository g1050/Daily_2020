#ifndef _COMPLEX_
#define _COMPLEX_

/* template <typename T>//模板 */
class Complex
{
public:
    Complex(double r = 0,double i = 0);//参数默认值
    ~Complex() {}

public:
    double real() const;
    double imag() const;

private:
    /* T re,im;//实部和虚部 */
    double re,im;

};

#endif

