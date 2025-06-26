package models

type Configs struct {
	AuthParams     AuthParams     `json:"auth_params"`
	AppParams      AppParams      `json:"app_params"`
	PostgresParams PostgresParams `json:"postgres_params"`
}

type AuthParams struct {
	JwtSecretKey  string `json:"jwt_secret_key"`
	JwtTtlMinutes int    `json:"jwt_ttl_minutes"`
}

type AppParams struct {
	PortRun    string `json:"port_run"`
	ServerName string `json:"server_name"`
	GinMode    string `json:"gin_mode"`
}

type PostgresParams struct {
	User     string `json:"user"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}
