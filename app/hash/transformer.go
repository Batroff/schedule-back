package hash

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/pkg/errors"
	"os"
)

func ByteTransform(bytes []byte) string {
	hash := sha1.New()
	hash.Write(bytes)
	sha1Hash := hex.EncodeToString(hash.Sum(nil))

	return sha1Hash
}

func ExcelTransform(filepath string) (string, error) {
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return "", errors.Wrapf(err, "File read error %s", filepath)
	}
	return ByteTransform(fileBytes), nil
}

func ExcelManyTransform(paths []string) ([]string, error) {
	transformed := make([]string, len(paths))

	for i, path := range paths {
		h, err := ExcelTransform(path)
		if err != nil {
			return nil, err
		}

		transformed[i] = h
	}

	return transformed, nil
}
