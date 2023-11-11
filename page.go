package telegraph

type Node interface{}

type NodeElement struct {
	Tag      string            `json:"tag"`
	Attrs    map[string]string `json:"attrs,omitempty"`
	Children []Node            `json:"children,omitempty"`
}

type Page struct {
	Path        string `json:"path"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Views       int    `json:"views"`
	AuthorName  string `json:"author_name"`
	AuthorURL   string `json:"author_url"`
	ImageURL    string `json:"image_url"`
	Content     []Node `json:"content"`
	CanEdit     bool   `json:"can_edit"`
}

type PageList struct {
	TotalCount int    `json:"total_count"`
	Pages      []Page `json:"pages"`
}

type PageViews struct {
	Views int `json:"views"`
}

type CreatePageParameters struct {
	AccessToken   string `json:"access_token"`
	Title         string `json:"title"`
	AuthorName    string `json:"author_name,omitempty"`
	AuthorURL     string `json:"author_url,omitempty"`
	Content       []Node `json:"content"`
	ReturnContent bool   `json:"return_content,omitempty"`
}

type EditPageParameters struct {
	AccessToken   string `json:"access_token"`
	Path          string `json:"path"`
	Title         string `json:"title"`
	Content       []Node `json:"content"`
	AuthorName    string `json:"author_name,omitempty"`
	AuthorURL     string `json:"author_url,omitempty"`
	ReturnContent bool   `json:"return_content,omitempty"`
}

type GetPageParameters struct {
	Path          string `json:"path"`
	ReturnContent bool   `json:"return_content,omitempty"`
}

type GetPageListParameters struct {
	AccessToken string `json:"access_token"`
	Offset      int    `json:"offset,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}

type GetViewsParameters struct {
	Path  string `json:"path"`
	Year  int    `json:"year,omitempty"`
	Month int    `json:"month,omitempty"`
	Day   int    `json:"day,omitempty"`
	Hour  int    `json:"hour,omitempty"`
}

func (c *client) GetPage(p *GetPageParameters) (result *Page, err error) {
	var (
		response *Response
	)

	response, err = c.request("/getPage", p)
	if err != nil {
		return
	}

	result = &Page{}
	err = jsonUnmarshal(response.Result, result)
	if err != nil {
		return
	}

	return
}

func (c *client) GetViews(p *GetViewsParameters) (result *PageViews, err error) {
	var (
		response *Response
	)

	response, err = c.request("/getViews", p)
	if err != nil {
		return
	}

	result = &PageViews{}
	err = jsonUnmarshal(response.Result, result)
	if err != nil {
		return
	}

	return
}

func (c *client) CreatePage(p *CreatePageParameters) (result *Page, err error) {
	var (
		parameters interface{}
		response   *Response
	)

	if len(p.AccessToken) == 0 {
		parameters = c.withToken(p)
	} else {
		parameters = p
	}

	response, err = c.request("/createPage", parameters)
	if err != nil {
		return
	}

	result = &Page{}
	err = jsonUnmarshal(response.Result, result)
	if err != nil {
		return
	}

	return
}

func (c *client) EditPage(p *EditPageParameters) (result *Page, err error) {
	var (
		parameters interface{}
		response   *Response
	)

	if len(p.AccessToken) == 0 {
		parameters = c.withToken(p)
	} else {
		parameters = p
	}

	response, err = c.request("/editPage", parameters)
	if err != nil {
		return
	}

	result = &Page{}
	err = jsonUnmarshal(response.Result, result)
	if err != nil {
		return
	}

	return
}

func (c *client) GetPageList(p *GetPageListParameters) (result *PageList, err error) {
	var (
		parameters interface{}
		response   *Response
	)

	if len(p.AccessToken) == 0 {
		parameters = c.withToken(p)
	} else {
		parameters = p
	}

	response, err = c.request("/getPageList", parameters)
	if err != nil {
		return
	}

	result = &PageList{}
	err = jsonUnmarshal(response.Result, result)
	if err != nil {
		return
	}

	return
}
