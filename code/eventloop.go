for {
		select {
		case f := <-c.events:
			f()

		case <-c.close:
			break
		}
	}
