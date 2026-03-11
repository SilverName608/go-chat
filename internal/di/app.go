package di

import (
	"fmt"
	"net/http"

	"github.com/SilverName608/go-chat/internal/api"
	"github.com/SilverName608/go-chat/internal/application"
	"github.com/SilverName608/go-chat/internal/config"
	"github.com/SilverName608/go-chat/internal/domain/service"
	"github.com/SilverName608/go-chat/internal/hub"
	"github.com/SilverName608/go-chat/internal/infrastructure/db"
	"github.com/SilverName608/go-chat/internal/infrastructure/repository"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"

	goredis "github.com/redis/go-redis/v9"

	redisClient "github.com/SilverName608/go-chat/internal/infrastructure/redis"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(config.Load),
		fx.Provide(NewPool),
		fx.Provide(NewRedisClient),

		fx.Provide(fx.Annotate(
			repository.NewPostgresUserRepository,
			fx.As(new(repository.UserRepository)),
		)),
		fx.Provide(fx.Annotate(
			repository.NewPostgresRoomRepository,
			fx.As(new(repository.RoomRepository)),
		)),
		fx.Provide(fx.Annotate(
			repository.NewPostgresMessageRepository,
			fx.As(new(repository.MessageRepository)),
		)),

		fx.Provide(fx.Annotate(
			application.NewUserService,
			fx.As(new(service.UserService)),
		)),
		fx.Provide(fx.Annotate(
			application.NewRoomService,
			fx.As(new(service.RoomService)),
		)),
		fx.Provide(fx.Annotate(
			application.NewMessageService,
			fx.As(new(service.MessageService)),
		)),

		fx.Provide(hub.NewHub),

		fx.Provide(api.NewMiddleware),
		fx.Provide(api.NewUserHandler),
		fx.Provide(api.NewRoomHandler),
		fx.Provide(api.NewWSHandler),
		fx.Provide(api.NewRouter),

		fx.Invoke(RunServer),
	)
}

func RunServer(router chi.Router, cfg *config.Config) {
	if err := runMigrations(cfg); err != nil {
		panic(err)
	}
	fmt.Printf("Server launch → http://localhost:%s\n", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, router); err != nil {
		panic(err)
	}
}

func runMigrations(cfg *config.Config) error {
	m, err := migrate.New(
		"file://migrations",
		cfg.DBDSN,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func NewPool(cfg *config.Config) (*pgxpool.Pool, error) {
	return db.NewPool(cfg.DBDSN)
}

func NewRedisClient(cfg *config.Config) (*goredis.Client, error) {
	return redisClient.NewClient(cfg.RedisDSN)
}
