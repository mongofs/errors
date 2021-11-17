// 这里是建议处理最终错误的模式，如果在api接口调用过程中抛出的错误统一处理的地方
// 不论是api类型服务还是脚本类型服务，应该在项目统一错误处理的地方进行错误处理，
// 错误处理原则 1. 不应该太过分散的处理 2.不应该出现未知错误，一旦出现未知错误表明程序
// 已经进入非常糟糕的阶段
// 对内部错误记录应该记录：调用堆栈，code ，httpCode，ext，
// 出现错误对外部调用应该返回：code，httpCode,ext
package errors_test

import (
	"fmt"

	"github.com/mongofs/errors"
)

// 创建具体错误，
func Example_handlerCreatedByEngineersErr() {
	// 关于error_code test
	errors.Register(100910,500,"user is not found ")
	defer errors.Flush()
	ErrUserNotFound := errors.NewCode(100910)

	// 模拟具体程序产生错误如何handler
	if errors.ErrWithCode == errors.ErrorType(ErrUserNotFound){
		// check the error cause
		fmt.Printf("error cause is  : %s \n" ,errors.Cause(ErrUserNotFound))
		// check this code
		fmt.Printf("error code is : %v \n" ,errors.Code(ErrUserNotFound))

		// output: error cause is  : user is not found
		// output: error code is  : 100910
	}
}



func getError()error{
	return errors.New("connection reset by peer ")
}


// 处理产生的第三方错误
func Example_handlerCreatedByThirdPartyErr() {
	var code = 10086 // 网络调用错误
	errors.Register(code,500,"网络调用错误")
	defer errors.Flush()
	err := getError();
	// 在你这边应该将第三方的错误进行归类
	withCode := errors.WithCode(err,code)
	fmt.Printf("%v\n\r",withCode)
	fmt.Printf("%v\n\r",errors.Cause(err))
	fmt.Printf("%v\n\r",errors.BaseCode(err))
	fmt.Printf("%v\n\r",errors.BaseHttpCode(err))

	// output : 网络调用错误
	// output : connection reset by peer
	// output : 10086
	// output : 500
}




// 错误统一处理层面的解决方案
func Example_handlerErr(){
	err:= getError()
	// 拿到一个错误

	// 判断错误是否为创建的错误 或者第三方错误
	if errors.ErrFundamental == errors.ErrorType(err) || errors.ErrWithCode == errors.ErrorType(err){
		// 获取到code
		fmt.Printf("%v \r\n", errors.Code(err))
		// 记录错误堆栈
		fmt.Printf("%v \r\n", errors.Code(err))
		// return res
	}
}