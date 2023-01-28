package utils

import (
	"encoding/json"
	"log"
)

func Serialize[T interface{}](v T) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Println("[Marshal Error]", err)
		return ""
	}
	return string(marshal)
}

func SerializeBytes[T interface{}](v T) []byte {
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Println("[Marshal Error]", err)
		return make([]byte, 0)
	}
	return marshal
}

func Deserialize[T interface{}](seq string) *T {
	if seq == "" {
		return nil
	}
	var res = new(T)
	err := json.Unmarshal([]byte(seq), res)
	if err != nil {
		log.Println("[Unmarshal Error]", err)
		return nil
	}
	return res
}

func DeserializeBytes[T interface{}](seq []byte) *T {
	if seq == nil || len(seq) == 0 {
		return nil
	}
	var res = new(T)
	err := json.Unmarshal(seq, res)
	if err != nil {
		log.Println("[Unmarshal Error]", err)
		return nil
	}
	return res
}
