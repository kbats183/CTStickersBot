package ocrapi

import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

type OCRClientConfig struct {
	BaseURL           string `yaml:"base_url"`
	ApiToken          string `yaml:"api_token" env:"OCRAPI_TOKEN"`
	IsOverlayRequired bool   `yaml:"is_overlay_required"`
	OCREngine         int    `yaml:"ocr_engine"`
	Language          string `yaml:"language"`
}

type OCRClient struct {
	config      OCRClientConfig
	restyClient *resty.Client
}

func (client *OCRClient) ParseImage(filePath string) (*ParseAnswer, error) {
	formData := map[string]string{
		"isOverlayRequired": strconv.FormatBool(client.config.IsOverlayRequired),
		"OCREngine":         strconv.Itoa(client.config.OCREngine),
		"language":          client.config.Language,
	}

	var answer ParseAnswer
	_, err := client.restyClient.
		R().
		SetHeader("apikey", client.config.ApiToken).
		SetFormData(formData).
		SetResult(&answer).
		SetFile("file", filePath).
		Post(client.config.BaseURL)

	if err != nil {
		return nil, err
	}
	//cs, e := ans.Result().()

	return &answer, nil
}

func NewOCRClient(config OCRClientConfig) *OCRClient {
	client := resty.New()

	return &OCRClient{
		config:      config,
		restyClient: client,
	}
}

type ParsedResultTextOverlay struct {
	Lines []struct {
		Words []struct {
			WordText string  `json:"WordText"`
			Left     float64 `json:"Left"`
			Top      float64 `json:"Top"`
			Height   float64 `json:"Height"`
			Width    float64 `json:"Width"`
		} `json:"Words"`
		MaxHeight float64 `json:"MaxHeight"`
		MinTop    float64 `json:"MinTop"`
	} `json:"Lines"`
	HasOverlay bool   `json:"HasOverlay"`
	Message    string `json:"Message"`
}

type ParsedResult struct {
	TextOverlay       *ParsedResultTextOverlay `json:"TextOverlay"`
	FileParseExitCode int                      `json:"FileParseExitCode"`
	ParsedText        *string                  `json:"ParsedText"`
	ErrorMessage      *string                  `json:"ErrorMessage"`
	ErrorDetails      *string                  `json:"ErrorDetails"`
}

type ParseAnswer struct {
	ParsedResults                []ParsedResult `json:"ParsedResults"`
	OCRExitCode                  int            `json:"OCRExitCode"`
	IsErroredOnProcessing        bool           `json:"IsErroredOnProcessing"`
	ErrorMessage                 interface{}    `json:"ErrorMessage"`
	ErrorDetails                 string         `json:"ErrorDetails"`
	SearchablePDFURL             string         `json:"SearchablePDFURL"`
	ProcessingTimeInMilliseconds string         `json:"ProcessingTimeInMilliseconds"`
}

func GetStringByParseAnswer(answer *ParseAnswer) string {
	str := ""
	for _, result := range answer.ParsedResults {
		if result.ParsedText != nil {
			str += *result.ParsedText
		}
	}
	return str
}
