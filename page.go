package telegraph

type Node interface{}

type NodeElement struct {
	// Name of the DOM element. Available tags: a, aside, b, blockquote, br, code, em, figcaption, figure, h3, h4, hr, i, iframe, img, li, ol, p, pre, s, strong, u, ul, video.
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

type GetPageRequest struct {
	Path          string `json:"path"`
	ReturnContent bool   `json:"return_content,omitempty"`
}

type GetViewsRequest struct {
	Path  string `json:"path"`
	Year  int    `json:"year,omitempty"`
	Month int    `json:"month,omitempty"`
	Day   int    `json:"day,omitempty"`
	Hour  int    `json:"hour,omitempty"`
}

type CreatePageRequest struct {
	AccessToken   string `json:"access_token"`
	Title         string `json:"title"`
	AuthorName    string `json:"author_name,omitempty"`
	AuthorURL     string `json:"author_url,omitempty"`
	Content       []Node `json:"content"`
	ReturnContent bool   `json:"return_content,omitempty"`
}

type EditPageRequest struct {
	AccessToken   string `json:"access_token"`
	Path          string `json:"path"`
	Title         string `json:"title"`
	Content       []Node `json:"content"`
	AuthorName    string `json:"author_name,omitempty"`
	AuthorURL     string `json:"author_url,omitempty"`
	ReturnContent bool   `json:"return_content,omitempty"`
}

type GetPageListRequest struct {
	AccessToken string `json:"access_token"`
	Offset      int    `json:"offset,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}

func (c *client) GetPage(p *GetPageRequest) (result *Page, err error) {
	var response *Response

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

func (c *client) GetViews(p *GetViewsRequest) (result *PageViews, err error) {
	var response *Response

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

func (c *client) CreatePage(p *CreatePageRequest) (result *Page, err error) {
	var response *Response

	if len(p.AccessToken) == 0 {
		p.AccessToken = c.token
		defer func() {
			p.AccessToken = ""
		}()
	}

	response, err = c.request("/createPage", p)
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

func (c *client) EditPage(p *EditPageRequest) (result *Page, err error) {
	var response *Response

	if len(p.AccessToken) == 0 {
		p.AccessToken = c.token
		defer func() {
			p.AccessToken = ""
		}()
	}

	response, err = c.request("/editPage", p)
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

func (c *client) GetPageList(p *GetPageListRequest) (result *PageList, err error) {
	var response *Response

	if len(p.AccessToken) == 0 {
		p.AccessToken = c.token
		defer func() {
			p.AccessToken = ""
		}()
	}

	response, err = c.request("/getPageList", p)
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

func (n *NodeElement) AddChildren(t *NodeElement) *NodeElement {
	n.Children = append(n.Children, t)
	return t
}

func (n *NodeElement) AddText(text string) *NodeElement {
	return n.AddChildren(&NodeElement{
		Tag:      "p",
		Children: []Node{text},
	})
}

func (n *NodeElement) AddImage(src string, description string) *NodeElement {
	image := NodeElement{
		Tag: "figure",
	}
	image.AddChildren(&NodeElement{
		Tag:   "img",
		Attrs: map[string]string{"src": src},
	})
	image.AddChildren(&NodeElement{
		Tag:      "figcaption",
		Children: []Node{description},
	})
	return n.AddChildren(&image)
}
