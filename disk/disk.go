package disk

import "io"

func ReadBlock(file io.ReadSeeker, index, bs int64) ([]byte, error) {
	return nil, nil
}

func WriteBlock(file io.WriteSeeker, index, bs int64, data []byte) error {
	return nil
}
