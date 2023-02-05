// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package erren

import (
	"errors"
	"fmt"
)

// 使用map类型的key来检查错误是否在错误代码组里
// example: if _,ok := erren.ErrorMap[10001]; ok{}
// 返回ok=true，10001在错误代码组，false则不在

var ErrorMap = map[int32]struct{}{10001: {}, 10002: {}, 10003: {}, 10004: {}, 10005: {}}

const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	AuthorizationFailedErrCode = 10003
	UserAlreadyExistErrCode    = 10004
	UserNotExistErrCode        = 10005
)

type ErrEn struct {
	ErrCode int32
	ErrMsg  string
}

/*
func (e ErrEn) ErrorMap() map[int]int {
	var ErrorMap map[int]int
	m := []int{10000,10001, 2}
	for _, v := range m {
		ErrorMap[v] = v
	}

	return ErrorMap
}*/

func (e ErrEn) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrEn {
	return ErrEn{code, msg}
}

func (e ErrEn) WithMessage(msg string) ErrEn {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, "Successfully")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User Already Exist")
	UserNotExistErr        = NewErrNo(UserNotExistErrCode, "User Not Exist")
)

// ConvertErr convert error to ErrEn
func ConvertErr(err error) ErrEn {
	Err := ErrEn{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
