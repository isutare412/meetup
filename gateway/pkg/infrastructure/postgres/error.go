package postgres

import (
	"errors"
	"regexp"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

var (
	// Format: Key ({column})=({value}) already exists.
	regexDuplicate = regexp.MustCompile(`\(([^\s]+)\)=\(([^\s]+)\)`)
)

func isErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)

}

func isErrDuplicateKey(err error) (key, val string) {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return
	} else if pgErr.Code != "23505" {
		return
	}

	matches := regexDuplicate.FindStringSubmatch(pgErr.Detail)
	if len(matches) < 3 {
		return
	}
	key = matches[1]
	val = matches[2]
	return
}
