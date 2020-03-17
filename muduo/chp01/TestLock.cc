//互斥锁测试代码
#include <muduo/base/CountDownLatch.h>
#include <muduo/base/Mutex.h>
#include <muduo/base/Thread.h>
#include <muduo/base/Timestamp.h>

#include <boost/bind.hpp>
#include <boost/ptr_container/ptr_vector.hpp>
#include <vector>
#include <stdio.h>

using namespace muduo;
using namespace std;

MutexLock g_mutex;//声明锁对象
vector<int> g_vec;//int动态数组(向量)
const int kCount = 10*1000*1000;//常量1千万

void threadFunc()
{
    for (int i = 0; i < kCount; ++i)
    {
        MutexLockGuard lock(g_mutex);//使用锁
        g_vec.push_back(i);//往向量中插入1000w个整数
    }
}

int main()
{
    const int kMaxThreads = 8;//最多8个线程
    g_vec.reserve(kMaxThreads * kCount);//预留8千万个整数(这个所占的内存空间有300多M)

    Timestamp start(Timestamp::now());//当前时间戳
    for (int i = 0; i < kCount; ++i)
    {
        g_vec.push_back(i);//往向量中插入1000w个整数
    }
    //输出插入这么多个数的时间
    printf("single thread without lock %f\n", timeDifference(Timestamp::now(), start));

    start = Timestamp::now();//更新当前时间戳
    threadFunc();//调用上面的函数
    //和上面一样，计算下插入这么多个数的时间
    printf("single thread with lock %f\n", timeDifference(Timestamp::now(), start));

    for (int nthreads = 1; nthreads < kMaxThreads; ++nthreads)
    {//ptr_vector指针的vector
        boost::ptr_vector<Thread> threads;
        g_vec.clear();//先清除g_vec向量
        start = Timestamp::now();//更新当前时间戳
        for (int i = 0; i < nthreads; ++i)
        {
            threads.push_back(new Thread(&threadFunc));//创建线程
            threads.back().start();//启动线程
        }
        for (int i = 0; i < nthreads; ++i)
        {
            threads[i].join();
        }
        //分别输出1到8个线程执行插入操作的时间
        printf("%d thread(s) with lock %f\n", nthreads, timeDifference(Timestamp::now(), start));
    }
}
}
}}
}
}
}
}
