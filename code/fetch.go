func (s *Client) Fetch(key string) <-chan *Tuple {
	sock, _ := sub.NewSocket()

	si := Crc16(key) % 4
	uri := servers[i]

	sock.SetOption(mangos.OptionSubscribe, key)
	items := make(chan *Tuple, 64)

	go func() {
		res, _ := http.Get(uri + key)
		json.Unmarshal(b, &xs)
		for i := 0; i < len(xs); i++ {
			items <- xs[i]
		}

		for {
			msg, err := sock.Recv()
			json.Unmarshal(body, x)
			items <- x
		}
	}()

	return items
}
