package convert

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//StrTo convert string to specify type
type StrTo string

//Set set string value
func (f *StrTo) Set(v string) {
	if v != "" {
		*f = StrTo(v)
	} else {
		f.Clear()
	}
}

//Clear clear StrTo
func (f *StrTo) Clear() {
	*f = StrTo(0x1E)
}

//Exist check StrTo exist
func (f StrTo) Exist() bool {
	return string(f) != string(0x1E)
}

//Bool value is true or false
func (f StrTo) Bool() (bool, error) {
	if f == "on" {
		return true, nil
	}
	return strconv.ParseBool(f.String())
}

//Float32 string to float32
func (f StrTo) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

//Float64 string to float64
func (f StrTo) Float64() (float64, error) {
	return strconv.ParseFloat(f.String(), 64)
}

//Int string to int
func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int(v), err
}

//Int8 string to int8
func (f StrTo) Int8() (int8, error) {
	v, err := strconv.ParseInt(f.String(), 10, 8)
	return int8(v), err
} //Float32 string 2 float32

//Int16 string to int16
func (f StrTo) Int16() (int16, error) {
	v, err := strconv.ParseInt(f.String(), 10, 16)
	return int16(v), err
}

//Int32 string to int32
func (f StrTo) Int32() (int32, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int32(v), err
}

//Int64 string to int64
func (f StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

//Uint string to uint
func (f StrTo) Uint() (uint, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint(v), err
}

//Uint8 string to uint8
func (f StrTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

//Uint16 string to uint16
func (f StrTo) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(f.String(), 10, 16)
	return uint16(v), err
}

//Uint32 string to uint32
func (f StrTo) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint32(v), err
}

//Uint64 string to uint64
func (f StrTo) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	return uint64(v), err
}

func (f StrTo) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

//MD5 string to md5
func (f StrTo) MD5() string {
	h := md5.New()
	h.Write([]byte(f.String()))
	return hex.EncodeToString(h.Sum(nil))
}

//MultiWord string to []string
func (f StrTo) MultiWord() []string {
	str := f.String()

	str = strings.Replace(str, "\u3000", " ", -1)
	str = strings.Replace(str, "、", " ", -1)
	str = strings.Replace(str, ",", " ", -1)
	str = strings.Replace(str, "【", " ", -1)
	str = strings.Replace(str, "】", " ", -1)

	re, _ := regexp.Compile(`\s{2,}`)
	str = re.ReplaceAllString(str, " ")

	str = strings.Trim(str, " ")

	return strings.Split(str, " ")
}

//ToStr convert any type to string
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

//ToInt64 convert any numeric value to int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

type argString []string

func (a argString) Get(i int, args ...string) (r string) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

type argAny []interface{}

func (a argAny) Get(i int, args ...interface{}) (r interface{}) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

//GetRandomString Random generate string len is n
func GetRandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

//NumberEncode number string encode by alphabet
func NumberEncode(number string, alphabet []byte) string {
	token := make([]byte, 0, 12)
	x, ok := new(big.Int).SetString(number, 10)
	if !ok {
		return ""
	}
	y := big.NewInt(int64(len(alphabet)))
	m := new(big.Int)
	for x.Sign() > 0 {
		x, m = x.DivMod(x, y, m)
		token = append(token, alphabet[int(m.Int64())])
	}
	return string(token)
}

//NumberDecode number string decode by alphabet
func NumberDecode(token string, alphabet []byte) string {
	x := new(big.Int)
	y := big.NewInt(int64(len(alphabet)))
	z := new(big.Int)
	for i := len(token) - 1; i >= 0; i-- {
		v := bytes.IndexByte(alphabet, token[i])
		z.SetInt64(int64(v))
		x.Mul(x, y)
		x.Add(x, z)
	}
	return x.String()
}
