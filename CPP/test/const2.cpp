#include <iostream>

int main()
{
    int const a = 1;
    int b = 2;
    const int c = 3;
    int d = 4;

    int * const p = &b;
    *p = 888;
    /* p =  &d; */

    const int *q = &a;
    int const *r = &a;
    q = &c;
    r = &c;
    /* *r = 888; */
    /* *q = 888; */
    /* a = 888; */
    return 0;
}

