package image_utils

import (
	"encoding/base64"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/nu7hatch/gouuid"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func encodeBytesToBase64String(bytes []byte) string {
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	base64String := base64.StdEncoding.EncodeToString(bytes)
	base64Encoding += "data:" + mimeType + ";base64," + base64String

	return base64Encoding
}

func decodeBase64String(base64Encoded string) ([]byte, rest_errors.RestErr) {
	decoded, err := base64.StdEncoding.DecodeString(base64Encoded)
	if err != nil {
		return nil, rest_errors.NewBadRequestError("Error when decoding image")
	}
	return decoded, nil
}

func generateUUID() (string, rest_errors.RestErr) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", rest_errors.NewInternalServerError("Error when generating image id", err)
	}
	return id.String(), nil
}

func writeBytesToFile(path string, content []byte) rest_errors.RestErr {
	file, err := os.Create(path)
	if err != nil {
		return rest_errors.NewInternalServerError("Error when creating file", err)
	}
	defer file.Close()
	if _, err := file.Write(content); err != nil {
		return rest_errors.NewInternalServerError("Error when writing to file", err)
	}
	if err := file.Sync(); err != nil {
		return rest_errors.NewInternalServerError("Error when writing to file", err)
	}
	return nil
}

func readBytesFromFile(path string) ([]byte, rest_errors.RestErr) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("Error when reading file", err)
	}
	return bytes, nil
}

func SaveImage(base64Encoded string, basePath string) (string, rest_errors.RestErr) {
	decoded, err := decodeBase64String(base64Encoded)
	if err != nil {
		return "", err
	}

	id, err := generateUUID()
	imagePath := filepath.Join(basePath, id)

	if err := writeBytesToFile(imagePath, decoded); err != nil {
		return "", err
	}

	return imagePath, nil
}

func LoadImage(path string) (string, rest_errors.RestErr) {
	bytes, err := readBytesFromFile(path)
	if err != nil {
		return "", err
	}
	return encodeBytesToBase64String(bytes), nil
}
