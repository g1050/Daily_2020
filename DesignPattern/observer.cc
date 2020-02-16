#include <iostream>
#include <string>
#include <vector>

class IProgress
{
public:
    virtual void Doprogress(float value) = 0;//纯虚函数没有具体的实现，而要留到子类去实现
    virtual ~IProgress(){};//虚函数有实现，子类可以重载
};

class Form
{

};

class FileSpliter;

//多继承推荐用法，主继承和辅继承
class MainForm : public Form,public IProgress
{
    std::string file_path;//文件路径
    int file_num;//文件分割数目
    
    ProgressBar *progressbar();
    //因为MainForm和progressbar是一个紧耦合的关系
    
public:
    void button_click()
    {
        file_path = get_file_path();
        file_num = get_file_num();

        FileSpliter spliter(file_path,file_num,this);
        spliter.split();
    }

    virtual void Doprogress(float value)
    {
        progressbar->setValue(value);
    }

};

class FileSpliter
{
    std::string file_path;    
    int file_num;
    //如果说提供一个进度条的话怎么做
    /* ProgessBar *progressbar;//具体的通知事件 */
    IProgress *m_iprogress;//抽象的通知机制
    /* FileSpliter不再紧耦合界面类，而是依赖于抽象 */
    std::vector<IProgress> *m_iprogressVector;//支持多个观察者

public:
    FileSpliter(const std::string& path,int num,IProgress *iprogress):
        file_path(path),
        file_num(num),
        m_iprogress(iprogress){};  
    void split()
    {
        //1.读取大文件
        
        //2.分批次写入
        for(int i = 0;i<file_num;i++)
        {
            for()//遍历vector取到所有通知的东西
            if(m_iprogress == NULL) ;//没有通知机制
            else
            {
                m_iprogress->Doprogress(0.99);
            }
        }
    }

    void AddIprogress(){} 
    void DeleteIprogress(){}
};

int main()
{

    return 0;
}

