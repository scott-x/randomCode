package randomCode

import (
	"errors"
	"math/rand"
	"time"
)

const data = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var ParamError  error = errors.New("param should be expect")

func Random(n int) (string,error){
	max := len(data)
	if n <= 0 || n> max{
		return "",ParamError
	}
	rand.Seed(time.Now().Unix())
	var result []byte
	for i:=0;i<n;i++ {
		result = append(result,data[rand.Intn(max)])
	}
	return string(result),nil
}