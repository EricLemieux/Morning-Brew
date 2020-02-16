package audio

// SynthesizeAudioToUrl - take the formatted script, create an audio file and store in a publicly accessible location.
func SynthesizeAudioToUrl(text string) (string, error) {
	audio, err := synthesizeAudioToByteStream(text)
	if err != nil {
		return "", err
	}

	url, err := storeAudioFileInBucket(audio)
	if err != nil {
		return "", err
	}

	return url, nil
}

func synthesizeAudioToByteStream(text string) ([]byte, error) {
	// TODO

	return nil, nil
}

func storeAudioFileInBucket(audio []byte) (string, error) {
	// TODO

	return "", nil
}
