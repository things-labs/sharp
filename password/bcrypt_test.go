package password

import (
	"testing"
)

func TestBCrypt(t *testing.T) {
	org := "hahaha"
	cpt := NewBCrypt("")
	dst, err := cpt.Hash(org)
	if err != nil {
		t.Fatalf("hash, %v", err)
	}

	if err := cpt.Compare(org, dst); err != nil {
		t.Fatalf("Compare, %v", err)
	}
}

func BenchmarkBCrypt_Hash(b *testing.B) {
	cpt := NewBCrypt("111")

	for i := 0; i < b.N; i++ {
		_, _ = cpt.Hash("hahaha")
	}
}

func BenchmarkBCrypt_Compare(b *testing.B) {
	org := "hahaha"
	cpt := NewBCrypt("")
	dst, _ := cpt.Hash(org)

	for i := 0; i < b.N; i++ {
		_ = cpt.Compare(org, dst)
	}
}
