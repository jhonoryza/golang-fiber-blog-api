package web_controllers

import (
	"fiber_blog/app/models"
	"fiber_blog/app/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/jhonoryza/inertia-fiber"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Column struct {
	Key      string `json:"key"`
	Label    string `json:"label"`
	Visible  bool   `json:"visible"`
	Sortable bool   `json:"sortable"`
}

type Paginate struct {
	Data        *[]responses.PostResponses `json:"data"`
	From        int                        `json:"from"`
	To          int                        `json:"to"`
	NextPageUrl string                     `json:"next_page_url"`
	PrevPageUrl string                     `json:"prev_page_url"`
}

func PostIndex(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		filterSearch := c.Query("filter[search]")
		filterTitle := c.Query("filter[title]")

		sort := c.Query("sort")
		sortKey := "id"
		sortDir := "desc"
		splitSort := strings.Split(sort, "-")
		if sort != "" && len(splitSort) > 1 {
			sortKey = splitSort[0]
			sortDir = splitSort[1]
		}

		limitStr := c.Query("limit", "10")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			panic(err)
		}
		if limit == 0 {
			limit = 10
		}
		pageStr := c.Query("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			panic(err)
		}
		if page == 0 {
			page = 1
		}
		offset := (page - 1) * limit
		columns := []Column{
			{Key: "id", Label: "Id", Visible: true, Sortable: true},
			{Key: "title", Label: "Title", Visible: true, Sortable: true},
			{Key: "published_at", Label: "Published", Visible: true, Sortable: true},
		}

		var posts []models.Post

		query := db
		if filterSearch != "" {
			query = query.Where("lower(title) LIKE ?", "%"+strings.ToLower(filterSearch)+"%")
		}
		if filterTitle != "" {
			query = query.Where("lower(title) = ?", strings.ToLower(filterTitle))
		}
		if sort != "" {
			query = query.Order(sortKey + " " + sortDir)
		}
		query.Where("published_at is not null").Offset(offset).Limit(limit).Find(&posts)

		postResponses := responses.NewPostResponses(&posts)

		//baseUrl := env.GetEnv().GetString("APP_URL")
		queryParams := c.Request().URI().QueryArgs()

		nextPage := page + 1
		nextPageUrl := buildPaginationURL(c, nextPage, limit, queryParams)
		if len(*postResponses) != limit {
			nextPageUrl = ""
		}

		prevPage := page - 1
		prevPageUrl := buildPaginationURL(c, prevPage, limit, queryParams)
		if prevPage <= 0 {
			prevPageUrl = ""
		}

		from := (page-1)*limit + 1
		to := min(from+limit-1, len(*postResponses))

		paginate := Paginate{
			Data:        postResponses,
			From:        from,
			To:          to,
			NextPageUrl: nextPageUrl,
			PrevPageUrl: prevPageUrl,
		}

		return inertia.Render(c, http.StatusOK, "Admin/Post/Index", fiber.Map{
			"posts":       paginate,
			"pageOptions": []int{5, 10, 25, 50, 100},
			"limit":       limit,
			"columns":     columns,
			"allIds":      []int{},
			"filters":     []string{"search", "title"},
			"defaultSort": "-id",
		})
	}
}

func buildPaginationURL(c *fiber.Ctx, page int, limit int, queryParams *fasthttp.Args) string {
	baseURL := c.BaseURL() + c.Path()
	query := url.Values{}

	queryParams.VisitAll(func(key, value []byte) {
		if string(key) != "page" || string(key) != "limit" {
			query.Add(string(key), string(value))
		}
	})

	query.Set("page", strconv.Itoa(page))
	query.Set("limit", strconv.Itoa(limit))

	return baseURL + "?" + query.Encode()
}
