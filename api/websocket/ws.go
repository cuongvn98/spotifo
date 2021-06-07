package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"spotifo/types"
)

type WS struct {
	upgrader *websocket.Upgrader
	logger   *logrus.Logger
	hub      *Hub
}

func NewWS(logger *logrus.Logger) *WS {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	hub := NewHub()
	return &WS{logger: logger, upgrader: upgrader, hub: hub}
}

func (ws WS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	iUser := r.Context().Value("iUser")
	if iUser == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("UnAuthorized"))
		return
	}
	var user User

	if u, ok := iUser.(types.User); ok {
		user = User{
			Id:    u.Id,
			Email: user.Email,
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("UnAuthorized"))
		return
	}

	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		ws.logger.Errorf("failed to upgrade connection :%s \n", err)
		return
	}

	client := &Client{
		hub:  ws.hub,
		send: make(chan []byte, 256),
		conn: conn,
		user: user,
	}
	ws.hub.register <- client

	go client.writePump()
	go client.readPump()

}
