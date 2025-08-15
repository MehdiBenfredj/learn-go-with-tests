package context

import (
	"context"
	"net/http"
)

//func Server(store Store) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		ctx := r.Context()
//		ch := make(chan string, 1)
//		go func() {
//			res, _ := store.Fetch(ctx)
//			w.Write([]byte(res))
//			ch <- res
//		}()
//		select {
//		case <-ch:
//		case <-ctx.Done():
//		}
//
//	}
//}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		res, err := store.Fetch(ctx)
		if err == nil {
			w.Write([]byte(res))
		}
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
