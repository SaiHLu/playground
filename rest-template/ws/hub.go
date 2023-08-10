package ws

type message struct {
	data []byte
	room string
}

type hub struct {
	rooms      map[string]map[*connection]bool
	broadcast  chan message
	register   chan subscription
	unregister chan subscription
}

func NewHub() hub {
	return hub{
		rooms:      make(map[string]map[*connection]bool),
		broadcast:  make(chan message),
		register:   make(chan subscription),
		unregister: make(chan subscription),
	}
}

func (h *hub) Run() {
	for {
		select {
		case register := <-h.register:
			connections := h.rooms[register.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[register.room] = connections
			}
			h.rooms[register.room][register.conn] = true

		case unregister := <-h.unregister:
			connections := h.rooms[unregister.room]
			if connections != nil {
				if _, ok := connections[unregister.conn]; ok {
					delete(connections, unregister.conn)
					close(unregister.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, unregister.room)
					}
				}
			}

		case broadcast := <-h.broadcast:
			connections := h.rooms[broadcast.room]

			for c := range connections {
				select {
				case c.send <- broadcast.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, broadcast.room)
					}
				}
			}
		}

	}
}
