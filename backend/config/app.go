package config

import "go.mongodb.org/mongo-driver/v2/mongo"


type App struct {
	Env *Env
	DB mongo.Client
}

func NewApp() (*App, error) {
	env, err := NewEnv()
	if err != nil {
		return nil, err
	}
	db := NewMongoDBConfig(env)
	return &App{
		Env: env,
		DB: *db,
	}, nil
}

func (app *App) CloseDBConnection() {
	CloseMongoDBConnection(&app.DB)
}