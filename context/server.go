package context

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//ctx := r.Context()
		//
		//data := make(chan string, 1)
		//
		//go func() {
		//	data <- store.Fetch()
		//}()
		//
		//select {
		//case d := <-data:
		//	fmt.Fprint(w, d)
		//case <-ctx.Done():
		//	store.Cancel()
		//}
		data, err := store.Fetch(r.Context())

		if err != nil {
			log.Println("error encountered")
			return
		}
		fmt.Fprint(w, data)
	}
}
