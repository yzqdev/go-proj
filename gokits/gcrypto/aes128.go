package gcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

const (
	keyRootIterCount = 5000
	emptyStr = ""
)

var rootKeySalt = []byte{
	0x45, 0xEF, 0x2F, 0x62,
	0xAB, 0x9C, 0xE8, 0x03,
	0x28, 0xB1, 0xF2, 0x61,
	0xDE, 0xF1, 0xD2, 0x58,
}

var (
	// ErrAESTextSize ...
	ErrAESTextSize = errors.New("cipher text length is not a multiple of the block size")
	// ErrAESPadding ...
	ErrAESPadding = errors.New("cipher padding size error")
	// ErrInvalidKey ...
	ErrInvalidKey = errors.New("invalid key, it's length should more than 16")
)

// Crypto the crypto object
type Crypto struct {
	rootBlock cipher.Block
	rootKey   []byte
}

// NewCrypto create a instance of *Crypto
func NewCrypto(factor string) (*Crypto, error) {
	c := new(Crypto)
	return c, c.init(factor)
}

func (c *Crypto) calKey(factor string) []byte {
	fbs := []byte(factor)
	mac := hmac.New(sha256.New, rootKeySalt)
	for i := 0; i < keyRootIterCount; i++ {
		mac.Reset()
		mac.Write(fbs)
		mac.Write([]byte{byte(i >> 24 & 0xFF), byte(i >> 16 & 0xFF), byte(i >> 8 & 0xFF), byte(i & 0xFF)})
		fbs = mac.Sum(nil)
	}

	return c.normalKey(fbs)
}

func (c *Crypto) normalKey(key []byte) []byte {
	blen := aes.BlockSize
	if len(key) >= blen {
		return key[:blen]
	}
	panic("export key fatal failed.")
}

// init
func (c *Crypto) init(factor string) error {
	c.rootKey = c.calKey(factor)
	block, err := aes.NewCipher([]byte(c.rootKey))
	if err != nil {
		return err
	}
	c.rootBlock = block
	return nil
}

// Decrypt from an encrypted array of byte
func (c *Crypto) Decrypt(src []byte) ([]byte, error) {
	return c.decrypt(c.rootBlock, src)
}

func (c *Crypto) decrypt(block cipher.Block, src []byte) ([]byte, error) {
	blen := aes.BlockSize

	// check the length
	if len(src) < blen*2 || len(src)%blen != 0 {
		return nil, ErrAESTextSize
	}

	// IV
	iv := src[:blen]
	// encrypt(text)
	srcLen := len(src) - blen
	decryptText := make([]byte, srcLen)

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decryptText, src[blen:])

	// unpadding
	paddingLen := int(decryptText[srcLen-1])
	if paddingLen > 16 {
		return nil, ErrAESPadding
	}

	return decryptText[:srcLen-paddingLen], nil
}

// DecryptStr decrypt from an encrypted base64 string
func (c *Crypto) DecryptStr(encryptedText string) (string, error) {
	src, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return emptyStr, err
	}

	d, err := c.Decrypt(src)
	if err != nil {
		return emptyStr, err
	}
	return string(d), err
}

// Encrypt an array byte
func (c *Crypto) Encrypt(src []byte) ([]byte, error) {
	return c.encrypt(c.rootBlock, src)
}

func (c *Crypto) encrypt(block cipher.Block, src []byte) ([]byte, error) {
	blen := aes.BlockSize

	// padding
	padLen := blen - (len(src) % blen)
	for i := 0; i < padLen; i++ {
		src = append(src, byte(padLen))
	}

	// iv || encrypt(text)
	srcLen := len(src)
	encryptText := make([]byte, blen+srcLen)

	// iv
	iv := encryptText[:blen]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encryptText[blen:], src)

	return encryptText, nil

}

// EncryptStr encrypt a string
func (c *Crypto) EncryptStr(plainText string) (string, error) {
	src := []byte(plainText)
	encrypted, err := c.Encrypt(src)
	if err != nil {
		return emptyStr, err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// EncryptWithKey encrypt a string with a given work key
// encrypt work key using root key firstly
// then encrypt text using work key
func (c *Crypto) EncryptWithKey(key []byte, plainText []byte) (encryptedKey []byte, encryptedText []byte, err error) {
	if len(key) < aes.BlockSize {
		err = ErrInvalidKey
		return
	}
	encryptedKey, err = c.Encrypt(key)
	if err != nil {
		return
	}

	workKey := c.normalKey(key)
	block, err := aes.NewCipher([]byte(workKey))
	if err != nil {
		return
	}

	encryptedText, err = c.encrypt(block, []byte(plainText))
	if err != nil {
		return
	}
	return
}

// EncryptStrWithKey encrypt a byte array with a given work key
// encrypt work key using root key firstly
// then encrypt text using work key
func (c *Crypto) EncryptStrWithKey(key []byte, plainText string) (encryptedKey string, encryptedText string, err error) {
	encryptedKeyBs, encryptedTextBs, err1 := c.EncryptWithKey(key, []byte(plainText))
	if err1 != nil {
		err = err1
		return
	}

	encryptedKey = base64.StdEncoding.EncodeToString(encryptedKeyBs)
	encryptedText = base64.StdEncoding.EncodeToString(encryptedTextBs)
	return
}

// EncryptStrWithRandKey encrypt a string with a random work key
func (c *Crypto) EncryptStrWithRandKey(plainText string) (encryptedKey string, encryptedText string, err error) {
	key := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return emptyStr, emptyStr, err
	}
	return c.EncryptStrWithKey(key, plainText)
}

// DecryptStrWithKey decrypt a byte array with encrypted work key
func (c *Crypto) DecryptWithKey(encryptedKey []byte, encryptedText []byte) (plainText []byte, err error) {
	plainKey, err1 := c.Decrypt(encryptedKey)
	if err1 != nil {
		err = err1
		return
	}

	workKey := c.normalKey(plainKey)
	block, err1 := aes.NewCipher([]byte(workKey))
	if err1 != nil {
		err = err1
		return
	}

	plainText, err = c.decrypt(block, encryptedText)
	if err != nil {
		return
	}
	return
}

// DecryptStrWithKey decrypt a string with encrypted work key
func (c *Crypto) DecryptStrWithKey(encryptedKey string, encryptedText string) (plainText string, err error) {
	srcKey, err := base64.StdEncoding.DecodeString(encryptedKey)
	if err != nil {
		return emptyStr, err
	}

	srcText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return emptyStr, err
	}

	dText, err1 := c.DecryptWithKey(srcKey, srcText)
	if err1 != nil {
		err = err1
		return
	}

	return string(dText), err
}
