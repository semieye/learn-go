# goroutine

goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。

    go hello(a, b, c)

### runtime包中有几个处理goroutine的函数：

# Goexit

退出当前执行的goroutine，但是defer函数还会继续调用

# Gosched

让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

# NumCPU

返回 CPU 核数量

# NumGoroutine

返回正在执行和排队的任务总数

# GOMAXPROCS

用来设置可以并行计算的CPU核数的最大值，并返回之前的值。