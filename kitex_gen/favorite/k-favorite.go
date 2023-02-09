// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favorite

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
	"tiktok-server/kitex_gen/feed"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
	_ = feed.KitexUnusedProtection
)

func (p *FavoriteActionRequest) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetToken bool = false
	var issetVideoId bool = false
	var issetActionType bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetToken = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetVideoId = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetActionType = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetToken {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetVideoId {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetActionType {
		fieldId = 3
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteActionRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_FavoriteActionRequest[fieldId]))
}

func (p *FavoriteActionRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Token = v

	}
	return offset, nil
}

func (p *FavoriteActionRequest) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.VideoId = v

	}
	return offset, nil
}

func (p *FavoriteActionRequest) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.ActionType = v

	}
	return offset, nil
}

// for compatibility
func (p *FavoriteActionRequest) FastWrite(buf []byte) int {
	return 0
}

func (p *FavoriteActionRequest) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "FavoriteActionRequest")
	if p != nil {
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *FavoriteActionRequest) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("FavoriteActionRequest")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *FavoriteActionRequest) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "token", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Token)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteActionRequest) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "video_id", thrift.I64, 2)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.VideoId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteActionRequest) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "action_type", thrift.I32, 3)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.ActionType)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteActionRequest) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("token", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.Token)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteActionRequest) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("video_id", thrift.I64, 2)
	l += bthrift.Binary.I64Length(p.VideoId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteActionRequest) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("action_type", thrift.I32, 3)
	l += bthrift.Binary.I32Length(p.ActionType)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteActionResponse) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetStatusCode bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetStatusCode = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetStatusCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteActionResponse[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_FavoriteActionResponse[fieldId]))
}

func (p *FavoriteActionResponse) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.StatusCode = v

	}
	return offset, nil
}

func (p *FavoriteActionResponse) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.StatusMsg = &v

	}
	return offset, nil
}

// for compatibility
func (p *FavoriteActionResponse) FastWrite(buf []byte) int {
	return 0
}

func (p *FavoriteActionResponse) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "FavoriteActionResponse")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *FavoriteActionResponse) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("FavoriteActionResponse")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *FavoriteActionResponse) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status_code", thrift.I32, 1)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.StatusCode)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteActionResponse) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetStatusMsg() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status_msg", thrift.STRING, 2)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.StatusMsg)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *FavoriteActionResponse) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("status_code", thrift.I32, 1)
	l += bthrift.Binary.I32Length(p.StatusCode)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteActionResponse) field2Length() int {
	l := 0
	if p.IsSetStatusMsg() {
		l += bthrift.Binary.FieldBeginLength("status_msg", thrift.STRING, 2)
		l += bthrift.Binary.StringLengthNocopy(*p.StatusMsg)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *FavoriteListRequest) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetUserId bool = false
	var issetToken bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetUserId = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetToken = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetUserId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetToken {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteListRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_FavoriteListRequest[fieldId]))
}

func (p *FavoriteListRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.UserId = v

	}
	return offset, nil
}

func (p *FavoriteListRequest) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Token = v

	}
	return offset, nil
}

// for compatibility
func (p *FavoriteListRequest) FastWrite(buf []byte) int {
	return 0
}

func (p *FavoriteListRequest) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "FavoriteListRequest")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *FavoriteListRequest) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("FavoriteListRequest")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *FavoriteListRequest) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "user_id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.UserId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteListRequest) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "token", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Token)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteListRequest) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("user_id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.UserId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteListRequest) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("token", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.Token)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteListResponse) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetStatusCode bool = false
	var issetVideoList bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetStatusCode = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetVideoList = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetStatusCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetVideoList {
		fieldId = 3
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteListResponse[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_FavoriteListResponse[fieldId]))
}

func (p *FavoriteListResponse) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.StatusCode = v

	}
	return offset, nil
}

func (p *FavoriteListResponse) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.StatusMsg = &v

	}
	return offset, nil
}

