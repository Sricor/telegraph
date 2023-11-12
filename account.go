package telegraph

type Account struct {
	ShortName   string `json:"short_name"`
	AuthorName  string `json:"author_name"`
	AuthorURL   string `json:"author_url"`
	AccessToken string `json:"access_token"`
	AuthURL     string `json:"auth_url"`
	PageCount   int    `json:"page_count"`
}

type AccountInfoFields string

type CreateAccountRequest struct {
	ShortName  string `json:"short_name"`
	AuthorName string `json:"author_name,omitempty"`
	AuthorURL  string `json:"author_url,omitempty"`
}

type EditAccountInfoRequest struct {
	AccessToken string `json:"access_token"`
	ShortName   string `json:"short_name,omitempty"`
	AuthorName  string `json:"author_name,omitempty"`
	AuthorURL   string `json:"author_url,omitempty"`
}

type GetAccountInfoRequest struct {
	AccessToken string              `json:"access_token"`
	Fields      []AccountInfoFields `json:"fields,omitempty"`
}

type RevokeAccessTokenRequest struct {
	AccessToken string `json:"access_token"`
}

const (
	AccountFieldShortName  AccountInfoFields = "short_name"
	AccountFieldAuthorName AccountInfoFields = "author_name"
	AccountFieldAuthorURL  AccountInfoFields = "author_url"
	AccountFieldAuthURL    AccountInfoFields = "auth_url"
	AccountFieldPageCount  AccountInfoFields = "page_count"
)

func (c *client) CreateAccount(p *CreateAccountRequest) (result *Account, err error) {
	var response *Response

	response, err = c.request("/createAccount", p)
	if err != nil {
		return
	}

	result = &Account{}
	err = jsonUnmarshal(response.Result, result)
	if err != nil {
		return
	}

	return
}

func (c *client) EditAccountInfo(p *EditAccountInfoRequest) (result *Account, err error) {
	var response *Response

	if len(p.AccessToken) == 0 {
		p.AccessToken = c.token
		defer func() {
			p.AccessToken = ""
		}()
	}

	response, err = c.request("/editAccountInfo", p)
	if err != nil {
		return
	}

	result = &Account{}
	err = jsonUnmarshal(response.Result, result)
	if err != nil {
		return
	}

	return
}

func (c *client) GetAccountInfo(p *GetAccountInfoRequest) (result *Account, err error) {
	var response *Response

	if len(p.AccessToken) == 0 {
		p.AccessToken = c.token
		defer func() {
			p.AccessToken = ""
		}()
	}

	response, err = c.request("/getAccountInfo", p)
	if err != nil {
		return
	}

	result = &Account{}
	err = jsonUnmarshal(response.Result, result)
	if err != nil {
		return
	}

	return
}

func (c *client) RevokeAccessToken(p *RevokeAccessTokenRequest) (result *Account, err error) {
	var response *Response

	if len(p.AccessToken) == 0 {
		p.AccessToken = c.token
		defer func() {
			p.AccessToken = ""
		}()
	}

	response, err = c.request("/revokeAccessToken", p)
	if err != nil {
		return
	}

	result = &Account{}
	err = jsonUnmarshal(response.Result, result)
	if err != nil {
		return
	}

	return
}
