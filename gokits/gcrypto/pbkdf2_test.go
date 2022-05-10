package gcrypto

import "testing"

func TestPbkdf2(t *testing.T) {
	password := "ywlSRb80TaCQ4b7b"
	encrypted, salt := GenPbkdf2Passwd(password, Pbkdf2SaltLen, Pbkdf2IterTimes, Pbkdf2keyLen)
	t.Log(encrypted, salt)
	if !CmpPbkdf2Passwd(password, salt, encrypted, Pbkdf2IterTimes, Pbkdf2keyLen) {
		t.Fail()
	}
}
