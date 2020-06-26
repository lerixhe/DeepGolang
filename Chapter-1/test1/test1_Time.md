# 程序运行时间

## 分析程序的运行时间

只需在前面+time命令

```shell
 time go run test1_Time.go 
```

运行结果

    real    0m0.371s 
    //程序从开始到结束的运行总时间

    user    0m0.156s
    //程序在用户态度过的时间

    sys     0m0.359s
    //程序在内核态度过的时间

## 用户态与内核态

CPU有不同等级的特权指令，特权最高的指令对硬件的权限越大。
在操作系统的进程中，会调用这些指令，拥有特权指令的进程所处的状态就是内核态，拥有低权限指令的进程所处的状态就是用户态。
用户态拥有自己的用户空间，只能在这个范围内，执行低特权等级的指令。
这样就能保证，系统程序与应用程序和平相处，而不是混乱。

比如处理用户逻辑，只需要在用户态执行，调用cpu、物理内存、网卡等硬件资源的内核逻辑，需要在内核态进行。

拓展：

用户态切换到内核态的唯一途径——>中断/异常/陷入
内核态切换到用户态的途径——>设置程序状态字

## real >= user + sys ?

一般来说是对的，因为real time实际要包含阻塞、切换等额外花费些时间。但这里陷入出现real time反而更短的情况，查了下可能是多核cpu导致的。

## 更详细的信息 usr/bin/time

事实上，直接调用time二进制文件，可以获得更详细的信息。
显示CPU占用率，内存使用情况、进程切换、文件系统IO、socket情况

```shell
/usr/bin/time -v go run test1_Time.go
```

        lerix@DESKTOP-VS7IE94:~/DeepGolang/Chapter-1$ /usr/bin/time go run test1_Time.go
        2
        0.10user 0.28system 0:00.35elapsed 111%CPU (0avgtext+0avgdata 44276maxresident)k
        0inputs+0outputs (0major+23909minor)pagefaults 0swaps
        lerix@DESKTOP-VS7IE94:~/DeepGolang/Chapter-1$ /usr/bin/time -v go run test1_Time.go
        2
        Command being timed: "go run test1_Time.go"
        User time (seconds): 0.20
        System time (seconds): 0.37
        Percent of CPU this job got: 168%
        Elapsed (wall clock) time (h:mm:ss or m:ss): 0:00.34
        Average shared text size (kbytes): 0
        Average unshared data size (kbytes): 0
        Average stack size (kbytes): 0
        Average total size (kbytes): 0
        Maximum resident set size (kbytes): 44228
        Average resident set size (kbytes): 0
        Major (requiring I/O) page faults: 0
        Minor (reclaiming a frame) page faults: 23690
        Voluntary context switches: 0
        Involuntary context switches: 0
        Swaps: 0
        File system inputs: 0
        File system outputs: 0
        Socket messages sent: 0
        Socket messages received: 0
        Signals delivered: 0
        Page size (bytes): 4096
        Exit status: 0
