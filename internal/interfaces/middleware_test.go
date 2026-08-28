package interfaces

import "testing"

func TestBearerToken(t *testing.T) {
	tests := []struct {
		name   string
		header string
		token  string
		valid  bool
	}{
		{name: "valid", header: "Bearer token", token: "token", valid: true},
		{name: "trims token whitespace", header: "Bearer  token ", token: "token", valid: true},
		{name: "missing scheme", header: "token", valid: false},
		{name: "wrong scheme", header: "Basic token", valid: false},
		{name: "empty token", header: "Bearer ", valid: false},
		{name: "embedded bearer is not stripped", header: "Basic Bearer token", valid: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, valid := bearerToken(tt.header)
			if token != tt.token || valid != tt.valid {
				t.Fatalf("bearerToken(%q) = (%q, %v), want (%q, %v)", tt.header, token, valid, tt.token, tt.valid)
			}
		})
	}
}
