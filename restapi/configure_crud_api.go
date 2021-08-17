// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"test_task/internal/handlers"
	"test_task/internal/pkg/config"

	userss "test_task/internal/repositories/users"
	"test_task/restapi/operations"
	"test_task/restapi/operations/users"

	// postgresql
	_ "github.com/lib/pq"
)

//go:generate swagger generate server --target ..\..\test_kafka_redis_v1 --name CrudAPI --spec ..\api\openapi-spec\swagger.yml --principal interface{}

func configureFlags(api *operations.CrudAPIAPI) {
	type ConfigServer struct {
		Server struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"server"`
	}

	var cfgServer ConfigServer
	config.Get(&cfgServer)
	os.Setenv("HOST", cfgServer.Server.Host)
	os.Setenv("PORT", cfgServer.Server.Port)
}

func configureAPI(api *operations.CrudAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	type ConfigDB struct {
		DB struct {
			Name     string `json:"name"`
			Host     string `json:"host"`
			Port     int    `json:"port"`
			DBname   string `json:"dbname"`
			User     string `json:"user"`
			Password string `json:"password"`
		} `json:"db"`
	}

	var cfgDB ConfigDB
	config.Get(&cfgDB)
	// connstr := fmt.Sprintf("host=%s port=%v user=%s password=%s sslmode=disable", cfgDB.DB.Host, cfgDB.DB.Port, cfgDB.DB.User, cfgDB.DB.Password)
	connstr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", cfgDB.DB.Host, cfgDB.DB.Port, cfgDB.DB.User, cfgDB.DB.Password, cfgDB.DB.DBname)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	repos := Repositories{}
	repos.Users = userss.New(db)

	// Получить всех пользователей
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams) middleware.Responder {
		return handlers.GetAllUsers(repos.Users)
	})

	// Добавить нового пользователя
	api.UsersPostUsersHandler = users.PostUsersHandlerFunc(func(params users.PostUsersParams) middleware.Responder {
		return handlers.AddNewUser(repos.Users, params)
	})

	// Удалить пользователя
	api.UsersDeleteUsersHandler = users.DeleteUsersHandlerFunc(func(params users.DeleteUsersParams) middleware.Responder {
		return handlers.DeleteUser(repos.Users, params)
	})

	// Обновить данные пользователя
	api.UsersPatchUsersHandler = users.PatchUsersHandlerFunc(func(params users.PatchUsersParams) middleware.Responder {
		return handlers.UpdateUser(repos.Users, params)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
