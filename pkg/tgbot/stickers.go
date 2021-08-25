package tgbot

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func prepareStickerToParsing(stickerUrl string) (string, error) {
	err := makeTemporaryStickerDir()
	if err != nil {
		return "", err
	}

	_, err = url.Parse(stickerUrl)
	if err != nil {
		return "", err
	}
	//pathSegments := strings.Split(fileURL.Path, ".")
	fileName := "./tmp/"+ strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".webp"

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(stickerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	defer func() { _ = file.Close() }()

	return fileName, nil
}

func makeTemporaryStickerDir() error {
	tmpDirPath := "./tmp/"
	if _, err := os.Stat(tmpDirPath); os.IsNotExist(err) {
		err := os.Mkdir(tmpDirPath, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
