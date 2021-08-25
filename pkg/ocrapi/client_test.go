package ocrapi

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func prepareClient() *OCRClient {
	return NewOCRClient(OCRClientConfig{
		BaseURL:           "https://api.ocr.space/parse/image",
		ApiToken:          os.Getenv("OCRAPI_TOKEN"),
		Language:          "rus",
		OCREngine:         1,
		IsOverlayRequired: false,
	})
}

func TestWithOutToken(t *testing.T) {
	client := NewOCRClient(OCRClientConfig{
		BaseURL: "https://api.ocr.space/parse/image",
	})
	image, err := client.ParseImage("extra/test/ocr_test1.jpg")
	require.NoError(t, err)
	require.Equal(t, 0, image.OCRExitCode)
	require.Nil(t, image.ErrorMessage)
	require.Nil(t, image.ParsedResults)
}

func TestEmptyImage(t *testing.T) {
	client := prepareClient()
	image, err := client.ParseImage("extra/test/ocr_test1.jpg")
	require.NoError(t, err)
	require.Nil(t, image.ErrorMessage)
	require.Equal(t, 1, image.OCRExitCode)
	require.Len(t, image.ParsedResults, 1)
	require.Equal(t, 1, image.ParsedResults[0].FileParseExitCode)
	require.NotNil(t, image.ParsedResults[0].ParsedText)
	require.Equal(t, "", *image.ParsedResults[0].ParsedText)
}

func TestSmallText(t *testing.T) {
	client := prepareClient()
	image, err := client.ParseImage("extra/test/ocr_test2.jpg")
	require.NoError(t, err)
	require.Nil(t, image.ErrorMessage)
	require.Equal(t, 1, image.OCRExitCode)
	require.Len(t, image.ParsedResults, 1)
	require.Equal(t, 1, image.ParsedResults[0].FileParseExitCode)
	require.NotNil(t, image.ParsedResults[0].ParsedText)
	require.Equal(t, "small simple text", strings.TrimSpace(*image.ParsedResults[0].ParsedText))
	x, _ := json.Marshal(image)
	fmt.Println(string(x))
}

func TestBigText(t *testing.T) {
	client := prepareClient()
	image, err := client.ParseImage("extra/test/ocr_test3.jpg")
	require.NoError(t, err)
	require.Nil(t, image.ErrorMessage)
	require.Equal(t, 1, image.OCRExitCode)
	require.Len(t, image.ParsedResults, 1)
	require.Equal(t, 1, image.ParsedResults[0].FileParseExitCode)
	require.NotNil(t, image.ParsedResults[0].ParsedText)
	require.Equal(t, "текст немного больше u русский", strings.TrimSpace(*image.ParsedResults[0].ParsedText))
	x, _ := json.Marshal(image)
	fmt.Println(string(x))
}

func TestParseSticker(t *testing.T) {
	client := prepareClient()
	image, err := client.ParseImage("extra/test/ocr_test4.webp")
	require.NoError(t, err)
	require.Nil(t, image.ErrorMessage)
	require.Equal(t, 1, image.OCRExitCode)
	require.Len(t, image.ParsedResults, 1)
	require.Equal(t, 1, image.ParsedResults[0].FileParseExitCode)
	require.NotNil(t, image.ParsedResults[0].ParsedText)
	require.Equal(t, "Кирилл Пешков\r\n"+
		"Я считаю, что возможность\r\n" +
		"получить 4 должна быть у самого\r\n" +
		"тупого, но усердного человека", strings.TrimSpace(*image.ParsedResults[0].ParsedText))
}
