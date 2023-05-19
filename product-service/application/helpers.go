package application

import (
	"log"
	"os"

	"github.com/cristalhq/jwt/v4"
	"github.com/go-playground/validator/v10"
)

func getEnvConfig() (config, error) {
	//get configuration from enviroment and validate
	postgresDsn := os.Getenv("pgDSN")
	jwtSecret := os.Getenv("jwtSecret")
	port := os.Getenv("port")
	config := config{
		jwtSecretKey: jwtSecret,
		//default
		jwtSigningMethod: jwt.Algorithm(jwt.HS256),
		postgresDsn:      postgresDsn,
		port:             port,
	}
	err := config.Validate()
	if err != nil {
		log.Println(err)
		return config, err
	}
	return config, nil
}
func (c *config) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
