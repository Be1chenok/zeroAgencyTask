package server

import (
	"context"
	"fmt"

	"github.com/Be1chenok/zeroAgencyTask/internal/config"
	"github.com/Be1chenok/zeroAgencyTask/internal/delivery/handler"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	handler  handler.Handler
	fiberApp *fiber.App
}

func New(conf *config.Config, handler handler.Handler) *Server {
	return &Server{
		fiberApp: fiber.New(),
		handler:  handler,
	}
}

func (srv *Server) InitRoutes() {
	srv.fiberApp.Post("/edit/:id", srv.handler.EditNewsById)
	srv.fiberApp.Get("/list", srv.handler.GetNewsList)
}

func (srv *Server) Start(conf *config.Config) error {
	if err := srv.fiberApp.Listen(fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)); err != nil {
		return err
	}

	return nil
}

func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.fiberApp.ShutdownWithContext(ctx); err != nil {
		return err
	}
	return nil
}
