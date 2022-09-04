package handlers

import (
	"context"
	"go-redis/db"
	"io/ioutil"
	"net/http"
	"time"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var ctx = context.Background()
		red := db.GetRedisConnection()
		res, err := red.Get(ctx, "user").Result()
		if err != nil {
			http.Error(w, "Error!", 404)
			return

		}
		w.Write([]byte(res))
	case "POST":
		var ctx = context.Background()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		defer r.Body.Close()

		red := db.GetRedisConnection()
		_, err = red.SetNX(ctx, "user", body, 3*time.Second).Result()
		// _, err = red.SetNX(ctx, "user", body, 0).Result()
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return

		}
		w.Write([]byte(body))

	}

}
