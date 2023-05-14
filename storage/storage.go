package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"urler/lib/e"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userNmae string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

var ErrNoSavedPages = errors.New("no saved page")

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}
	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}
	return fmt.Sprint(h.Sum(nil)), nil
}
