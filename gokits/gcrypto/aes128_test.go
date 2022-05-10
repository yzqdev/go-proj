package gcrypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCryptoWithRootKey(t *testing.T) {
	crypto, _ := NewCrypto("ywlSRb80TaCQ4b7b")
	usr, _ := crypto.EncryptStr("goman")
	t.Logf("%s", usr)
	d, _ := crypto.DecryptStr(usr)
	assert.Equal(t, "goman", d)

	pwd, _ := crypto.EncryptStr("123456")
	t.Logf("%s", pwd)
	d, _ = crypto.DecryptStr(pwd)
	assert.Equal(t, "123456", d)
}

func TestCryptoWithWorkKey(t *testing.T) {
	crypto, _ := NewCrypto("ywlSRb80TaCQ4b7b")
	key, usr, _ := crypto.EncryptStrWithRandKey("goman")
	t.Logf("%s", key)
	t.Logf("%s", usr)
	d, _ := crypto.DecryptStrWithKey(key, usr)
	assert.Equal(t, "goman", d)

	key2, pwd, _ := crypto.EncryptStrWithRandKey("123456")
	t.Logf("%s", key2)
	t.Logf("%s", pwd)
	d, _ = crypto.DecryptStrWithKey(key2, pwd)
	assert.Equal(t, "123456", d)
}
