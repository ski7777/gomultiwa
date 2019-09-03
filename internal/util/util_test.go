package util

import "testing"

func TestWidToJid(t *testing.T) {
	type pair struct {
		i string
		o string
	}
	testdata := []pair{
		pair{"a@b.com", "a@s.whatsapp.net"},
		pair{"1234", "1234@s.whatsapp.net"},
		pair{"abc@a@b", "abc@s.whatsapp.net"},
	}
	for _, p := range testdata {
		if r := WidToJid(p.i); r != p.o {
			t.Errorf("Conversion from %s to %s was incorrect. Got %s to %s!", p.i, p.o, p.i, r)
		}
	}
}
