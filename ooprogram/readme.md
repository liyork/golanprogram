演示面向对象特性，设计并实现了一个音乐播放器程序

示范以下的关键流程：
(1) 音乐库功能，使用者可以查看、添加和删除里面的音乐曲目；
(2) 播放音乐；
(3) 支持MP3和WAV，但也能随时扩展以支持更多的音乐类型(基于接口编程，但是go对接口有很大的弹性)；
(4) 退出程序

该程序在运行后进入一个循环，用于监听命令输入的状态。该程序将接受以下命令。
音乐库管理命令： lib，包括list/add/remove命令。
播放管理： play命令， play后带歌曲名参数。
退出程序： q命令。

cd golanprogram/ooprogram
export GOPATH=$PWD

cd src/main
go run mplayer.go
lib add HugeStone MJ ~/MusicLib/hs.mp3 MP3
play HugeStone
lib list
lib view
q


面向对象，就不用很多包进行区分了。
面向对象三特征：封装(使用对象)、继承(不用显示，直接组合)、多态(只需要实现方法，灵活)