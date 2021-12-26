package env

type AppEnv struct {
	GitHub *GitHub
}

type GitHub struct {
	GraphQLURL  string `envconfig:"GITHUB_GRAPHQL_URL"`
	AccessToken string `envconfig:"GITHUB_ACCESS_TOKEN"`
}
