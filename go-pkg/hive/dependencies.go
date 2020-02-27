package hive

func (b *Bee) ConcatDeps (bees []*Bee) *Bee {
	urls := BeesToURLs(bees)
	for _, url := range urls {
		if !contains(b.Deps, url.Str()) {
			b.Deps = append(b.Deps, url.Str())
		}
	}
	return b
}


func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
