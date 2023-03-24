package sitemap

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"qiublog/model"
	"qiublog/utils"
	"qiublog/utils/feeds"
	"strconv"
	"time"
)

var Feed *feeds.Feed
var author *feeds.Author

func Db() {
	// 先清空之前数据
	Feed.Items = nil
	articles := model.GetAllArticle()
	Feed.Items = make([]*feeds.Item, 0, len(articles)+100)

	for _, v := range articles {
		Feed.Items = append(Feed.Items, &feeds.Item{
			Title:       v.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("%s/post/%d", utils.Config.Server.Url, v.ID)},
			Author:      author,
			Description: v.Desc,
			Id:          strconv.Itoa(int(v.ID)),
			Updated:     v.UpdatedAt,
			Created:     v.CreatedAt,
		})
	}
	file, err := os.OpenFile("./sitemap.cache", os.O_CREATE|os.O_WRONLY, 0200)
	if err != nil {
		fmt.Println("sitemap缓存写入失败!", err)
		return
	}
	data, _ := json.Marshal(&Feed.Items)
	file.Write(data)
	file.Close()
}

func InitSitemap() {
	// 获取博客建站时间
	CreatedTime := time.Unix(utils.Config.ConstructionTime, 0)
	author = &feeds.Author{Name: utils.Config.Server.AuthorName, Email: utils.Config.Server.AuthorEmail}
	Feed = &feeds.Feed{
		Title:       "Ocyss 的博客",
		Link:        &feeds.Link{Href: "https://邱.cf/"},
		Description: "故事很短，满是遗憾。",
		Author:      author,
		Created:     CreatedTime,
	}
	// 先读取本地的缓存，没有在遍历数据库
	sitemapCache, err := os.Open("./sitemap.cache")
	if err != nil {
		Db()
	} else {
		defer sitemapCache.Close()
		sitemapByte, err := io.ReadAll(sitemapCache)
		if err != nil {
			Db()
		}
		err = json.Unmarshal(sitemapByte, &Feed.Items)
		if err != nil {
			Db()
		}
	}

}
