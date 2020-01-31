#include <iostream>

int main()
{
    double a = 3.14;
    const int &b = a;
    std::cout << b << std::endl;

    /* int &c = a;//不允许 */

    int x = 666;
    int &y = x;
    const int &z = x;
    y = 888;
    std::cout << z << std::endl;


    return 0;
}

