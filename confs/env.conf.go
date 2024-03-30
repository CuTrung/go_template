package confs

import "github.com/joho/godotenv"

func LoadEnv() map[string]string {
	envs, _ := godotenv.Read(".env")
	return envs
}
