package context

import (
	"context"
	"github.com/kbats183/CTStickersBot/pkg/ocrapi"
	"go.uber.org/zap"
)

type Context struct {
	context.Context
	Logger *zap.Logger
	OCRClient *ocrapi.OCRClient
}

