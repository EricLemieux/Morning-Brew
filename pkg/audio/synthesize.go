package audio

import "github.com/ericlemieux/morning-brew/pkg/storage"

// SynthesizeAudioToUrl - take the formatted script, create an audio file and store in a publicly accessible location.
func SynthesizeAudioToUrl(store storage.Storage, text string) (string, error) {
	audio, err := synthesizeAudioToByteStream(text)
	if err != nil {
		return "", err
	}

	url, err := storeAudioFileInBucket(store, audio)
	if err != nil {
		return "", err
	}

	return url, nil
}

func synthesizeAudioToByteStream(text string) ([]byte, error) {
	// TODO

	return nil, nil
}

func storeAudioFileInBucket(store storage.Storage, audio []byte) (string, error) {
	url, err := store.Upload(audio)
	if err != nil {
		return "", err
	}

	return url, nil
}
