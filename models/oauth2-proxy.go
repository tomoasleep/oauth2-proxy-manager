package models

// ServiceSettings - Settings from individual services annotations.
type ServiceSettings struct {
	IngressClass string
	AuthURL      string
	AuthSignIn   string
	GitHub       GitHubProvider
}

// GitHubProvider - GitHub Provicer
type GitHubProvider struct {
	Organization string
	Teams        []string
}
