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

type CreateAccountParameters struct {
	ShortName  string `json:"short_name"`
	AuthorName string `json:"author_name,omitempty"`
	AuthorURL  string `json:"author_url,omitempty"`
}

type EditAccountInfoParameters struct {
	AccessToken string `json:"access_token"`
	ShortName   string `json:"short_name,omitempty"`
	AuthorName  string `json:"author_name,omitempty"`
	AuthorURL   string `json:"author_url,omitempty"`
}

type GetAccountInfoParameters struct {
	AccessToken string              `json:"access_token"`
	Fields      []AccountInfoFields `json:"fields,omitempty"`
}

type RevokeAccessTokenParameters struct {
	AccessToken string `json:"access_token"`
}

const (
	AccountShortName  AccountInfoFields = "short_name"
	AccountAuthorName AccountInfoFields = "author_name"
	AccountAuthorURL  AccountInfoFields = "author_url"
	AccountAuthURL    AccountInfoFields = "auth_url"
	AccountPageCount  AccountInfoFields = "page_count"
)

func (c *client) CreateAccount(p *CreateAccountParameters) (result *Account, err error) {
	var (
		response *Response
	)

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

func (c *client) EditAccountInfo(p *EditAccountInfoParameters) (result *Account, err error) {
	var (
		parameters interface{}
		response   *Response
	)

	if len(p.AccessToken) == 0 {
		parameters = c.withToken(p)
	} else {
		parameters = p
	}

	response, err = c.request("/editAccountInfo", parameters)
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

func (c *client) GetAccountInfo(p *GetAccountInfoParameters) (result *Account, err error) {
	var (
		parameters interface{}
		response   *Response
	)

	if len(p.AccessToken) == 0 {
		parameters = c.withToken(p)
	} else {
		parameters = p
	}

	response, err = c.request("/getAccountInfo", parameters)
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

func (c *client) RevokeAccessToken(p *RevokeAccessTokenParameters) (result *Account, err error) {
	var (
		parameters interface{}
		response   *Response
	)

	if len(p.AccessToken) == 0 {
		parameters = c.withToken(p)
	} else {
		parameters = p
	}

	response, err = c.request("/revokeAccessToken", parameters)
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
