package services

import (
	"github.com/Nistagram-Organization/Agent-Application/src/utils/rest_errors"
	"github.com/stretchr/testify/mock"
)

type ImageUtilsServiceMock struct {
	mock.Mock
}

func (i *ImageUtilsServiceMock) SaveImage(s string, s2 string) (string, rest_errors.RestErr) {
	args := i.Called(s, s2)
	if args.Get(1) == nil {
		return args.String(0), nil
	}
	return args.String(0), args.Get(1).(rest_errors.RestErr)
}

func (i *ImageUtilsServiceMock) LoadImage(s string) (string, rest_errors.RestErr) {
	args := i.Called(s)
	return args.String(0), nil
}
