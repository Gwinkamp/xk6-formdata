package formdata

import (
	"bytes"
	"mime/multipart"
)

// Builder builds formdata body
type Builder struct {
	body   *bytes.Buffer
	writer *multipart.Writer
}

func NewBuilder() *Builder {
	buff := new(bytes.Buffer)
	return &Builder{
		body:   buff,
		writer: multipart.NewWriter(buff),
	}
}

func (b *Builder) Add(key, value string) error {
	return b.writer.WriteField(key, value)
}

func (b *Builder) AddBytes(key string, value []byte) error {
	w, err := b.writer.CreateFormField(key)
	if err != nil {
		return err
	}
	_, err = w.Write(value)
	return err
}

func (b *Builder) AddFile(key, filename string, value []byte) error {
	w, err := b.writer.CreateFormFile(key, filename)
	if err != nil {
		return err
	}
	_, err = w.Write(value)
	return err
}

func (b *Builder) Build() ([]byte, error) {
	err := b.writer.Close()
	if err != nil {
		return nil, err
	}
	return b.body.Bytes(), nil
}

func (b *Builder) GetContentType() string {
	return b.writer.FormDataContentType()
}
