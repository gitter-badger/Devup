package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

type User struct {
	Login             *string `json:"login,omitempty"`
	ID                *int64  `json:"id,omitempty"`
	AvatarURL         *string `json:"avatar_url,omitempty"`
	HTMLURL           *string `json:"html_url,omitempty"`
	GravatarID        *string `json:"gravatar_id,omitempty"`
	Type              *string `json:"type,omitempty"`
	SiteAdmin         *bool   `json:"site_admin,omitempty"`
	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
}

type Permission struct {
	Admin *bool `json:"admin,omitempty"`
	Push  *bool `json:"push,omitempty"`
	Pull  *bool `json:"pull,omitempty"`
}

type Repository struct {
	ID               *int64 `json:"id,omitempty"`
	Owner            *User
	Name             *string `json:"name,omitempty"`
	FullName         *string `json:"full_name,omitempty"`
	Description      *string `json:"description,omitempty"`
	Homepage         *string `json:"homepage,omitempty"`
	DefaultBranch    *string `json:"default_branch,omitempty"`
	MasterBranch     *string `json:"master_branch,omitempty"`
	CreatedAt        *string `json:"created_at,omitempty"`
	PushedAt         *string `json:"pushed_at,omitempty"`
	UpdatedAt        *string `json:"updated_at,omitempty"`
	HTMLURL          *string `json:"html_url,omitempty"`
	CloneURL         *string `json:"clone_url,omitempty"`
	GitURL           *string `json:"git_url,omitempty"`
	SSHURL           *string `json:"ssh_url,omitempty"`
	SVNURL           *string `json:"svn_url,omitempty"`
	Language         *string `json:"language,omitempty"`
	Fork             *bool   `json:"fork,omitempty"`
	ForksCount       *int64  `json:"forks_count,omitempty"`
	OpenIssuesCount  *int64  `json:"open_issues_count,omitempty"`
	StargazersCount  *int64  `json:"stargazers_count,omitempty"`
	WatchersCount    *int64  `json:"watchers_count,omitempty"`
	Size             *int64  `json:"size,omitempty"`
	Permissions      *Permission
	Private          *bool   `json:"private,omitempty"`
	HasIssues        *bool   `json:"has_issues,omitempty"`
	HasWiki          *bool   `json:"has_wiki,omitempty"`
	HasDownloads     *bool   `json:"has_downloads,omitempty"`
	URL              *string `json:"url,omitempty"`
	ArchiveURL       *string `json:"archive_url,omitempty"`
	AssigneesURL     *string `json:"assignees_url,omitempty"`
	BlobsURL         *string `json:"blobs_url,omitempty"`
	BranchesURL      *string `json:"branches_url,omitempty"`
	CollaboratorsURL *string `json:"collaborators_url,omitempty"`
	CommentsURL      *string `json:"comments_url,omitempty"`
	CommitsURL       *string `json:"commits_url,omitempty"`
	CompareURL       *string `json:"compare_url,omitempty"`
	ContentsURL      *string `json:"contents_url,omitempty"`
	ContributorsURL  *string `json:"contributors_url,omitempty"`
	DownloadsURL     *string `json:"downloads_url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	ForksURL         *string `json:"forks_url,omitempty"`
	GitCommitsURL    *string `json:"git_commits_url,omitempty"`
	GitRefsURL       *string `json:"git_refs_url,omitempty"`
	GitTagsURL       *string `json:"git_tags_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssueCommentURL  *string `json:"issue_comment_url,omitempty"`
	IssueEventsURL   *string `json:"issue_events_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	KeysURL          *string `json:"keys_url,omitempty"`
	LabelsURL        *string `json:"labels_url,omitempty"`
	LanguagesURL     *string `json:"languages_url,omitempty"`
	MergesURL        *string `json:"merges_url,omitempty"`
	MilestonesURL    *string `json:"milestones_url,omitempty"`
	NotificationsURL *string `json:"notifications_url,omitempty"`
	PullsURL         *string `json:"pulls_url,omitempty"`
	ReleasesURL      *string `json:"releases_url,omitempty"`
	StargazersURL    *string `json:"stargazers_url,omitempty"`
	StatusesURL      *string `json:"statuses_url,omitempty"`
	SubscribersURL   *string `json:"subscribers_url,omitempty"`
	SubscriptionURL  *string `json:"subscription_url,omitempty"`
	TagsURL          *string `json:"tags_url,omitempty"`
	TreesURL         *string `json:"trees_url,omitempty"`
	TeamsURL         *string `json:"teams_url,omitempty"`
}

type RepositoryComment struct {
	HTMLURL   *string `json:"html_url,omitempty"`
	URL       *string `json:"url,omitempty"`
	ID        *int    `json:"id,omitempty"`
	CommitID  *string `json:"commit_id,omitempty"`
	User      *User
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Line      *string `json:"line, omitempty"`
	Body      *string `json:"body"`
	Path      *string `json:"path,omitempty"`
	Position  *int    `json:"position,omitempty"`
}

type RepositoryCommit struct {
	SHA       *string      `json:"sha,omitempty"`
	Commit    *Commit      `json:"commit,omitempty"`
	Author    *User        `json:"author,omitempty"`
	Committer *User        `json:"committer,omitempty"`
	Parents   []Commit     `json:"parents,omitempty"`
	Message   *string      `json:"message,omitempty"`
	HTMLURL   *string      `json:"html_url,omitempty"`
	Stats     *CommitStats `json:"stats,omitempty"`
	Files     []CommitFile `json:"files,omitempty"`
}

type Commit struct {
	SHA       *string       `json:"sha,omitempty"`
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
	Message   *string       `json:"message,omitempty"`
	/*Tree         *Tree       `json:"tree,omitempty"`*/
	Parents      []Commit     `json:"parents,omitempty"`
	Stats        *CommitStats `json:"stats,omitempty"`
	URL          *string      `json:"url,omitempty"`
	CommentCount *int         `json:"comment_count,omitempty"`
}

type CommitAuthor struct {
	Date  *string `json:"date,omitempty"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type CommitStats struct {
	Additions *int `json:"additions,omitempty"`
	Deletions *int `json:"deletions,omitempty"`
	Total     *int `json:"total,omitempty"`
}

type CommitFile struct {
	SHA       *string `json:"sha,omitempty"`
	Filename  *string `json:"filename,omitempty"`
	Additions *int    `json:"additions,omitempty"`
	Deletions *int    `json:"deletions,omitempty"`
	Changes   *int    `json:"changes,omitempty"`
	Status    *string `json:"status,omitempty"`
	Patch     *string `json:"patch,omitempty"`
}

func serveRest(w http.ResponseWriter, r *http.Request) {

	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))
}

