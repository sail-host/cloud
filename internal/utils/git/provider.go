package git

type GitProvider interface {
	CheckAccount() (bool, error)
	GetRepos() ([]Repository, error)
}
