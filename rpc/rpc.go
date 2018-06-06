package rpcdemo

// 我们的RPC最终是怎么被客户所调用的呢？Service.Method的方法
type DemoService struct{}

//Go语言语法框架对于RPC的要求就是要有args和result两个参数
//func (DemoService) Div(args, result) error {
//
//}
