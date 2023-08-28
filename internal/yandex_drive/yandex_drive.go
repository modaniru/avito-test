package yandexdrive

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	yadisk "github.com/nikitaksv/yandex-disk-sdk-go"
)

//go:generate mockery --name Disk
type Disk interface {
	CreateFile(name string, data []byte) (string, error)
	IsAvailible() bool
}

type YandexDisk struct {
	disk        yadisk.YaDisk
	client      *http.Client
	isAvailible bool
}

func NewYandexDisk(token string) (*YandexDisk, error) {
	if token == "" {
		return &YandexDisk{isAvailible: false}, nil
	}
	yaDisk, err := yadisk.NewYaDisk(context.Background(), http.DefaultClient, &yadisk.Token{AccessToken: token})
	if err != nil {
		return nil, err
	}
	return &YandexDisk{disk: yaDisk, client: http.DefaultClient, isAvailible: true}, nil
}

func (y *YandexDisk) CreateFile(name string, data []byte) (string, error) {
	op := "internal.yandex_drive.YandexDisk.CreateFile"
	name = name + ".csv"

	link, err := y.disk.GetResourceUploadLink(name, nil, true)
	if err != nil {
		return "", fmt.Errorf("%s get upload link error: %w", op, err)
	}
	request, err := http.NewRequest(http.MethodPut, link.Href, bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("%s create request error: %w", op, err)
	}
	_, err = y.client.Do(request)
	if err != nil {
		return "", fmt.Errorf("%s invoke request error: %w", op, err)
	}

	responseLink, err := y.disk.GetResourceDownloadLink(name, nil)
	if err != nil {
		return "", fmt.Errorf("%s get resource link error: %w", op, err)
	}
	return responseLink.Href, nil
}

func (y *YandexDisk) IsAvailible() bool {
	return y.isAvailible
}
