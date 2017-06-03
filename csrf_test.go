package csrf

import (
	"fmt"
	"testing"
)

const secret = "erHUnxuhBMRIsVB1LfqmiWCgB83ZEerH"

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestToolset_GenerateSalt(t *testing.T) {
	tests := []struct {
		saltLen int
	}{
		{
			saltLen: 8,
		},
		{
			saltLen: 16,
		},
		{
			saltLen: 32,
		},
	}

	var generated []string
	for i, test := range tests {
		toolset := New(Options{
			Secret:  secret,
			SaltLen: test.saltLen,
		})
		errorPrefix := fmt.Sprintf("Test [%d]: ", i)
		actual := toolset.GenerateSalt()
		if contains(generated, actual) {
			t.Errorf(errorPrefix + "Generation isn't random!")
		}
		generated = append(generated, actual)
		if len(actual) != test.saltLen {
			t.Errorf(errorPrefix+"Expected string length of %d, got %d", test.saltLen, len(actual))
		}
	}
}

func TestToolset_GenerateToken(t *testing.T) {
	toolset := New(Options{
		Secret:  secret,
		SaltLen: 16,
	})
	tests := []struct {
		salt     string
		expected string
	}{
		{
			salt:     "uvqIUfqJ3W0qlszj",
			expected: "uvqIUfqJ3W0qlszjdJkEzNyB_hee0YLbJJoGIR_voUZGjuftaoiWxakVRN0",
		},
		{
			salt:     "JZFIYpBjXTRgXpsP",
			expected: "JZFIYpBjXTRgXpsPgNMpYW_nSVkaMM8unSDJwO5VIj2sTaqfn90v8MHM5tQ",
		},
		{
			salt:     "MmXAtVud3K6pq1XA",
			expected: "MmXAtVud3K6pq1XAXzTM-DnOdD_XWEtZWJaMoDWiXZc8oa8JMoirx00Qthg",
		},
	}

	for i, test := range tests {
		errorPrefix := fmt.Sprintf("Test [%d]: ", i)
		actual := toolset.GenerateToken(test.salt)
		if actual != test.expected {
			t.Errorf(errorPrefix+"Expected %b, got %b", test.expected, actual)
		}
	}
}

func TestToolset_Verify(t *testing.T) {
	tests := []struct {
		saltLen  int
		token    string
		expected bool
	}{
		{
			saltLen:  16,
			token:    "uvqIUfqJ3W0qlszjdJkEzNyB_hee0YLbJJoGIR_voUZGjuftaoiWxakVRN0",
			expected: true,
		},
		{
			saltLen:  8,
			token:    "uvqIUfqJ3W0qlszjdJkEzNyB_hee0YLbJJoGIR_voUZGjuftaoiWxakVRN0",
			expected: false,
		},
		{
			saltLen:  16,
			token:    "JZFIYpBjXTRgXpsPgNMpYW_nSVkaMM8unSDJwO5VIj2sTaqfn90v8MHM5tQ",
			expected: true,
		},
	}

	for i, test := range tests {
		toolset := New(Options{
			Secret:  secret,
			SaltLen: test.saltLen,
		})
		errorPrefix := fmt.Sprintf("Test [%d]: ", i)
		actual := toolset.Verify(test.token)
		if actual != test.expected {
			t.Errorf(errorPrefix+"Expected %d, got %d", test.expected, actual)
		}
	}
}
