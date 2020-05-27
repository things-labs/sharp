package password

import (
	"testing"
)

func TestSimple(t *testing.T) {
	org := "hahaha"
	cpt := NewSimple("")
	dst, err := cpt.Hash(org)
	if err != nil {
		t.Fatalf("hash, %v", err)
	}

	if err := cpt.Compare(org, dst); err != nil {
		t.Fatalf("Compare, %v", err)
	}
}

func BenchmarkSimple_Hash(b *testing.B) {
	cpt := NewSimple("")

	for i := 0; i < b.N; i++ {
		_, _ = cpt.Hash("hahaha")
	}
}

func BenchmarkSimple_Compare(b *testing.B) {
	org := "hahaha"
	cpt := NewSimple("")
	dst, _ := cpt.Hash(org)

	for i := 0; i < b.N; i++ {
		_ = cpt.Compare(org, dst)
	}
}
