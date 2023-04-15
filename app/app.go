package app

import (
	"bbs/common/errorx"
	"bbs/common/middleware"
	"bbs/config"
	"bbs/ent"
	"bbs/user"
	"context"
	"github.com/casbin/casbin/v2"
	entadapter "github.com/casbin/ent-adapter"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type App struct {
	Server *rest.Server
	Config config.Config
	Ent    *ent.Client
	Casbin *casbin.Enforcer
}

func (a *App) register() {
	user.NewUserServie(a.Config, a.Server, a.Ent)
}
func (a *App) Run() {
	a.register()
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	a.Server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/static/:file",
		Handler: http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))).ServeHTTP,
	})
	a.Server.Use(middleware.AuthMiddleware)
	a.Server.Start()
}
func NewApp(ser *rest.Server, c config.Config, db *ent.Client, enforcer *casbin.Enforcer) (*App, func()) {
	app := &App{
		Server: ser,
		Config: c,
		Ent:    db,
		Casbin: enforcer,
	}
	return app, func() {
		app.Server.Stop()
	}
}
func NewServer(c config.Config) *rest.Server {
	return rest.MustNewServer(c.RestConf)
}
func NewCasbin(c config.Config) (*casbin.Enforcer, error) {
	a, _ := entadapter.NewAdapter(c.DB.Type, c.DB.Source)
	return casbin.NewEnforcer(c.Model, a)
}
func NewDB(c config.Config) (*ent.Client, func(), error) {
	var db *ent.Client
	var err error
	logx.Debugf("db:%s", c.DB.Source)
	db, err = ent.Open(c.DB.Type, c.DB.Source)
	if err != nil {
		logx.Errorf("failed opening connection to mysql: %v", err)
		return nil, func() {

		}, err
	}
	if err := db.Schema.Create(context.Background()); err != nil {
		logx.Errorf("failed creating schema resources: %v", err)
		return nil, func() {

		}, err
	}
	return db, func() {
		db.Close()
	}, err
}
