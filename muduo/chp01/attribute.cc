#include <iostream>
#include <cstdarg>
using namespace std;

void func() __attribute__((noreturn));

/* param attribute */
struct  student{
    int score;
    char sex;
} __attribute__((aligned(16)));

/* All member get togother. */
struct packed_struct{
    int number;
    student stu;
} __attribute__((packed));

/* Not togother. */
struct unpacked_struct{
    int number;
    student stu;
};

/* func attribute */
extern void func1(const char *format,...) __attribute__((format(printf,1,2)));

void func1(const char *format,...){
    va_list list;
    va_start(list,format);
    printf(format,list);
    va_end(list);
}


int main()
{
    /* param */
    cout << "sizeof(student): " << sizeof(student) << endl;
    cout << "sizeof(packed_struct)" << sizeof(packed_struct) << endl;
    cout << "sizeof(unpacked_struct)" << sizeof(unpacked_struct) << endl;
    /* func */
    func1("%d-%d-%d\n",2020,3,12);
    return 0;
}

/* void  func(){ */
/*     cout << "test1" << endl; */
/* } */
