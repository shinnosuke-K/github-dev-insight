package env

type AppEnv struct {
	GitHub *GitHub
	AWS    *AWS
	RDB    *RDB
}

type GitHub struct {
	GraphQLURL  string `envconfig:"GITHUB_GRAPHQL_URL"`
	AccessToken string `envconfig:"GITHUB_ACCESS_TOKEN"`
	LoginUser   string `envconfig:"GITHUB_LOGIN_USER"`
}

type AWS struct {
	SecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY" default:"fake"`
	AccessToken     string `envconfig:"AWS_ACCESS_TOKEN" default:"fake"`
	Region          string `envconfig:"AWS_REGION" default:"ap-northeast-1"`
	EndPoint        string `envconfig:"AWS_END_POINT" default:"http://localhost:8000"`
}

type RDB struct {
	Driver   string `envconfig:"RDB_DRIVER"`
	User     string `envconfig:"RDB_USER"`
	Password string `envconfig:"RDB_PASSWORD"`
	Host     string `envconfig:"RDB_HOST"`
	Port     string `envconfig:"RDB_PORT"`
	Name     string `envconfig:"RDB_NAME"`
}
