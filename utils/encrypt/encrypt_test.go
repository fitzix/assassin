package encrypt

import "testing"

func TestPassEncrypt(t *testing.T) {
	pwd := PassEncrypt("Ng7KQ!dcasj7yVJGJXrmH")
	t.Log(pwd)
}