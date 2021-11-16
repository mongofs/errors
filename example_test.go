//
package errors_test

import (
	"fmt"
	errors_map "maoti_sdk/logic/pkg/errors-map"
)

// 创建具体错误，
func Example_handlerCreatedByEngineersErr() {
	// 关于error_code test
	errors_map.Register(100910,500,"user is not found ")
	defer errors_map.Flush()
	ErrUserNotFound := errors_map.NewCode(100910)

	// 模拟具体程序产生错误如何handler
	if errors_map.ErrWithCode == errors_map.ErrorType(ErrUserNotFound){
		// check the error cause
		fmt.Printf("error cause is  : %s \n" ,errors_map.Cause(ErrUserNotFound))
		// check this code
		fmt.Printf("error code is : %v \n" ,errors_map.Code(ErrUserNotFound))
	}

	// 这里是建议处理最终错误的模式，如果在api接口调用过程中抛出的错误统一处理的地方
	// 不论是api类型服务还是脚本类型服务，应该在项目统一错误处理的地方进行错误处理，
	// 错误处理原则 1. 不应该太过分散的处理 2.不应该出现未知错误，一旦出现未知错误表明程序
	// 已经进入非常糟糕的阶段
	// 对内部错误记录应该记录：调用堆栈，code ，httpCode，ext，
	// 出现错误对外部调用应该返回：code，httpCode,ext
}



func getError()error{
	return errors_map.New("connection reset by peer ")
}


// 处理产生的第三方错误
func Example_handlerCreatedByThirdPartyErr() {
	var code = 10086 // 网络调用错误
	errors_map.Register(code,500,"网络调用错误")
	defer errors_map.Flush()
	err := getError();
	// 在你这边应该将第三方的错误进行归类
	withCode := errors_map.WithCode(err,code)
	// 通过返回withcode 创建一个新的错误 ，类似wrapper
	// 直接将withCode 返回出去
	fmt.Printf("%v\n\r",withCode)
}




// 错误统一处理层面的解决方案
func Example_handlerErr(){
	err:= getError()
	// 拿到一个错误

	// 判断错误是否为创建的错误 或者第三方错误
	if errors_map.ErrFundamental == errors_map.ErrorType(err) || errors_map.ErrWithCode == errors_map.ErrorType(err){
		// 获取到code
		fmt.Printf("%v \r\n", errors_map.Code(err))
		// 记录错误堆栈
		fmt.Printf("%v \r\n", errors_map.Code(err))
		// return res
	}
	// 当出现程序错误走到这里，代表存在错误是未知的
	// 这里错误就变得十分严重，原则上所有的错误都应该在错误码中包含，至少描述必须包含

	// 这里是建议处理最终错误的模式，如果在api接口调用过程中抛出的错误统一处理的地方
	// 不论是api类型服务还是脚本类型服务，应该在项目统一错误处理的地方进行错误处理，
	// 错误处理原则 1. 不应该太过分散的处理 2.不应该出现未知错误，一旦出现未知错误表明程序
	// 已经进入非常糟糕的阶段
	// 对内部错误记录应该记录：调用堆栈，code ，httpCode，ext，
	// 出现错误对外部调用应该返回：code，httpCode,ext
}