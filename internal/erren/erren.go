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

const (
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	AuthorizationFailedErrCode = 10003
)

type ErrEn struct {
	ErrCode int32
	ErrMsg  string
}

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
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
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
