Go命令行工具的革命性之处在于彻底消除了工程文件的概念，完全用目录结构和包名来推导工程结构和构建顺序。

展示Go语言的基本工程管理方法

场景：我们需要开发一个基于命令行的计算器程序。
基本用法
$ calc help
USAGE: calc command [arguments] ...
The commands are:
sqrt Square root of a non-negative value.
add Addition of two values.

$ calc sqrt 4 # 开根号
2
$ calc add 1 2 # 加法
3

```
<calcproj>
├─<src>
    ├─<calc> #可执行程序
        ├─calc.go  
    ├─<simplemath> #算法库，每个command对应一个同名go文件
        ├─add.go
        ├─add_test.go #表示对于add.go的单元测试
        ├─sqrt.go
        ├─sqrt_test.go
├─<bin>
├─<pkg> ＃包将被安装到此处
```

cd calcproj
export GOPATH=$PWD
cd bin
go build calc #生成可执行文件

执行：
./calc sqrt 4
./calc add 1 2

单元测试：
go test simplemath

GDB调试
gdb calc