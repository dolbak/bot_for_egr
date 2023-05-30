package storage

import (
	"crypto/sha1"
	"fmt"
	"io"
	"myGoApp/lib/e"
)

type Storage interface {
	Save(p *Messege) error
	IsExist(p *Messege) (bool, error)
}

type Messege struct {
	Text     string
	UserName string
}

func (p Messege) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.Text); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
