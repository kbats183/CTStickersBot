package bot_admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kbats183/CTStickersBot/pkg/core"
	botcontext "github.com/kbats183/CTStickersBot/pkg/core/context"
	"github.com/kbats183/CTStickersBot/pkg/storage"
	"net/http"
)

type BotAdminServer struct {
	Config   core.ServerConfig
	context  botcontext.Context
	storage  *storage.Storage
	router   *gin.Engine
	handlers map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewBotAdminServer(config core.ServerConfig, ctx botcontext.Context, storage *storage.Storage) *BotAdminServer {
	router := gin.New()
	router.Use(gin.Logger())
	return &BotAdminServer{
		Config:  config,
		context: ctx,
		storage: storage,
		router:  router,
	}
}

func (server *BotAdminServer) handlerHome(c *gin.Context) {
	stickerCount, userCount, requestCount, adminCount, err := server.storage.PingDB(server.context.Context)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf(`<html><body>
<h1>CTStickerBot</h1>
<p>Detabase info: %d stickers, %d requests, %d unique users, %d admins</p>
</body></html>`, stickerCount, requestCount, userCount, adminCount)))
}

func (server *BotAdminServer) handlerPing(c *gin.Context) {
	_, _, _, _, err := server.storage.PingDB(server.context.Context)
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
