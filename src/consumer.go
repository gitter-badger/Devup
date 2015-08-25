
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
/*{
        "url": "https://api.github.com/repos/smahawar/myApp/comments/12686327",
        "html_url": "https://github.com/smahawar/myApp/commit/a086e7ff7a6dc53ff5cbc1f5ed0e600d2243865a#commitcomment-12686327",
        "id": 12686327,
        "user": {
            "login": "smahawar",
            "id": 9653412,
            "avatar_url": "https://avatars.githubusercontent.com/u/9653412?v=3",
            "gravatar_id": "",
            "url": "https://api.github.com/users/smahawar",
            "html_url": "https://github.com/smahawar",
            "followers_url": "https://api.github.com/users/smahawar/followers",
            "following_url": "https://api.github.com/users/smahawar/following{/other_user}",
            "gists_url": "https://api.github.com/users/smahawar/gists{/gist_id}",
            "starred_url": "https://api.github.com/users/smahawar/starred{/owner}{/repo}",
            "subscriptions_url": "https://api.github.com/users/smahawar/subscriptions",
            "organizations_url": "https://api.github.com/users/smahawar/orgs",
            "repos_url": "https://api.github.com/users/smahawar/repos",
            "events_url": "https://api.github.com/users/smahawar/events{/privacy}",
            "received_events_url": "https://api.github.com/users/smahawar/received_events",
            "type": "User",
            "site_admin": false
        },
        "position": null,
        "line": null,
        "path": "",
        "commit_id": "a086e7ff7a6dc53ff5cbc1f5ed0e600d2243865a",
        "created_at": "2015-08-13T10:44:17Z",
        "updated_at": "2015-08-13T10:44:17Z",
        "body": "This is testing comment. ignore it."
    },*/

type Users struct{
	Login string "json:login"
	Id int "json:id"
	Avatar_url string "json:avatar_url"
	Gravatar_id string "json:gravatar_id"
	Url string "json:url"
	Html_url string "json:html_url"
	Followers_url string "json:followers_url"
	Followers_url string "json:followers_url"
	Gists_url string "json:gists_url"
	Starred_url string "json:starred_url"
	Subscriptions_url string "json:subscriptions_url"
	Organizations_url string "json:organizations_url"
	Repos_url string "json:repos_url"
	Events_url string "json:events_url"
	Received_events_url "json:received_events_url"
	Type string "json:type"
	Site_admin bool "json:site_admin"
}

type Comments struct{
	Url string "json:url"
	Html_url string "json:html_url"
	Id int "json:id"
	User Users
	Position string "json:position"
	Line string "json:line"
	Path string "json:path"
	Commit_id string "json:commit_id"
	Created_at string "json:created_at"
	Updated_at string "json:updated_at"
	Body string "json:body"
}

func Decode(r io.Reader) (x *Comments, err error) {
    x = new(Comments)
    if err = json.NewDecoder(r).Decode(x); err != nil {
        return
    }
    err = json.Unmarshal(x.AuthorRaw, &x.Author); err == nil {
        return
    }
    var s string
    err = json.Unmarshal(x.AuthorRaw, &s); err == nil {
        x.Author.Email = s
        return
    }
    var n uint64
    err = json.Unmarshal(x.AuthorRaw, &n); err == nil {
        x.Author.ID = n
    }
    return
}


func main() {

	url := "https://api.github.com/repos/smahawar/myApp/commits/a086e7ff7a6dc53ff5cbc1f5ed0e600d2243865a/comments"

	res, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil{
		panic(err)
	}
	p := Comments{}
	p Complex128;	err = json.Unmarshal(body, &p)

	if err != nil{
		panic(err)
	}

	fmt.Println(p)
}
