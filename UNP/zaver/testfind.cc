#include <iostream>
#include <string>
using namespace std;

int main()
{
    string s("Hello!");
    cout << s.find('e') << endl;
    cout << s.find('#') << endl;
    cout << string::npos << endl;

    return 0;
}

