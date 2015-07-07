package internal

import (
	"log"
	"net/http"
)

type CookieContainer struct {
	cookies []*http.Cookie
}

func (container *CookieContainer) Merge(newCookies []*http.Cookie) {
	for i,override := range newCookies {
		found := false
		for _,existing := range container.cookies {
			if existing.Name != override.Name {
				continue
			}
			container.cookies[i] = override
			found = true
		}
		if !found {
			container.cookies = append(container.cookies, override)
		}
	}
}

func (container *CookieContainer) AddToRequest(req *http.Request) {
	for _,cookie := range container.cookies {
		req.AddCookie(cookie)
	}
}

func (container *CookieContainer) Get(name string) *http.Cookie {
	for _,cookie := range container.cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

func (container *CookieContainer) Log() {
	for _,cookie := range container.cookies {
		log.Println(cookie.Name, "=", cookie.Value)
	}
}