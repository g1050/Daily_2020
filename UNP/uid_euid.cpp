#include <iostream>
#include <unistd.h>
using namespace std;
int main()
{
    uid_t uid = getuid();
    uid_t euid = geteuid();

    cout << "uid:" << uid << endl;
    cout << "euid:" << euid << endl;
    return 0;
}

