package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 创建错误，以及创建错误的类型判断
func TestNewError(t *testing.T) {
	tests := []struct {
		code     int
		give     error
		expectTy string //期望的错误类型
	}{
		{code: ErrDatabase, give: NewCode(ErrDatabase) /* 得到一个被创建的error */, expectTy: ErrFundamental},
		{code: ErrDatabase, give: NewCode(ErrDatabase) /* 得到一个被创建的error */, expectTy: ErrFundamental},
		{
			code:     ErrDatabase,
			give:     WithCode(NewCode(ErrDatabase), ErrDecodingJSON) /* 得到一个被创建的error */,
			expectTy: ErrWithCode,
		},
	}

	for _, v := range tests {
		// 查看创建出的内容
		assert.Equal(t, v.expectTy, ErrorType(v.give))
	}
}

// withCode 将第三方错误保存下来，并用内部注册的code 将错误进行转为内部错误
// 在错误处理层面会将错误
func TestWithCode(t *testing.T) {
	tests := []struct {
		give       error  // 包装过后的error
		expectCode int    // 给的code
		expectTy   string // withCode 是期望的错误
		cause      error  // 调用cause 方法会显示包裹的信息
	}{
		{give: WithCode(NewCode(ErrBind), ErrDatabase), expectCode: ErrDatabase, expectTy: ErrWithCode, cause: NewCode(ErrBind)},
		{give: WithCode(NewCode(ErrBind), ErrDatabase), expectCode: ErrDatabase, expectTy: ErrWithCode, cause: NewCode(ErrBind)},
	}

	for _, v := range tests {
		assert.Equal(t, v.expectTy, ErrorType(v.give))          // 错误类型
		assert.Equal(t, v.cause.Error(), Cause(v.give).Error()) // 包裹的错误内容
		assert.Equal(t, v.expectCode, Code(v.give))
	}
}

// 测试循环包裹code ，循环包裹code 应该以
func TestLoopWithCode(t *testing.T) {

	var TestOrigin =fmt.Errorf("user is not good ",)

	tests := []struct {
		giveOrigin  error
		giveCodeOne int
		giveCodeTwo int
		expectCode int //should like giveCodeOne
		expectErr  error
	}{
		{
			giveOrigin:  TestOrigin,
			giveCodeOne: ErrDatabase,
			giveCodeTwo: ErrBind,
			expectCode:  ErrDatabase,
			expectErr:   TestOrigin,
		},
		{
			giveOrigin:  TestOrigin,
			giveCodeOne: ErrDatabase,
			giveCodeTwo: ErrBind,
			expectCode:  ErrDatabase,
			expectErr:   TestOrigin,
		},
	}

	// 循环创建一部分的code，测试错误Code的时候应该是和最初始的错误码一样
	for _,v := range tests {
		errOne := WithCode(v.giveOrigin,v.giveCodeOne)
		errTwo := WithCode(errOne,v.giveCodeTwo)
		assert.Equal(t,v.expectCode, Code(errTwo)) //获取到第一个withcode 的code
		assert.Equal(t,v.expectErr, Cause(errTwo)) // 对比原始的错误
	}
}
