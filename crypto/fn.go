package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Md5(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}

func Sha1(data []byte) []byte {
	hash := sha1.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64DecodeNoPadding(data string) ([]byte, error) {
	return base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(data)
}

// SortData 对数据递归排序
func SortData(data interface{}) interface{} {
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Pointer:
		return SortData(val.Elem().Interface())
	case reflect.Struct:
		return sortStructFields(val)
	case reflect.Map:
		return sortMap(val)
	case reflect.Slice:
		return sortSlice(val)
	default:

		return fmt.Sprintf("%+v", data)
	}

}

// 对结构体字段按首字母排序
func sortStructFields(val reflect.Value) string {
	typ := val.Type()

	// 获取并排序字段名称
	fieldNames := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		fieldNames[i] = typ.Field(i).Name
	}
	sort.Strings(fieldNames)

	s := make([]string, 0)

	for _, name := range fieldNames {
		field := val.FieldByName(name)
		s = append(s, fmt.Sprintf("%s=%+v", name, SortData(field.Interface())))
	}
	return strings.Join(s, "&")
}

// 对 map 的键按顺序排序
func sortMap(val reflect.Value) string {
	keys := val.MapKeys()
	// 将键排序
	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j])
	})

	s := make([]string, 0)
	// 按排序后的键填充有序 map
	for _, key := range keys {
		s = append(s, fmt.Sprintf("%+v=%+v", key.Interface(), SortData(val.MapIndex(key).Interface())))
	}
	return strings.Join(s, "&")
}

// 对 slice 的值进行排序
func sortSlice(val reflect.Value) string {
	length := val.Len()
	s := make([]string, 0)

	// 复制并递归排序 slice 内的值
	for i := 0; i < length; i++ {
		s = append(s, fmt.Sprintf("%d=%+v", i, SortData(val.Index(i).Interface())))
	}

	return strings.Join(s, "&")
}
