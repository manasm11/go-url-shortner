package inmemory

import "github.com/manasm11/go-url-shortner/shortner"

type inmemoryRepository struct {
	redirects map[string]*shortner.Redirect
}

func NewInMemoryRepository() *inmemoryRepository {
	return &inmemoryRepository{
		redirects: make(map[string]*shortner.Redirect),
	}
}

func (r inmemoryRepository) Find(code string) (*shortner.Redirect, error) {
	if redirect, ok := r.redirects[code]; ok {
		return redirect, nil
	}
	return nil, shortner.ErrRedirectNotFound
}

func (r inmemoryRepository) Store(redirect *shortner.Redirect) error {
	r.redirects[redirect.Code] = redirect
	return nil
}
