package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strings"

	"github.com/matoous/go-nanoid"
	"golang.org/x/crypto/scrypt"
)

func GetNanoId() string {
	id, _ := gonanoid.Nanoid(16)
	return id
}

func PassEncrypt(password, salt string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 16)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", dk), nil
}

/*CBC加密 按照golang标准库的例子代码
不过里面没有填充的部分,所以补上
*/

// 使用PKCS7进行填充，IOS也是7
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// func PKCS7UnPadding(origData []byte) []byte {
// 	length := len(origData)
// 	unPadding := int(origData[length-1])
// 	return origData[:(length - unPadding)]
// }

// aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func aesCBCEncrypt(rawData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 填充原文
	blockSize := block.BlockSize()
	rawData = PKCS7Padding(rawData, blockSize)
	// 初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize)
	// block大小 16

	// block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, rawData)

	return cipherText, nil
}

func AesEncrypt(rawData, key, iv string) (string, error) {
	data, err := aesCBCEncrypt([]byte(rawData), []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}
	return strings.ToUpper(fmt.Sprintf("%x", data)), nil
}
