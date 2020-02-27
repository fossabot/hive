package hive

import (
	"fmt"
	"strings"
)

type URL string
type URLs []URL
type Deps []string

func implode(author, name string) URL {
	return URL(fmt.Sprintf("%s/%s", author, name))
}

func (url URL) Str() string {
	return string(url)
}

func (urls URLs) Array() []string {
	arr := make([]string, len(urls))
	for i, url := range urls{
		arr[i] = string(url)
	}
	return arr
}

func appendOne(dest URLs, src URL) URLs {
	if !dest.contains(src) {
		return append(dest, src)
	}
	return dest
}

func AppendURL(dest URLs, src ...URL) URLs {
	if len(src) == 0 {
		return dest
	} else if len(src) == 1 {
		return appendOne(dest, src[0])
	} else {
		return AppendURL(appendOne(dest, src[0]), src[1:]...)
	}
}

func (arr URLs) contains( str URL) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func BeesToURLs(arr []*Bee) URLs {
	urls := make(URLs, len(arr))
	for i, bee := range arr {
		urls[i] = bee.URL()
	}
	return urls
}
func ArrToURLs(arr []string) URLs {
	urls := make(URLs, len(arr))
	for i, url := range arr {
		urls[i] = URL(url)
	}
	return urls
}

func explode(id string) (author, name string) {
	arr := strings.Split(string(id), "/")
	if len(arr) == 2 {
		return arr[0], arr[1]
	}
	return "error", "error"
}

func (b *Bee) URL() URL {
	return implode(b.Author, b.Name)
}

func (b *Bee) DepURLs() URLs {
	urls := make(URLs, len(b.Deps))
	for i, url := range b.Deps {
		urls[i] = URL(url)
	}
	return urls
}

func (h *Hive) DepURLs() URLs {
	urls := make(URLs, len(h.Deps))
	i := 0
	for url, _ := range h.Deps {
		urls[i] = URL(url)
	}
	return urls
}

func (u URL) ToBee() (Bee, error) {
	b := Bee{}
	return  b, nil
}

