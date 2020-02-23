#include <iostream>
#include "readconfig.h"
using namespace std;
int main()
{
    ReadConfig readconfig("test.config");
    if(!readconfig.readFile()){
        cout << "error" << endl;
    }
    readconfig.print();
    return 0;
}

