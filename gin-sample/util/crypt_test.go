package util

import "testing"

func TestEnCrypt(t *testing.T) {
	tests := []string{"12345667234", "xwerfqwew@@#4d"}
	for _, c := range tests {
		actual := EnCrypt(c)
		if ok := ValidatePassword(c, actual); !ok {
			t.Fatal("encrypt password failed: ", c)
		}
	}
}
