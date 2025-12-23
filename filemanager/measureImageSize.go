package filemanager

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strings"
	"time"

	_ "golang.org/x/image/webp"
)

type ImageSize struct {
	Width  int
	Height int
}

func MeasureLocalImageSize(path string) (*ImageSize, error) {
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, os.ErrNotExist
		}
		fmt.Println("error in filemanager/MeasureImageSize/os.Open")
		return nil, err
	}
	defer f.Close()

	cfg, format, err := image.DecodeConfig(f)
	if err != nil {
		if strings.Contains(err.Error(), "unknown format") {
			fmt.Println("format:" + format)
			_ = WriteJson(format, "notion_data/not_supported_img_format.json")
			return &ImageSize{
				Width:  0,
				Height: 0,
			}, nil
		}
		fmt.Println("error in filemanager/MeasureImageSize/image.DecodeConfig format:" + format)
		return nil, err
	}

	return &ImageSize{
		Width:  cfg.Width,
		Height: cfg.Height,
	}, nil
}

var (
	ErrRequestFailed = errors.New("request failed")
	ErrHTTPStatus    = errors.New("invalid http status")
	ErrNotImage      = errors.New("content is not image")
	ErrDecodeImage   = errors.New("failed to decode image")
)

func MeasureImageSizeFromURL(url string) (*ImageSize, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}
	defer resp.Body.Close()

	// 404 / 403 など
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %d", ErrHTTPStatus, resp.StatusCode)
	}

	// Content-Type チェック
	ct := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "image/") {
		return nil, fmt.Errorf("%w: %s", ErrNotImage, ct)
	}

	cfg, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDecodeImage, err)
	}

	return &ImageSize{
		Width:  cfg.Width,
		Height: cfg.Height,
	}, nil
}
