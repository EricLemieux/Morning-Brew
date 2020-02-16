package audio

import (
	"github.com/ericlemieux/morning-brew/pkg/storage"
	"testing"
)

type mockStorage struct{}

func (store mockStorage) Upload(file []byte) (string, error) {
	return "https://some.file.url/abc123", nil
}

func TestSynthesizeAudioToUrl(t *testing.T) {
	type args struct {
		text  string
		store storage.Storage
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Smoke test",
			args:    args{text: "input script", store: mockStorage{}},
			want:    "https://some.file.url/abc123",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SynthesizeAudioToUrl(tt.args.store, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("SynthesizeAudioToUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SynthesizeAudioToUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
