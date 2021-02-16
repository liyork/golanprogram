棋牌游戏服务器的例子展现Go语言并发编程

详细需求:
需要支持玩家进行下面的基本操作：
登录游戏
查看房间列表
创建房间
加入房间
进行游戏
房间内聊天
游戏完成，退出房间
退出登录


将该示例划分为以下子系统：
玩家会话管理系统，用于管理每一位登录的玩家，包括玩家信息和玩家状态
大厅管理
房间管理系统，创建、管理和销毁每一个房间
游戏会话管理系统，管理房间内的所有动作，包括游戏进程和房间内聊天
聊天管理系统，用于接收管理员的广播

只实现了最基础的会话管理系统和聊天管理系统。因为
它们足以展示以下的技术问题：
goroutine生命周期管理
goroutine之间的通信
共享资源访问控制

目录结构
<cgss>
├─<src>
    ├─<cg>
        ├─center.go
        ├─centerclient.go
        ├─player.go
    ├─<ipc> //IPC（进程间通信）框架的目的:封装通信包的编码细节
        ├─server.go
        ├─client.go
        ├─ipc_test.go
    ├─cgss.go
    
运行
cd concurrentprogram
export GOPATH=$PWD

go run cgss.go
login Tom 1 101
login Jerry 2 321
listplayer
send Hello everybody.
logout Tom
listplayer
send Hello the people online.
logout Jerry
listplayer
q

#问题；
Command> send xxx
server resp: &{200 }
Command> j received message: xxx
t received message: xxx

产生这种现象的问题并不是并发，因为client发送send到server都是同步执行，而且还是用的channel，所以回来的消息也只有一个。
问题应该是player的print和主程序的Command>快慢了，应该是player接收到了，然后server这边也从channel的写事件返回了，但是
player还没来及打印，主线程就打印了Command>
