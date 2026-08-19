package base64

import (
	"testing"
)

func TestBase64(t *testing.T) {
	input := "this is a test"
	expectedOutput := "dGhpcyBpcyBhIHRlc3Q="

	base64output, err := EncodeBase64(input)

	if err != nil {
		t.Fatalf("EncodeBase64 returned an error %e", err)
	}

	if expectedOutput != base64output {
		t.Errorf("Expect output string of \"%v\" does not match actual output value \"%v\"", expectedOutput, base64output)
	}
}
