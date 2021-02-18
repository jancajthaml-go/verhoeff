package verhoeff

import "testing"

func TestVerhoeff(t *testing.T) {
	expectDigit := func(a string, b int) {
		c, err := Digit(a)
		if err != nil {
			t.Errorf("%v", err)
		} else if b != c {
			t.Errorf("for input " + a + " expected digit " + string(rune(b+48)) + " but got " + string(rune(c+48)) + " instead")
		}
	}

	expectSignature := func(a string, b string) {
		c, err := Generate(a)
		if err != nil {
			t.Errorf("%v", err)
		} else if b != c {
			t.Errorf("for input " + a + " expected signature " + b + " but got " + c + " instead")
		}
		if !Validate(c) {
			t.Errorf("Unable to validate signature that was generated")
		}
	}

	expectDigit("75872", 2)
	expectDigit("12345", 1)
	expectDigit("142857", 0)
	expectDigit("123456789012", 0)
	expectDigit("8473643095483728456789", 2)

	if !Validate("84736430954837284567892") {
		t.Errorf("Checksum failed for valid input")
	}

	if Validate("12") {
		t.Errorf("Checksum passed for invalid input")
	}

	if Validate("xy-1") {
		t.Errorf("Checksum passed for invalid alphabet")
	}

	expectSignature("8473643095483728456789", "84736430954837284567892")
}


func BenchmarkVerhoeffSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Digit("123")
	}
}

func BenchmarkVerhoeffLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Digit("00123014764700968325")
	}
}

func BenchmarkVerhoeffSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digit("123")
		}
	})
}

func BenchmarkVerhoeffLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Digit("00123014764700968325")
		}
	})
}
