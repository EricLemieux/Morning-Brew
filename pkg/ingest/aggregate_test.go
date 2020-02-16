package ingest

import "testing"

func TestGetText(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Smoke test",
			want:    "Hello world. I am some sample text.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetText()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetText() got = %v, want %v", got, tt.want)
			}
		})
	}
}
