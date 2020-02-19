#include <iostream>
#include <string>
#define cout std::cout
typedef std::string string;

class Test{
    friend void print_data(Test &test);
public:
    Test(int data):
        data(data){};
    static void print_data(Test &test){
        cout << test.data << std::endl;
    }
private:
    int data;
};

int main()
{
    Test test(100);
    test.print_data(test);
    string s("Hello World!");
    cout << s << std::endl;
    return 0;
}

