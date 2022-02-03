package bot_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	botcontext "github.com/kbats183/CTStickersBot/pkg/core/context"
	"github.com/kbats183/CTStickersBot/pkg/storage"
	"go.uber.org/zap"
	"net/http"
)

type ServerConfig struct {
	Port string `yaml:"port" env:"PORT"`
}

type BotAdminServer struct {
	Config   ServerConfig
	context  botcontext.Context
	storage  *storage.Storage
	router   *gin.Engine
	handlers map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewBotAdminServer(config ServerConfig, ctx botcontext.Context, storage *storage.Storage) *BotAdminServer {
	router := gin.New()
	router.Use(func(context *gin.Context) {
		ctx.Logger.Debug("HTTP request",
			zap.String("url", context.Request.RequestURI),
		)
	})
	return &BotAdminServer{
		Config:  config,
		context: ctx,
		storage: storage,
		router:  router,
	}
}

func (server *BotAdminServer) handlerHome(c *gin.Context) {
	stickerCount, userCount, requestCount, chosenStickerCount, adminCount, err := server.storage.PingDB(server.context.Context)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf(`<html><body>
<h1>CTStickerBot</h1>
<p>Detabase info: %d stickers, %d requests, %d chosen stickers, %d unique users, %d admins</p>
</body></html>`, stickerCount, requestCount, chosenStickerCount, userCount, adminCount)))
}

func (server *BotAdminServer) handlerPing(c *gin.Context) {
	_, _, _, _, _, err := server.storage.PingDB(server.context.Context)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, "pong")
}

func (server *BotAdminServer) Listen() error {
	server.router.GET("/", server.handlerHome)
	server.router.GET("/ping", server.handlerPing)
	err := server.router.Run(":" + server.Config.Port)
	if err != nil {
		return err
	}
	return nil
}
