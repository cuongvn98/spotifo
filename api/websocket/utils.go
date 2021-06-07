package websocket

func IndexOfInClientSlice(sl []*Client, el *Client) int {
	for i, v := range sl {
		if v == el {
			return i
		}
	}

	return -1
}
