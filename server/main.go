package main

import (
	"fmt"
)

// CustomHTTPError 型は error インターフェースを満たす
type CustomHTTPError struct {
	Code       int
	Message    string
	ErrorCode  string
	stackTrace []uintptr
}

// Error メソッドを実装することで error インターフェースを満たす
func (e *CustomHTTPError) Error() string {
	return fmt.Sprintf("Error: %s (code: %d)", e.Message, e.Code)
}

func HandleError(err error) {
	// 引数は error 型なので、CustomHTTPError も渡せる
	fmt.Println(err.Error()) // CustomHTTPError の Error() が呼ばれる
}

func main() {
	// CustomHTTPError を作成
	err := &CustomHTTPError{
		Code:       404,
		Message:    "Page not found",
		ErrorCode:  "NOT_FOUND",
		stackTrace: []uintptr{0x001, 0x002, 0x003}, // スタックトレースの例
	}

	// HandleError に CustomHTTPError を渡す
	HandleError(err) // error インターフェースとして渡される
}