func (p *FavoriteListResponse) FastReadField3(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.VideoList = make([]*feed.Video, 0, size)
	for i := 0; i < size; i++ {
		_elem := feed.NewVideo()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.VideoList = append(p.VideoList, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

// for compatibility
func (p *FavoriteListResponse) FastWrite(buf []byte) int {
	return 0
}

func (p *FavoriteListResponse) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "FavoriteListResponse")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *FavoriteListResponse) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("FavoriteListResponse")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *FavoriteListResponse) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status_code", thrift.I32, 1)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.StatusCode)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteListResponse) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetStatusMsg() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status_msg", thrift.STRING, 2)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.StatusMsg)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *FavoriteListResponse) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "video_list", thrift.LIST, 3)
	listBeginOffset := offset
	offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
	var length int
	for _, v := range p.VideoList {
		length++
		offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
	}
	bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
	offset += bthrift.Binary.WriteListEnd(buf[offset:])
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteListResponse) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("status_code", thrift.I32, 1)
	l += bthrift.Binary.I32Length(p.StatusCode)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteListResponse) field2Length() int {
	l := 0
	if p.IsSetStatusMsg() {
		l += bthrift.Binary.FieldBeginLength("status_msg", thrift.STRING, 2)
		l += bthrift.Binary.StringLengthNocopy(*p.StatusMsg)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *FavoriteListResponse) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("video_list", thrift.LIST, 3)
	l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.VideoList))
	for _, v := range p.VideoList {
		l += v.BLength()
	}
	l += bthrift.Binary.ListEndLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteServiceActionFavoriteArgs) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetReq bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetReq = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetReq {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceActionFavoriteArgs[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_FavoriteServiceActionFavoriteArgs[fieldId]))
}

func (p *FavoriteServiceActionFavoriteArgs) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewFavoriteActionRequest()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Req = tmp
	return offset, nil
}

// for compatibility
func (p *FavoriteServiceActionFavoriteArgs) FastWrite(buf []byte) int {
	return 0
}

func (p *FavoriteServiceActionFavoriteArgs) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ActionFavorite_args")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *FavoriteServiceActionFavoriteArgs) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ActionFavorite_args")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *FavoriteServiceActionFavoriteArgs) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "req", thrift.STRUCT, 1)
	offset += p.Req.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteServiceActionFavoriteArgs) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("req", thrift.STRUCT, 1)
	l += p.Req.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteServiceActionFavoriteResult) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField0(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceActionFavoriteResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteServiceActionFavoriteResult) FastReadField0(buf []byte) (int, error) {
	offset := 0

	tmp := NewFavoriteActionResponse()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Success = tmp
	return offset, nil
}

// for compatibility
func (p *FavoriteServiceActionFavoriteResult) FastWrite(buf []byte) int {
	return 0
}

func (p *FavoriteServiceActionFavoriteResult) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ActionFavorite_result")
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *FavoriteServiceActionFavoriteResult) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ActionFavorite_result")
	if p != nil {
		l += p.field0Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *FavoriteServiceActionFavoriteResult) fastWriteField0(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "success", thrift.STRUCT, 0)
		offset += p.Success.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *FavoriteServiceActionFavoriteResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += bthrift.Binary.FieldBeginLength("success", thrift.STRUCT, 0)
		l += p.Success.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *FavoriteServiceListFavoriteArgs) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetReq bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetReq = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetReq {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceListFavoriteArgs[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_FavoriteServiceListFavoriteArgs[fieldId]))
}

func (p *FavoriteServiceListFavoriteArgs) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewFavoriteListRequest()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Req = tmp
	return offset, nil
}

// for compatibility
func (p *FavoriteServiceListFavoriteArgs) FastWrite(buf []byte) int {
	return 0
}

func (p *FavoriteServiceListFavoriteArgs) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ListFavorite_args")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *FavoriteServiceListFavoriteArgs) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ListFavorite_args")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *FavoriteServiceListFavoriteArgs) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "req", thrift.STRUCT, 1)
	offset += p.Req.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *FavoriteServiceListFavoriteArgs) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("req", thrift.STRUCT, 1)
	l += p.Req.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *FavoriteServiceListFavoriteResult) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField0(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceListFavoriteResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteServiceListFavoriteResult) FastReadField0(buf []byte) (int, error) {
	offset := 0

	tmp := NewFavoriteListResponse()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Success = tmp
	return offset, nil
}

// for compatibility
func (p *FavoriteServiceListFavoriteResult) FastWrite(buf []byte) int {
	return 0
}

func (p *FavoriteServiceListFavoriteResult) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ListFavorite_result")
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *FavoriteServiceListFavoriteResult) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ListFavorite_result")
	if p != nil {
		l += p.field0Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *FavoriteServiceListFavoriteResult) fastWriteField0(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "success", thrift.STRUCT, 0)
		offset += p.Success.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *FavoriteServiceListFavoriteResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += bthrift.Binary.FieldBeginLength("success", thrift.STRUCT, 0)
		l += p.Success.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *FavoriteServiceActionFavoriteArgs) GetFirstArgument() interface{} {
	return p.Req
}

func (p *FavoriteServiceActionFavoriteResult) GetResult() interface{} {
	return p.Success
}

func (p *FavoriteServiceListFavoriteArgs) GetFirstArgument() interface{} {
	return p.Req
}

func (p *FavoriteServiceListFavoriteResult) GetResult() interface{} {
	return p.Success
}
