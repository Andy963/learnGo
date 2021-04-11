package add

//先执行全局变量,然后init函数
var Name string = "hello world"
var Age int = 10

func init() {
	Name = "I'm Name"
	Age = 29
}
