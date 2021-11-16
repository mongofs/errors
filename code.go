package errors

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

var (
	_rwMutex  = sync.RWMutex{}
	_errorMap = make(map[int]*inn, 30)
)

var (
	_httpCode = map[int]bool{
		200: true, //200
		400: true, //400
		401: true, //401
		402: true, //402
		403: true, //403
		404: true, //404
		500: true, //500
	}
)

var (
	// 在这里创建错误通过错误map创建错误，如果错误code没有在错误map中注册过，那么
	// 就会报错errCodeNotExist
	// 调用方应该首先判断erros.is( errCodeNotExist)
	ErrCodeNotExist = errors.New("errors-map:  code is not exist")
)

type inn struct {
	c    int    // code
	http int    // httpCode
	ext  string //extra
}

func Register(code int, httpCode int, ext string) {
	if _, ok := _errorMap[code]; ok {
		panic(fmt.Sprintf("code %d: is exist", code))
	}

	if _, ok := _httpCode[httpCode]; !ok {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}

	if len(ext) == 0 {
		panic("code have no decrption")
	}

	temV := &inn{
		c:    code,
		http: httpCode,
		ext:  ext,
	}
	_errorMap[code] = temV
}

// get information by errCode
func getExt(errCode int) (inn, error) {
	_rwMutex.RLock()
	defer _rwMutex.RUnlock()
	if e, ok := _errorMap[errCode]; !ok {
		panic(fmt.Sprintf("code %d: is not exist", errCode))
	} else {
		tem := *e
		return tem, nil
	}
}



func Flush(){
	_rwMutex.Lock()
	defer  _rwMutex.Unlock()
	_errorMap = make(map[int]*inn, 1)
}