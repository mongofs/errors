package errors

// 整体errors包管理理念是：用户注册错误码，
// 这个项目主要用户管理项目错误：内部整合了：github/pkg/errors-map ,底层使用了这个包内的错误，
// 整体用法是用户 registerError  此方法，将错误注册到当前包内，由包维护一个内部错误map，
// 当项目产生错误的时候，只需要调用方法 if ok :=errors_map.IsFundamental ;!ok { handler(otherErr)}
// else { fmt.Println( errors_map.GetCode() )}
