package encoders

func Marshal(encoder Encoder, v interface{}) ([]byte, error) {
	return encoder.Marshal(v)
}

func Unmarshal(encoder Encoder, data []byte, v interface{}) error {
	return encoder.Unmarshal(data, v)
}
