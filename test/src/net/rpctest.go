package main

//net/rpc包

//一个对象中只有满足如下这些条件的方法，才能被 RPC 服务端设置为可供远程访问：
// 必须是在对象外部可公开调用的方法（首字母大写）；
// 必须有两个参数，且参数的类型都必须是包外部可以访问的类型或者是Go内建支持的类
//型；
// 第二个参数必须是一个指针；
// 方法必须返回一个error类型的值
//func (t *T) MethodName(argType T1, replyType *T2) error
