package c14v1context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct{
	response string
	cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "hello, world"


	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf("got %s want %s", res.Body.String(), data)
		}

		store.assertWasNotCancceled(store)
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		store.assertWasCancelled(store)
	})
}

func (s *SpyStore) assertWasCancelled(store *SpyStore) {
	s.t.Helper()

	if !store.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancceled(store *SpyStore) {
	s.t.Helper()

	if store.cancelled {
		s.t.Error("it should not have cancelled the store")
	}
}

