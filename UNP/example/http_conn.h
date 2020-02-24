#ifndef HTTP_CONN_H
#define HTTP_CONN_H

#include <unistd.h>
#include <signal.h>
#include <sys/types.h>
#include <sys/epoll.h>
#include <fcntl.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <assert.h>
#include <sys/stat.h>
#include <string.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/mman.h>
#include <stdarg.h>
#include <errno.h>
/* #include "lock.h" */

class http_conn{
public:
    /* 文件名最大长度 */
    static const int FILENAME_LEN = 200;
    /* 读缓冲区大小 */
    static const int READ_BUFFER_SIZE = 2040;
    /* 写缓冲区大小 */
    static const int WRITE_BUFFER_SIZE = 1024;
    /* HTTP请求方法,但我们支持GET */
    enum METHOD{
        GET = 0,POST,HEAD,PUT,DELETE,TRACE,OPTIONS,CONNECT,PATCH
    };
    
    /* 解析客户请求时候，主状态机的状态 */
    enum CHECK_STATE{
        CHECK_STATE_REQUESTLINE = 0,
        CHECK_STATE_HEADER,
        CHECK_STATE_CONTENT
    };

    /* 服务器处理HTTP请求的可能结果 */
    enum HTTP_CODE{
        NO_REQUEST,GET_REQUEST,BAD_REQUEST,NO_RESOURCE,
        FORBIDDEN_REQUEST,FILE_REQUEST,INTERNAL_ERROR,
        CLOSED_CONNECTION
    };
    
    /* 行的读取状态 */
    enum LINE_STATUS{
        LINE_OK = 0,LINE_BAD,LINE_OPEN
    };
    
public:
    http_conn(){}
    ~http_conn(){}
public:
    /* 初始化新接受的连接 */
    void init(int sockfd,const sockaddr_in& addr);
    /* 关闭连接 */
    void close_conn(bool real_close = true);
    /* 处理客户请求 */
    void process();
    /* 非阻塞读操作 */
    bool read();
    /* 非阻塞写操作 */
    bool write();
private:
    /* 初始化连接 */
    void init();
    /* 解析HTTP请求 */
    HTTP_CODE process_read();
    /* 填充HTTP应答 */
    bool process_write(HTTP_CODE ret);

    /* 下面这一组函数被process_read调用用来分析HTTP请求 */
    HTTP_CODE parse_request_line(char *text);
    HTTP_CODE parse_headers(char *text);
    HTTP_CODE parse_content(char *text);
    HTTP_CODE do_request();
    char* get_line(return m_read_buf + m_start_line;);
    LINE_STATUS parse_line();

    /* 下面这组函数被process_write调用用来填充HTTP应答 */
    void unmap();
    bool add_response(const char *format,...);
    bool add_content(const char* content);
    bool add_status_line(int status,const char *title);
    bool add_headers(int content_lenth);
    bool add_content_length(int content_length);
    bool add_linger();
    bool add_blank_line();
public:
    /* 所有socket上的事件都被注册到同一个epoll内核时间中,所以将epoll描述符号设置为静态的 */
    static int m_epollfd;
    /* 统计用户数量 */
    static int m_user_count;

private:
    /* 读HTPP连接的socket和对方的socket地址 */
    int m_sockfd;
    sockaddr_in m_address;

    /* 读缓冲区 */
    char m_read_buf[READ_BUFFER_SIZE];
    /* 标示读缓冲中已经读入客户端数据的最后一个字节的下一个位置 */
    int m_read_idx;
    /* 当前正在分析的字符在读缓冲区的位置 */
    int m_checked_idx;
    int m_start_line;   
    char m_write_buf[WRITE_BUFFER_SIZE];
    int m_write_idx;

    CHECK_STATE m_check_state;
    /* 请求方法 */
    METHOD m_method;

    char m_read_file[FILENAME_LEN];
    char *m_url;
    
    char* m_version;
    char *m_host;
    int m_content_length;
    bool m_linger;

    char* m_file_address;
    struct stat m_file_stat;
    struct iovec m_iv[2];
    int m_iv_count;
};

#endif
