package cli

import (
	"reflect"
	"testing"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		want     Config
		wantErr  bool
		errorMsg string
	}{
		{
			name: "valid filepath with unique flag",
			args: []string{"testdata/input.txt", "-u"},
			want: Config{
				FilePath: "testdata/input.txt",
				Unique:   true,
			},
			wantErr: false,
		},
		{
			name: "valid filepath without unique flag",
			args: []string{"testdata/input.txt"},
			want: Config{
				FilePath: "testdata/input.txt",
				Unique:   false,
			},
			wantErr: false,
		},
		{
			name:     "missing filepath",
			args:     []string{},
			want:     Config{},
			wantErr:  true,
			errorMsg: "filepath is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := ""
			if len(tt.args) > 0 {
				filePath = tt.args[0]
			}
			uniq := false
			if len(tt.args) > 1 {
				uniq = true
			}
			got, err := ParseFlags(filePath, uniq)

			// Check if we expected an error
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseFlags() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("ParseFlags() error = %v, want %v", err, tt.errorMsg)
					return
				}
				return
			}

			// If we didn't expect an error, but got one
			if err != nil {
				t.Errorf("ParseFlags() unexpected error = %v", err)
				return
			}

			// Compare the results
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
