#include "http.h"
using namespace std;

int main(int argc,char *argv[])
{
    if(argc <= 2){
        cout << "Usage:" << argv[0] << "ip_address port_number" << endl;
        return 1;
    }
    
    const char *ip = argv[1];
    const int port = atoi(argv[2]);
    
    struct sockaddr_in address;
    bzero( &address,sizeof(address));
    address.sin_family = AF_INET;
    inet_pton(AF_INET,ip,&address.sin_addr);
    address.sin_port = htons(port);

    int listenfd = socket(PF_INET,SOCK_STREAM,0);
    assert(listenfd >= 0);
    int ret = bind(listenfd,(struct sockaddr*)&address,sizeof(address));
    assert(ret != -1);
    struct sockaddr_in client_address;
    socklen_t client_addrlength = sizeof(client_address);
    int fd = accept(listenfd,(struct sockaddr*)&client_address,&client_addrlength );

    if(fd < 0){
        cout << "errorno is " << errno << endl;
    }else{
        char buffer[BUFFER_SIZE];//读缓冲区
        memset(buffer,'\0',BUFFER_SIZE);
        int data_read = 0;
        int read_index = 0;//已经读取到的数据
        int checked_index = 0;//当前已经分析多少数据
        int start_line = 0;//记录行在buffer中的起始数据
        
        /* 初始化主状态机的状态 */
        CHECK_STATE checkstate = CHECK_STATE_REQUESTLINE;
        while(1){//循环读取数据并且分析
           data_read = recv(fd,buffer + read_index,BUFFER_SIZE - read_index,0);
           if(data_read == -1){
               cout << "read failed" << endl;
           }else if(data_read == 0){
                cout << "Remote client has closedd the connection " << endl;
           }
           read_index += data_read;

           /* 分析目前已获得的所有客户数据 */
           HTTP_CODE result = parse_content(buffer,checked_index,checkstate,read_index,start_line);
           if(result == NO_REQUEST){
               continue;
           }else if(result == GET_REQUEST){
               send(fd,szret[0],strlen(szret[0]),0);
               break;
           }else{
               send(fd,szret[1],strlen(szret[1]),0);
               break;
           }
        }


    }

    close(fd);
    close(listenfd);
    cout << "Have a nice day! ^.^" << endl;

    return 0;
}

