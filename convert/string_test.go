package convert

import "testing"

func TestString(t *testing.T) {

	tString := "goai is a test"

	md5 := StrTo(tString).MD5()
	t.Logf("string to MD5: %s", md5)

	mult := StrTo(tString).MultiWord()
	t.Logf("string to MultiWord: %+#v", mult)

	numberEn := NumberEncode("575757", []byte("123"))
	t.Logf("string to NumberEncode: %+#v", numberEn)

	numberDe := NumberDecode(numberEn, []byte("123"))
	t.Logf("string to NumberDecode: %+#v", numberDe)
}
