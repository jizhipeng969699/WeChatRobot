# groutine 调度器的GMP模型的设计思想

## GMP 模型的简介：
- G goroutine 协程
- P processor 处理器 用于处理 G
- M thread 线程 内核线程 物理线程

- 全局队列  存放等待运行的G
- P的本地队列 存放等待运行的G
- - 最大数量不超过256 G
  - 创建一个G  优先放到某个P的本地队列中，如果满了会放到全局队列中
- P列表
- - 程序启动是创建
  - 最多有 GOMAXPROCS 个  可以配置
- M列表： 当前操作系统分配到当前Go程序的内核线程数量
- P和M 的数量：
- - P的数量问题  环境变量设置$GOMAXPROCS  runtime.GOMAXPROCS 设置
  - M的数量问题   Go语言本身限定M的最大数量是10000   runtime/debug  SetMaxThreads 函数设置
  - - 有一个M阻塞，会创建一个新的M
    - 有m空闲，会回收或者睡眠


## 调度器的设计策略：
- work stealing 偷取机制   ， p 本地队列没有G  去其他p中偷取G 其他也没有 去全局队列取G
- hand off 握手机制    p执行G时发生阻塞， M释放绑定P ，让这个P 去执行其他G ， 阻塞G会和M绑定，M执行完G后会销毁或睡眠


## “go func{} 经历了什么过程”：
- 

## GMP 模型的简介：
- 
