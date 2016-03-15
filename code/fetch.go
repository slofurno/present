func (s *Client) Fetch(key string) <-chan *Tuple {
	idx := crc16(key) % 4
	sock.Dial(servers[idx])
	sock.SetOption(mangos.OptionSubscribe, key)
	
	go func() {
		http.Get(uri + key)
	

}
