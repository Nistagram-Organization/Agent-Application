package image_utils

import (
	"bytes"
	"encoding/base64"
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/nu7hatch/gouuid"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func encodeBytesToBase64String(bytes []byte) string {
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	base64String := base64.StdEncoding.EncodeToString(bytes)
	base64Encoding += "data:" + mimeType + ";base64," + base64String

	return base64Encoding
}

func returnImageType(base64Encoded string) string {
	return strings.TrimSuffix(base64Encoded[5:strings.Index(base64Encoded, ",")], ";base64")
}

func decodeBase64String(base64Encoded string, imageType string) (image.Image, rest_errors.RestErr) {
	var image image.Image
	var err error

	trimmmed := base64Encoded[strings.Index(base64Encoded, ",")+1:]
	decoded, decodingErr := base64.StdEncoding.DecodeString(trimmmed)
	if decodingErr != nil {
		return nil, rest_errors.NewBadRequestError("Error when decoding image")
	}
	reader := bytes.NewReader(decoded)
	if imageType == "image/png" {
		image, err = png.Decode(reader)
	} else {
		image, err = jpeg.Decode(reader)
	}
	if err != nil {
		return nil, rest_errors.NewBadRequestError("Error when decoding image")
	}
	return image, nil
}

func generateUUID() (string, rest_errors.RestErr) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", rest_errors.NewInternalServerError("Error when generating image id", err)
	}
	return id.String(), nil
}

func writeImageToFile(path string, image image.Image, imageType string) rest_errors.RestErr {
	file, err := os.Create(path)
	if err != nil {
		return rest_errors.NewInternalServerError("Error when creating file", err)
	}
	defer file.Close()

	if imageType == "image/png" {
		if err := png.Encode(file, image); err != nil {
			return rest_errors.NewInternalServerError("Error when writing to file", err)
		}
	} else {
		if err := jpeg.Encode(file, image, nil); err != nil {
			return rest_errors.NewInternalServerError("Error when writing to file", err)
		}
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
	imageType := returnImageType(base64Encoded)
	image, err := decodeBase64String(base64Encoded, imageType)
	if err != nil {
		return "", err
	}

	id, err := generateUUID()
	imageName := id + "." + strings.Split(imageType, "/")[1]
	imagePath := filepath.Join(basePath, imageName)

	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		os.MkdirAll(basePath, 0700)
	}

	if err := writeImageToFile(imagePath, image, imageType); err != nil {
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
