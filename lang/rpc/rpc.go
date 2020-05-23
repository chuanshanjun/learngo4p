package rpcdemo

import "errors"

// Service.Method

type DemoService struct {
}

// Args要大写，里面的参数也要大写
type Args struct {
	A, B int
}

// 参数一定是两个，一个是Args，另外一个是result，且Result一定要是*号，因为我们要把它写进去
// 两个参数一个是输入，一个是输出
// Args加不加信号都没有关系，如果加星号的话，就直接传递指针，不加星号的话会值拷贝
// 关键是result一定是要指针类型
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
