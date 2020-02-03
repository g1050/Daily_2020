#include <iostream>
namespace space1
{
    int a = 10;
}

namespace space2
{
    int a = 20;
}

using namespace std;
int main()
{
    cout << space1::a << endl;
    cout << space2::a << endl;
}

