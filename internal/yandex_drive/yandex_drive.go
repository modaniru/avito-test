package yandexdrive

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	yadisk "github.com/nikitaksv/yandex-disk-sdk-go"
)

type YandexDisk struct {
	disk   yadisk.YaDisk
	client *http.Client
}

func NewYandexDisk(token string) (*YandexDisk, error) {
	yaDisk, err := yadisk.NewYaDisk(context.Background(), http.DefaultClient, &yadisk.Token{AccessToken: token})
	if err != nil {
		return nil, err
	}
	return &YandexDisk{disk: yaDisk, client: http.DefaultClient}, nil
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
