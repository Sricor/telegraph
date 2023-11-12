# Telegraph

Telegra.ph API

## Installation
`go get github.com/sricor/telegraph`

## Usage
<h3 id="createAccount">Account > <a href="https://telegra.ph/api#createAccount"><code>createAccount</code></a></h3>
<details>
<summary>Creates Telegraph Account</summary>

```go
package main

import (
	"fmt"

	"github.com/sricor/telegraph"
)

func main() {
	client := telegraph.NewClient()
	account, err := client.CreateAccount(&telegraph.CreateAccountRequest{
		ShortName: "Sandbox",
		AuthorName: "Anonymous",
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(account)
}
```

</details>




<h3 id="editAccount">Account > <a href="https://telegra.ph/api#editAccountInfo"><code>editAccount</code></a></h3>
<details>
<summary>Edit Telegraph Account</summary>

```go
package main

import (
	"fmt"

	"github.com/sricor/telegraph"
)

func main() {
	client := telegraph.NewClient()
	client.SetToken("<Token>")

	account, err := client.EditAccountInfo(&telegraph.EditAccountInfoRequest{
		ShortName: "Sandbox2",
		AuthorName: "Anonymous2",
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(account)
}
```

</details>



<h3 id="getAccount">Account > <a href="https://telegra.ph/api#getAccountInfo"><code>getAccount</code></a></h3>
<details>
<summary>Get Telegraph Account Information</summary>

```go
package main

import (
	"fmt"

	"github.com/sricor/telegraph"
)

func main() {
	client := telegraph.NewClient()
	client.SetToken("<Token>")
	
	account, err := client.GetAccountInfo(&telegraph.GetAccountInfoRequest{
		Fields: []telegraph.AccountInfoFields{
			telegraph.AccountFieldShortName,
			telegraph.AccountFieldAuthorName,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(account)
}
```

</details>




<h3 id="revokeAccessToken">Account > <a href="https://telegra.ph/api#revokeAccessToken"><code>revokeAccessToken</code></a></h3>
<details>
<summary>Revoke Telegraph Account AccessToken</summary>

```go
package main

import (
	"fmt"

	"github.com/sricor/telegraph"
)

func main() {
	client := telegraph.NewClient()
	client.SetToken("<Token>")
    
	account, err := client.RevokeAccessToken(&telegraph.RevokeAccessTokenRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Print(account)
}
```

</details>




<h3 id="create">Page > <a href="https://telegra.ph/api#createPage"><code>createPage</code></a></h3>
<details>
<summary>Create Telegraph Page</summary>

```go
package main

import (
	"fmt"

	"github.com/sricor/telegraph"
)

func main() {
	client := telegraph.NewClient()
	client.SetToken("<Token>")

	content := telegraph.NodeElement{}
	content.AddText("Hello World!")

	page, err := client.CreatePage(&telegraph.CreatePageRequest{
		Title: "Sample Page",
		AuthorName: "Anonymous",
		Content: []telegraph.Node{content},
		// Content: []telegraph.Node{
		// 	telegraph.NodeElement{
		// 		Tag: "p",
		// 		Children: []telegraph.Node{
		// 			"Hello World!",
		// 		},
		// 	},
		// },
		ReturnContent: false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(page)
}
```

</details>



<h3 id="edit">Page > <a href="https://telegra.ph/api#editPage"><code>editPage</code></a></h3>
<details>
<summary>Edit Telegraph Page</summary>

```go
package main

import (
	"fmt"

	"github.com/sricor/telegraph"
)

func main() {
	client := telegraph.NewClient()
	client.SetToken("<Token>")

	content := telegraph.NodeElement{}
	content.AddText("Hello World!")

	page, err := client.EditPage(&telegraph.EditPageRequest{
		Title: "Sample Page",
		AuthorName: "Anonymous",
		Content: []telegraph.Node{content},
		// Content: []telegraph.Node{
		// 	telegraph.NodeElement{
		// 		Tag: "p",
		// 		Children: []telegraph.Node{
		// 			"Hello World!",
		// 		},
		// 	},
		// },
		ReturnContent: false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(page)
}
```

</details>



<h3 id="get">Page > <a href="https://telegra.ph/api#getPage"><code>getPage</code></a></h3>
<details>
<summary>Get Telegraph Page</summary>

```go
package main

import (
	"fmt"

	"github.com/sricor/telegraph"
)

func main() {
	client := telegraph.NewClient()

	page, err := client.GetPage(&telegraph.GetPageRequest{
		Path: "Sample-Page-12-15",
		ReturnContent: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(page)
}
```

</details>



<h3 id="getPages">Page > <a href="https://telegra.ph/api#getPageList"><code>getPageList</code></a></h3>
<details>
<summary>Get Telegraph Page List</summary>

```go
package main

import (
	"fmt"

	"github.com/sricor/telegraph"
)

func main() {
	client := telegraph.NewClient()
	client.SetToken("<Token>")

	pages, err := client.GetPageList(&telegraph.GetPageListRequest{
		Limit: 3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(pages)
}
```

</details>




<h3 id="getPageViews">Page > <a href="https://telegra.ph/api#getViews"><code>getPageViews</code></a></h3>
<details>
<summary>Get Telegraph Page Views</summary>

```go
package main

import (
	"fmt"

	"github.com/sricor/telegraph"
)

func main() {
	client := telegraph.NewClient()

	page, err := client.GetViews(&telegraph.GetViewsRequest{
		Path: "Sample-Page-12-15",
		Year: 2016,
		Month: 12,
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(page)
}
```

</details>
