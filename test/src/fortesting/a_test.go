package fortesting

import "testing"

//Go本身提供了一套轻量级的测试框架。符合规则的测试代码会在运行测试时被自动识别并执行。
//在需要测试的包下面创建"target_test.go"文件，target为目标测试go文件

//功能测试函数，以Test开头
//功能测试：go test fortesting 或进入fortesting执行go test
func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 { //2，则有错误
		t.Errorf("Add(1,2) failed. Got %d, expected 3.", r) //打印了错误信息后中止测试
	}
}

//性能测试
//cd fortesting 后 go test-test.bench a_test.go???
func BenchmarkAdd(b *testing.B) {
	b.StopTimer() //暂停计时器
	//费时准备操作
	doPreparation()
	b.StopTimer() //开启计时器
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func doPreparation() {

}
