package perror

import "fmt"

func TryRollback(err error, rb func() error) error {
	if err == nil {
		return nil
	}
	if rbErr := rb(); rbErr != nil {
		err = fmt.Errorf("%v: %w", rbErr, err)
	}
	return err
}
