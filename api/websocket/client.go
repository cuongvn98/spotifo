package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	writeWait            = 10 * time.Second
	pongWait             = 60 * time.Second
	pingPeriod           = (pongWait * 9) / 10
	maxMessageSize int64 = 512
)

type Client struct {
	hub    *Hub
	send   chan []byte
	conn   *websocket.Conn
	logger *logrus.Logger
	user   User
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.conn.Close(); err != nil {
			c.logger.Errorf("failed when close connection: %s", err)
		}
	}()
	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					c.logger.Errorf("failed when sent close message: %s", err)
				}
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.logger.Errorf("failed when get next writer: %s", w)
				return
			}
			if _, err := w.Write(message); err != nil {
				c.logger.Errorf("failed while write message to client: %s", err)
			}
			if err := w.Close(); err != nil {
				c.logger.Errorf("failed while close message writer: %s", err)
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.logger.Errorf("failed to write ping message: %s", err)
				return
			}
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		if err := c.conn.Close(); err != nil {
			c.logger.Errorf("failed when close connection: %s", err)
		}
	}()
	c.conn.SetReadLimit(maxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		c.logger.Errorf("failed to set read deadline: %s", err)
	}
	c.conn.SetPongHandler(func(string) error {
		if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			c.logger.Errorf("failed to set read deadline: %s", err)
			return err
		}
		return nil
	})
	for {
		typ, reader, err := c.conn.NextReader()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Errorf("socket is close while try to get reader: %s", err)
			}
			break
		}
		if typ != websocket.TextMessage {
			break
		}

		var msg Message

		if err := json.NewDecoder(reader).Decode(&msg); err != nil {
			c.logger.Errorf("failed while decode message")
			break
		}

		// do somethings awesome
	}
}
