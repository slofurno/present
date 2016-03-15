func fetch(user string) chan *Tuple {
	client := &Client{}
	subs := client.Fetch(user)
	res := make(chan *Tuple, 64)

	push := func(xs <-chan *Tuple) {
		for x := range xs {
			res <- x
		}
	}

	go func() {
		for s := range subs {
			c := client.Fetch(string(s.Key))
			go push(c)
		}
	}()

	return res
}

