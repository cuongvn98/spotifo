package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"spotifo/types"
)

type Authorization struct {
	Next     http.Handler
	Endpoint string
}

func (a Authorization) Middleware(next http.Handler) http.Handler {
	a.Next = next
	return a
}

func (a Authorization) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")
	if len(authorization) <= 12 {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write(nil)
		return
	}
	user, err := a.fetchUserInfo(authorization)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	ctx := r.Context()
	ctx = context.WithValue(ctx, "user", user)
	r = r.WithContext(ctx)
	a.Next.ServeHTTP(w, r)
}

func (a Authorization) fetchUserInfo(token string) (types.User, error) {
	req, err := http.NewRequest("GET", a.Endpoint, nil)
	if err != nil {
		return types.User{}, err
	}
	req.Header.Add("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return types.User{}, err
	}

	var data struct {
		User []types.User `json:"user"`
	}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return types.User{}, err
	}
	if len(data.User) != 0 {
		return data.User[0], nil
	}

	return types.User{}, fmt.Errorf("user not found")
}
