package upload

import (
	"io"
	"mime/multipart"
)

// Tool ...
type Tool interface {
	UploadImage(fileHeader *multipart.FileHeader) (string, error)
	CheckUploadImageAllowExt(ext string) bool
	CheckUploadImageMaxSize(file io.Reader) bool
}

// Setting ...
type Setting interface {
	GetUploadSavePath() string
	GetUploadImageAllowExts() []string
	GetUploadImageMaxSize() int
}
