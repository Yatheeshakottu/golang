type MyStruct struct {
	Val int
}

func myfunc() MyStruct {
	return MyStruct{Val: 1}
}

func myfunc() *MyStruct {
	return &MyStruct{}
}

func myfunc(s *MyStruct) {
	s.Val = 1
}
/*
1.first function returns a copy of MyStruct (or) the first function  returns a  new instance of MyStruct with "val" set to be 1
2.the second function is a pointer is a struct value passed to a function (or) the second function  returns a pointer to a new instance of MyStruct
3.The third function myfunc(s *MyStruct) takes a pointer to a MyStruct instance as an argument and sets its Val field to 1.