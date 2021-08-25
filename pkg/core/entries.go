package core

type AppConfig struct {
	APPName string `default:"CTStickerBot"`

	DB StorageConfig `yaml:"db"`

	TelegramBot BotConfig `yaml:"telegram_bot"`

	OCR OCRClientConfig `yaml:"ocr"`
}

type StorageConfig struct {
	Host     string `yaml:"host" env:"DBHost"`
	UserName string `yaml:"user" env:"DBUser"`
	Password string `yaml:"password" env:"DBPassword"`
	DBName   string `yaml:"name" env:"DBName"`
}

type BotConfig struct {
	BotAuthToken       string `yaml:"auth_token" env:"BOT_AUTH_TOKEN"`
	EnableDebug        bool   `yaml:"enable_debug" env:"BOT_ENABLE_DEBUG" default:"true"`
	InlineStickerLimit int   `yaml:"inline_stickers_limit" env:"BOT_INLINE_STICKER_LIMIT" default:"10"`
}

type OCRClientConfig struct {
	BaseURL           string `yaml:"base_url"`
	ApiToken          string `yaml:"api_token" env:"OCRAPI_TOKEN"`
	IsOverlayRequired bool   `yaml:"is_overlay_required"`
	OCREngine         int    `yaml:"ocr_engine"`
	Language          string `yaml:"language"`
}

type StickerAnswer struct {
	ID           int
	FileID       string
	StickerTitle string
}
