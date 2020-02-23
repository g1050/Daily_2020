#include <iostream>
#include <string>
using namespace std;

int main()
{
    string s("Hello!");
    cout << s.substr(2) << endl;
    cout << s.substr(2,-2) << endl;
    return 0;
}

