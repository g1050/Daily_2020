#include <iostream>

int main(int argc,char **argv)
{
    /* Tell the compiler I now it,and i don't care it. */
    (void)argc;
    (void)argv;
    std::cout << "Hello world" << std::endl;
    return 0;
}

