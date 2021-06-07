package websocket

type Hub struct {
	clients    map[*Client]bool
	users      map[User][]*Client
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		users:      make(map[User][]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			// add client to user group
			if _, ok := h.users[client.user]; !ok {
				h.users[client.user] = make([]*Client, 0)
			}
			h.users[client.user] = append(h.users[client.user], client)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)

			}
			if clients, ok := h.users[client.user]; ok {
				//delete client from user
				index := IndexOfInClientSlice(clients, client)
				if index != -1 {
					clients[index] = clients[len(clients)-1]
					h.users[client.user] = clients[:len(clients)-1]
				}
			}
			close(client.send)
		}
	}
}
