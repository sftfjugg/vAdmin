package tools

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

func StrToInt(err error, index string) int {
	result, err := strconv.Atoi(index)
	if err != nil {
		HasError(err, "string to int error"+err.Error(), -1)
	}
	return result
}

/*
bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
           Must be the already hashed PW ^              ^ Plain Text Password to compare
*/
func CompareHashAndPassword(e string, p string) (bool, error) {

	byteHash := []byte(e)  // e: hashedPassword
	plainPwd := []byte(p)
	log.Println("vvvvvvvvvvvvvvvvvvvvvvv--->hashedPassword", e)
	log.Println("vvvvvvvvvvvvvvvvvvvvvvv--->PlainPassword", p)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}

func CompareHashAndPassword2(e string, p string) (bool, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(e), bcrypt.DefaultCost) //加密处理
	if err != nil {
		log.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可

	byteHash := []byte(encodePWD)  // e: hashedPassword
	plainPwd := []byte(p)
	log.Println("vvvvvvvvvvvvvvvvvvvvvvv--->hashedPassword", encodePWD)
	log.Println("vvvvvvvvvvvvvvvvvvvvvvv--->PlainPassword", p)
	err = bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}

/*
func CompareHashAndPassword(e string, p []byte) (bool, error) {

	byteHash := []byte(e)
	err := bcrypt.CompareHashAndPassword(byteHash, p)
	if err != nil {
		log.Print(err.Error())
		return false, err
	}
	return true, nil
}
*/

// Assert 条件断言
// 当断言条件为 假 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
func Assert(condition bool, msg string, code ...int) {
	if !condition {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func HasError(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		if msg == "" {
			msg = err.Error()
		}
		log.Println(err)
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}