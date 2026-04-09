package exercise1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	t.Run("Succeful Get Request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello")
		}))

		defer server.Close()

		got, err := Fetch(context.Background(), server.URL)

		if err != nil {
			t.Fatal(err)
		}
		if string(got) != "hello" {
			t.Errorf("got %s want hello", got)
		}
	})
	t.Run("context gets cancelled", func(t *testing.T) {
		cancellingCtx, cancel := context.WithCancel(t.Context())
		cancel()

		got, err := Fetch(cancellingCtx, "/")
		if got != nil {
			t.Error("expected no output")
		}
		if !errors.Is(err, ErrCancelled) {
			t.Fatalf("expected %v, got %v", ErrCancelled, err)
		}
	})
	t.Run("times out when server is slow", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			select{
			case <- ctx.Done():
				return
			case <-time.After(5 * time.Second):
				fmt.Fprint(w, "too late")
			}
		}))
		defer server.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		got, err := Fetch(ctx, server.URL)
		if got != nil {
			t.Error("expected no output")
		}
		if !errors.Is(err, ErrTimeout) {
			t.Fatalf("expected %v, got %v", ErrTimeout, err)
		}
	})
}
