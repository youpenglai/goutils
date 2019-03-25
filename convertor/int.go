package convertor

import (
	"encoding/binary"
	"strconv"
	"strings"
)

func ConvertInt(n interface{}) int {
	var result int
	switch n.(type) {
	case int:
		result = n.(int)
	case int64:
		v, _ := n.(int64)
		result = int(v)
	case float64:
		v, _ := n.(float64)
		result = int(v)
	case string:
		result, _ = strconv.Atoi(n.(string))
	}
	return result
}

func ConvertInt32(n interface{}) int32 {
	var result int32
	switch n.(type) {
	case int64:
		result = n.(int32)
	case float64:
		v, _ := n.(float64)
		result = int32(v)
	case int:
		i, _ := n.(int)
		result = int32(i)
	case string:
		resultInt64, _ := strconv.ParseInt(n.(string), 10, 32)
		result = ConvertInt32(resultInt64)
	default:
	}
	return result
}

func ConvertInt64(n interface{}) int64 {
	var result int64
	switch n.(type) {
	case int64:
		result = n.(int64)
	case float64:
		v, _ := n.(float64)
		result = int64(v)
	case int:
		i, _ := n.(int)
		result = int64(i)
	case string:
		result, _ = strconv.ParseInt(n.(string), 10, 64)
	default:
	}
	return result
}

func ConvertUint64(n interface{}) uint64 {
	var result uint64
	switch n.(type) {
	case uint64:
		result = n.(uint64)
	case int64:
		i, _ := n.(int64)
		result = uint64(i)
	case float64:
		v, _ := n.(float64)
		result = uint64(v)
	case int:
		i, _ := n.(int)
		result = uint64(i)
	case string:
		result, _ = strconv.ParseUint(n.(string), 10, 64)
	default:
	}
	return result
}

func ConvertInt64List(s string, sep string) ([]int64, error) {
	strList := strings.Split(s, sep)
	int64List := make([]int64, len(strList))
	for index, str := range strList {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		int64List[index] = i
	}

	return int64List, nil
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
