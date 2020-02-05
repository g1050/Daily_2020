#include <iostream>
using namespace std;
class Library{
public:
        virtual ~Library(){}    //父类析构函数写作虚函数

public:
        void run(){
            step1();

            if(step2()){
                step3();
            }

            for(int i = 0;i<4;i++){
                step4();
            }
            
            step5();
        }

protected:
        virtual void step1(){
            cout << "step1" << endl;
        }
        void step3(){
            cout << "step3" << endl;
        }
        void step5(){
            cout << "step5" << endl;
        }
        virtual bool step2() = 0;
        virtual void step4() = 0;//变化

        

};

class Application: public Library{
public:
    Application(){};
    ~Application(){
        cout << "Application析构函数" << endl;
    }

protected:

    /* virtual void step1(){ */
    /*     cout << "子类拥有自己的想法step1" << endl; */
    /* } */

    virtual bool step2(){
        cout << "step2为真" << endl;
        return true;
    }

    virtual void step4(){
        cout << "step4" << endl;
    }
};

int main()
{
    Library *pointer = new Application();
    pointer->run();
    delete pointer;
    return 0;
}

