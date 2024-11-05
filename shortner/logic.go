package shortner

import (
	"errors"
	"time"

	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrRedirectNotFound = errors.New("Redirect not found")
	ErrRedirectInvalid  = errors.New("Redirect invalid")
)

type redirectService struct {
	redirectRepo RedirectRepositoryInterface
}

func (s *redirectService) Find(code string) (*Redirect, error) {
	return s.redirectRepo.Find(code)
}
func (s *redirectService) Store(redirect *Redirect) error {

	// check validity. conditions are defined in the template strings of Redirect struct
	if err := validate.Validate(redirect); err != nil {
		return err
	}

	// generate short code
	redirect.Code = shortid.MustGenerate()

	// generate created at timestamp
	loc, _ := time.LoadLocation("Asia/Kolkata")
	redirect.CreatedAt = time.Now().In(loc).Unix()

	return s.redirectRepo.Store(redirect)
}

func NewRedirectService(redirectRepo RedirectRepositoryInterface) RedirectServiceInterface {
	return &redirectService{
		redirectRepo: redirectRepo,
	}
}
