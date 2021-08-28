package core

type AppConfig struct {
	APPName string `default:"CTStickerBot"`

	DB StorageConfig `yaml:"db"`

	TelegramBot BotConfig `yaml:"telegram_bot"`

	OCR OCRClientConfig `yaml:"ocr"`

	ServerConfig ServerConfig `yaml:"server"`
}

type StorageConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	UserName string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"name" env:"DB_NAME"`
}

type BotConfig struct {
	BotAuthToken       string `yaml:"auth_token" env:"BOT_AUTH_TOKEN"`
	EnableDebug        bool   `yaml:"enable_debug" env:"BOT_ENABLE_DEBUG" default:"true"`
	InlineStickerLimit int    `yaml:"inline_stickers_limit" env:"BOT_INLINE_STICKER_LIMIT" default:"10"`
}

type OCRClientConfig struct {
	BaseURL           string `yaml:"base_url"`
	ApiToken          string `yaml:"api_token" env:"OCRAPI_TOKEN"`
	IsOverlayRequired bool   `yaml:"is_overlay_required"`
	OCREngine         int    `yaml:"ocr_engine"`
	Language          string `yaml:"language"`
}

type ServerConfig struct {
	Port string `yaml:"port" env:"PORT"`
}

type StickerAnswer struct {
	ID           int
	FileID       string
	StickerTitle string
}
