package shortner

type RedirectRepositoryInterface interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
