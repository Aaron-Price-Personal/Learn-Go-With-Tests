package concurrency

type WebsiteChecker func(string) bool

// Anonymous struct values used when no names needed
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Pushing value onto channel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Pulling value out of channel
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
