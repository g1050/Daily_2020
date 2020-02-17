#include <iostream>

//业务操作
class Stream{
public:
    virtual char Read(int number) = 0;
    virtual void Seek(int position) = 0;
    virtual void Write(int data)= 0;

    virtual ~Stream(){};
};

/**************************主体类***************************/
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

    virtual void Write(int data){
        //写文件流
        return ;
    }
};

//其他文件流都是读写定位的操作
class NetworkStream :public Stream{

};

class MemoryStream : public Stream{

};

/**************************扩展操作***************************/
//对公共成员向上提，在这里做一个中间类
class DecoratorStream: public Stream{
protected:
    Stream *stream;//
    DecoratorStream(Stream *stm):
        stream(stm){}
};

//对流进行加密操作
class CryptoStream: public DecoratorStream{
    Stream* stream;
public:
    CryptoStream(Stream* stm):
        DecoratorStream(stm){}

    virtual char Read(int number){//继承加组合,继承是为了Override　Read方法,组合是...
        stream->Read(number);
        //进行额外的加密操作
        return 'a';
    }

    virtual void Seek(int position){
        //进行额外加密操作
        stream->Seek(position);
    }

    virtual void Write(int data){
       //进行额外的加密操作 
       stream->Write(data);
    }

};
//对流进行缓冲
class BufferedStream: public DecoratorStream{
    Stream *stream;
public:
    BufferedStream(Stream *stm):
        DecoratorStream(stm){};

    //...
    virtual char Read(int number){
        //对流进行缓冲操作
        return 'a';
    }

    virtual void Seek(int position){
        //进行额外加密操作
        stream->Seek(position);
    }

    virtual void Write(int data){
       //进行额外的加密操作 
       stream->Write(data);
    }


};

int main()
{
    FileStream *fs1 = new FileStream();
    CryptoStream *fs2 = new CryptoStream(fs1);//加密
    BufferedStream *fs3 = new BufferedStream(fs1);//缓冲
    BufferedStream *fs4 = new BufferedStream(fs2);//加密且缓冲

    //这就是组合由于继承的地方
    return 0;
}