func serveRest1(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/login/oauth/authorize?client_id=c5376f36b92a55ac20e1", 301)
}

func serveRest2(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	res, err := http.Get("https://github.com/login/oauth/access_token?client_id=c5376f36b92a55ac20e1&client_secret=a2933fb99800ed7b683b70a139f2691d5e2953e8&code=" + code + "")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	url, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	//fmt.Println(string(url))

	s := strings.Split(string(url), "&")

	//fmt.Println(s)

	at := s[0]

	//fmt.Println(at)

	access_token := strings.Split(at, "=")

	//fmt.Println(access_token[1])

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token[1]},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List("", nil)

	if err != nil {
		panic(err)
	}

	repo, err := json.Marshal(repos)

	if err != nil {
		panic(err)
	}

	var dat []Repository
	var d1 Repository

	if err := json.Unmarshal([]byte(string(repo)), &dat); err != nil {
		fmt.Println(err)
	}

	m := make(map[string]string)
	n := make(map[string]string)
	m1 := make(map[string]string)

	for key := range dat {
		dat1, _ := json.Marshal(dat[key])

		if err := json.Unmarshal([]byte(string(dat1)), &d1); err != nil {
			fmt.Println(err)
		}

		id1, _ := json.Marshal(d1.Name)
		userid, _ := json.Marshal(d1.Owner.Login)

		t1, err := strconv.Unquote(string(id1))
		t2, err := strconv.Unquote(string(userid))

		// list all comments of all repos for the authenticated user
		comments, _, err := client.Repositories.ListComments(t2, t1, nil)

		if err != nil {
			panic(err)
		}

		commits, _, err := client.Repositories.ListCommits(t2, t1, nil)

		if err != nil {
			panic(err)
		}

		comment, err := json.Marshal(comments)

		if err != nil {
			fmt.Println(err)
		}

		commit, err := json.Marshal(commits)

		if err != nil {
			fmt.Println(err)
		}

		var com []RepositoryCommit
		var cm1 RepositoryCommit

		if err := json.Unmarshal([]byte(string(commit)), &com); err != nil {
			fmt.Println(err)
		}

		for key := range com {
			com1, _ := json.Marshal(com[key])

			if err := json.Unmarshal([]byte(string(com1)), &cm1); err != nil {
				fmt.Println(err)
			}

			sha, _ := json.Marshal(cm1.SHA)

			t3, err := strconv.Unquote(string(sha))

			comments1, _, err := client.Repositories.ListCommitComments(t2, t1, t3, nil)

			if err != nil {
				panic(err)
			}

			comment1, err := json.Marshal(comments1)
			m1[string(sha)] = string(comment1)

		}

		m[string(id1)] = string(comment)
		n[string(id1)] = string(commit)

	}

	/*	for k, v := range m {
		fmt.Printf("comments of %s -> %s\n", k, v)
	}*/
	for k, v := range m1 {
		fmt.Printf("commit of %s -> %s\n", k, v)
	}

	//fmt.Println(string(repo))
	fmt.Fprintf(w, string(repo))

}

func main() {

	http.HandleFunc("/", serveRest)
	http.HandleFunc("/github", serveRest1)
	http.HandleFunc("/github/callback", serveRest2)
	http.ListenAndServe("localhost:1337", nil)

}

func getJsonResponse() ([]byte, error) {

	fruits := make(map[string]int)
	fruits["Banana"] = 23
	fruits["Apple"] = 4

	vegetables := make(map[string]int)
	vegetables["Carrets"] = 44
	vegetables["Goolse"] = 7

	d := Data{fruits, vegetables}
	p := Payload{d}

	//fmt.Println(p);

	return json.MarshalIndent(p, "", " ")

}
