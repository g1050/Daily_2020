#include <iostream>
using namespace std;

class Fraction{
public:
    Fraction(int num,int den = 1)
        :m_numreator(num),
        m_denominator(den){};

    /* 转换函数一般写作常函数 */
    operator double() const{//不用写返回值类型，转换函数
        return (double)(m_numreator/m_denominator);
    }

private:
    int m_numreator;//分子
    int m_denominator;//分母
};

int main()
{
    Fraction f(6,3);
    double a = (double)f;
    cout << a << endl;
    return 0;
}

