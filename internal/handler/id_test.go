package handler

import "testing"

func TestParseID(t *testing.T) {
	tests := []struct {
		arg     string
		want    int
		wantErr bool
	}{
		{"1", 1, false},
		{"42", 42, false},
		{"0", 0, true},
		{"-3", 0, true},
		{"abc", 0, true},
		{"", 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.arg, func(t *testing.T) {
			got, err := parseID(tc.arg)

			if (err != nil) != tc.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("id = %d, want %d", got, tc.want)
			}
		})
	}
}
