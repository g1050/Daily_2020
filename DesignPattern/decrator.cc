#include <iostream>

//业务操作
class Stream{
public:
    virtual char Read(int number) = 0;
    virtual void Seek(int position) = 0;
    virtual void Write(char data)= 0;

    virtual ~Stream(){};
};

/**************************主题类***************************/
//主体类FileStream,应该和拓展功能向两个方向发展，最后再组合在一起
//可能还有网络流、内存流等等各种流，但是基本操作是一样的都是Read,Seek,Write
class FileStream : public Stream{
    virtual char Read(int number){
        //读文件流
        return 'a';
    }

    virtual void Seek(int position){
        //定位文件流
        return ;
    }

    virtual void write(int data){
        //写文件流
        return ;
    }
};

class NetworkStream :public Stream{

};

class MemoryStream : public Stream{

};

/**************************扩展操作***************************/
//对流进行加密操作
class CryptoStream: public Stream{
    Stream* stream;
public:
    CryptoStream(Stream* stm):
        stream(stm){};

    virtual char Read(int number){//继承加组合,继承是为了Override　Read方法,组合是...
        stream->Read(number);
        //进行额外的加密操作
        return 'a';
    }

    virtual void Seek(int position){
        //进行额外加密操作
        stream->Seek(position);
    }

    virtual void Write() 
};
int main()
{
    return 0;
}

