package audio

import (
	"testing"
)

func TestSynthesizeAudioToUrl(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Smoke test",
			args:    args{text: "input script"},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SynthesizeAudioToUrl(tt.args.text)
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
