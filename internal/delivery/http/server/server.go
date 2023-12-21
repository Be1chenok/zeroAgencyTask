package server

import (
	"context"
	"fmt"

	"github.com/Be1chenok/zeroAgencyTask/internal/config"
	"github.com/Be1chenok/zeroAgencyTask/internal/delivery/http/handler"
	"github.com/gofiber/fiber/v2"
	_ "github.com/swaggo/fiber-swagger/example/docs"
)

type Server struct {
	handler  handler.Handler
	fiberApp *fiber.App
}

func New(conf *config.Config, handler handler.Handler) *Server {
	return &Server{
		fiberApp: fiber.New(fiber.Config{
			ReadTimeout:  conf.Server.ReadTimeout,
			WriteTimeout: conf.Server.WriteTimeout,
		}),
		handler: handler,
	}
}

func (srv *Server) InitRoutes() {
	srv.fiberApp.Post("/register", srv.handler.Register)
	srv.fiberApp.Post("/login", srv.handler.Login)

	secure := srv.fiberApp.Group("/", srv.handler.UserAccessIdentity)
	secure.Get("/logout", srv.handler.LogOut)
	secure.Get("/fullLogout", srv.handler.FullLogOut)
	secure.Get("/refresh", srv.handler.Refresh)

	secure.Post("/edit/:id", srv.handler.EditNewsById)
	secure.Get("/list", srv.handler.GetNewsList)
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
