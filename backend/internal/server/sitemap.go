package server

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/query"
)

var baseURL = "https://knnkt.dk"

var baseURLs = []string{
	"/",
	"/about",
	"/events",
	"/artists",
}

type sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	XMLNs   string       `xml:"xmlns,attr"`
	URL     []SitemapURL `xml:"url"`
}

type SitemapURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
	LastMod string   `xml:"lastmod"`
}

func (s Server) handleGetSitemap() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var sitemap sitemap

		sitemap.XMLNs = "http://www.sitemaps.org/schemas/sitemap/0.9"

		eventFilters := query.FilterCollection{
			"is_public": []query.Filter{{Cmp: query.Equal, Value: "true"}},
		}

		eventQuery, err := query.NewListQuery(
			query.WithFilters(eventFilters),
		)
		if err != nil {
			writeError(w, err)
			return
		}

		eventResult, err := s.eventService.List(ctx, eventQuery)
		if err != nil {
			writeError(w, err)
			return
		}

		artists := make(map[int64]artist.Artist)
		for _, e := range eventResult.Records {
			for _, c := range e.Concerts {
				artists[c.Artist.ID] = c.Artist
			}
		}

		urls := baseURLs
		for _, event := range eventResult.Records {
			urls = append(urls, fmt.Sprintf("/events/%d", event.ID))
		}

		for _, artist := range artists {
			urls = append(urls, fmt.Sprintf("/artists/%d", artist.ID))
		}

		y, m, _ := time.Now().Date()
		lastmod := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")

		for _, url := range urls {
			u := SitemapURL{
				Loc:     baseURL + url,
				LastMod: lastmod,
			}

			sitemap.URL = append(sitemap.URL, u)
		}

		outXML, _ := xml.MarshalIndent(sitemap, "", "  ")
		sitemapXML := append([]byte(xml.Header), outXML...)
		w.Header().Set("Content-Type", "application/xml")
		_, err = w.Write(sitemapXML)
		if err != nil {
			writeError(w, err)
			return
		}
	}
}
