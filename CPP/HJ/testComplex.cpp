#include <iostream>
#include "Complex.h"
int main()
{
    Complex c(1,2);
    std::cout << c << std::endl;
    Complex c2(2,3);
    std::cout << c2 << std::endl;
    std::cout << c+c2 << std::endl;
    return 0;
}

