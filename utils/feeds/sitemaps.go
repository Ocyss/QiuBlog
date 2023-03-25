package feeds

import (
	"encoding/xml"
	"time"
)

type SitemapFeedXml struct {
	XMLName          xml.Name `xml:"urlset"`
	ContentNamespace string   `xml:"xmlns,attr"`
	Channel          []*SitemapFeed
}

type SitemapBaidu struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type SitemapFeed struct {
	XMLName    xml.Name     `xml:"url"`
	Loc        string       `xml:"loc"`
	Lastmod    string       `xml:"lastmod,omitempty"`
	Changefreq string       `xml:"changefreq,omitempty"`
	Priority   string       `xml:"priority,omitempty"`
	Baidu      SitemapBaidu `xml:"mobile:mobile"`
}

type Sitemap struct {
	*Feed
}

func (r *Sitemap) SitemapFeed() []*SitemapFeed {
	var channel []*SitemapFeed
	for _, v := range r.Items {
		pub := anyTimeFormat(time.RFC3339, v.Updated, v.Created)
		channel = append(channel, &SitemapFeed{
			Loc:        v.Link.Href,
			Lastmod:    pub,
			Changefreq: v.Changefreq,
			Priority:   v.Priority,
			Baidu:      SitemapBaidu{Type: "pc,mobile"},
		})
	}
	return channel
}

func (r *Sitemap) FeedXml() interface{} {
	return &SitemapFeedXml{
		Channel:          r.SitemapFeed(),
		ContentNamespace: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
}
