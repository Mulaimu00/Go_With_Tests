package exercise1

import (
	"context"
	"errors"
	"io"
	"net/http"
)

var (
	ErrTimeout = errors.New("Time ran out")
	ErrCancelled = errors.New("user cancelled request")
)

func Fetch(ctx context.Context, url string) ([]byte, error) {
	cancelErr := ctx.Err()
	if cancelErr != nil {
		return nil ,ErrCancelled
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded){
			return nil, ErrTimeout
		}
		return  nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

