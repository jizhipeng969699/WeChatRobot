# golang GMP 模型和设计思想


## 进程
- 内存分配的最小单位 ， 进程间通信：tcp udp http unixdomain

## 线程（thread）
- CPU 分配时间轮片的最小执行单位，
  - 是进程的一个实体，  
  - 一个进程可以创建多个线程 
  - 线程通信 ： 共享内存，上线文切换资源消耗较小
  - 内存大小 G 级别
  - 内核态线程  物理线程
    
## 协程（co-routine）
- 用户态执行单位
- 更加轻量级，更灵活
- 内存大小 M 级别
- M : 1  没有利用多cpu
- 1 ： 1  内核态用户态切换浪费资源
- M ：N  多协程多进程 过度依赖 调度器

## go协程 （go routine）
- 依托GMP模型和设计思想 更灵活 切换成本更低
- 内存优化 4kb 级别  可以大量开辟



# GMP

## G （goroutine） go协程

## P （Processer） go调度器
- runtime.GOMaxProcess  设置调度器数量
- $GOMAXPROCESS 

## M  (thread) 内核线程



## 全局G队列
## p调度器G队列


- 一个 M 绑定一个 P 
- 一个 P 有一个G队列 

>M0 启动第一个G 也就是main,
> 
>M1 绑定 P 创建G队列，创建G0 
> 
>G0负责找G ， 先从本队列找，没有就去全局G队列找（从全局队列一次拿 min[ len GQ / maxprocess +1 , len GQ /2 ] ）   ，全局队列没有
>就去其他G队列偷   （线程服用） 一次偷这个G队列的一半
> 
> 当M执行G时发生syscall 调用发生阻塞， M会和当前G进行绑定，同时创建新的M或者从M睡眠队列找一个将自身的P队列 绑定给新的M
> 
> 
> 



 


