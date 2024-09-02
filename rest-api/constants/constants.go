package constants

type EnvVariableNames map[string]string

var ENV_VARIABLE_NAMES EnvVariableNames = map[string]string{
	"DB_PATH":           "DB_PATH",
	"DB_MAX_OPEN_CONNS": "DB_MAX_OPEN_CONNS",
	"DB_MAX_IDLE_CONNS": "DB_MAX_IDLE_CONNS",
	"JWT_SECRET":        "JWT_SECRET",
}
