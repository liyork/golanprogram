排序算法的比较程序

用法如下所示：
USAGE: sorter –i <in> –o <out> –a <qsort|bubblesort>

具体的执行过程如下：
$ ./sorter –i in.dat –o out.dat –a qsort
The sorting process costs 10us to complete.

目录的结构如下：
<sorter>
    ├─<src>
        ├─<sorter> #主程序包
            ├─sorter.go #主程序
        ├─<algorithms> #算法包
            ├─<qsort> 
                ├─qsort.go #快速排序
                ├─qsort_test.go
            ├─<bubblesort>
                ├─bubblesort.go #冒泡排序
                ├─bubblesort_test.go
    ├─<pkg>
    ├─<bin>
    
cd golanprogram/sorter
export GOPATH=$PWD

cd golanprogram/sorter/src/sorter
go build sorter.go
./sorter -i unsorted.dat -o sorted.dat -a bubblesort


cd golanprogram/sorter
构建并单元测试
go build algorithms/qsort
go build algorithms/bubblesort
go test algorithms/qsort
go test algorithms/bubblesort
安装
go install algorithms/qsort
go install algorithms/bubblesort
go build sorter
go install sorter
生成bin和pkg目录，pkg下有bubblesort.a和qsort.a，bin下有sorter的二进制可执行我呢间

cp src/sorter/unsorted.dat bin
cd bin
./sorter -i unsorted.dat -o sorted.dat -a qsort
cat sorted.dat
./sorter -i unsorted.dat -o sorted.dat -a bubblesort

