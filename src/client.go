package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var store = sessions.NewCookieStore([]byte("github"))

type User struct {
	Login *string `json:"login,omitempty"`
	ID    *int64  `json:"id,omitempty"` /*
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
		SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`*/
}

type Permission struct {
	Admin *bool `json:"admin,omitempty"`
	Push  *bool `json:"push,omitempty"`
	Pull  *bool `json:"pull,omitempty"`
}

type Repository struct {
	ID       *int64 `json:"id,omitempty"`
	Owner    *User
	Name     *string `json:"name,omitempty"`
	FullName *string `json:"full_name,omitempty"`
	/*Description      *string `json:"description,omitempty"`
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
	TeamsURL         *string `json:"teams_url,omitempty"`*/
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
	SHA          *string       `json:"sha,omitempty"`
	Author       *CommitAuthor `json:"author,omitempty"`
	Committer    *CommitAuthor `json:"committer,omitempty"`
	Message      *string       `json:"message,omitempty"`
	Parents      []Commit      `json:"parents,omitempty"`
	Stats        *CommitStats  `json:"stats,omitempty"`
	URL          *string       `json:"url,omitempty"`
	CommentCount *int          `json:"comment_count,omitempty"`
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

type Comments struct {
	Body *string `json:"body"`
	User *User
}

func authGithub(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "https://github.com/login/oauth/authorize?client_id=c5376f36b92a55ac20e1", 301)
}

func authGithubCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	code := r.URL.Query().Get("code")

	res, err := http.Get("https://github.com/login/oauth/access_token?client_id=c5376f36b92a55ac20e1&client_secret=a2933fb99800ed7b683b70a139f2691d5e2953e8&code=" + code + "")

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	url, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	s := strings.Split(string(url), "&")

	at := s[0]

	access_token := strings.Split(at, "=")

	session, err := store.Get(r, "github")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	session.Values["access_token"] = access_token[1]

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token[1]},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	owner, _, err := client.Users.Get("")

	if err != nil {
		fmt.Println(err)
	}

	user, err := json.Marshal(owner)

	if err != nil {
		fmt.Println(err)
	}

	var userDetails User

	if err := json.Unmarshal([]byte(string(user)), &userDetails); err != nil {
		fmt.Println(err)
	}

	username, err := json.Marshal(userDetails.Login)

	session.Values["userid"] = string(username)

	session.Save(r, w)

	if access_token != nil && username != nil {
		fmt.Fprintf(w, "Successfully Logged in!")
	} else {
		fmt.Fprintf(w, "You are not authenticated!")
	}

}

func repos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session, err := store.Get(r, "github")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	at, ok := session.Values["access_token"]

	if !ok {
		fmt.Println(ok)
	}

	access_token, ok := at.(string)

	if !ok {
		fmt.Println(ok)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)

	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List("", nil)

	if err != nil {
		fmt.Println(err)
	}

	repo, err := json.Marshal(repos)

	if err != nil {
		fmt.Println(err)
	}

	var repoList []Repository

	if err := json.Unmarshal([]byte(string(repo)), &repoList); err != nil {
		fmt.Println(err)
	}

	/*	for key := range repoList {
		r1, _ := json.Marshal(repoList[key])

		if err := json.Unmarshal([]byte(string(r1)), &repo1); err != nil {
			fmt.Println(err)
		}

		id, _ := json.Marshal(repo1.ID)
		name, _ := json.Marshal(repo1.Name)
		userid, _ := json.Marshal(repo1.Owner.Login)

		fmt.Println(string(id), string(name), string(userid))
	}*/

	repoL, err := json.Marshal(repoList)

	fmt.Println(string(repoL))
	fmt.Fprintf(w, string(repoL))

}

func repoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	session, err := store.Get(r, "github")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	at, ok := session.Values["access_token"]

	if !ok {
		fmt.Println(ok)
	}

	access_token, ok := at.(string)

	if !ok {
		fmt.Println(ok)
	}

	uid, ok := session.Values["userid"]

	if !ok {
		fmt.Println(ok)
	}

	userid, ok := uid.(string)

	if !ok {
		fmt.Println(ok)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	t1, err := strconv.Unquote(string(userid))
	t2 := string(ps.ByName("repo"))

	comments, _, err := client.Repositories.ListComments(t1, t2, nil)

	if err != nil {
		fmt.Println(err)
	}

	comment, err := json.Marshal(comments)

	if err != nil {
		fmt.Println(err)
	}

	//var cmt []Comments
	var repoComments []Comments
	//var repoComment Comments

	if err := json.Unmarshal([]byte(string(comment)), &repoComments); err != nil {
		fmt.Println(err)
	}

	/*	for key := range repoComments {
		repoCmt, _ := json.Marshal(repoComments[key])

		if err := json.Unmarshal([]byte(string(repoCmt)), &repoComment); err != nil {
			fmt.Println(err)
		}
		body, _ := json.Marshal(repoComment.Body)
		fmt.Println(string(body))
	}*/

	cmt, err := json.Marshal(repoComments)

	fmt.Println(string(cmt))

	fmt.Fprintf(w, string(cmt))

}

func reposComments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session, err := store.Get(r, "github")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	at, ok := session.Values["access_token"]

	if !ok {
		fmt.Println(ok)
	}

	access_token, ok := at.(string)

	if !ok {
		fmt.Println(ok)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List("", nil)

	if err != nil {
		fmt.Println(err)
	}

	repo, err := json.Marshal(repos)

	if err != nil {
		fmt.Println(err)
	}

	var repoList []Repository
	var repo1 Repository

	if err := json.Unmarshal([]byte(string(repo)), &repoList); err != nil {
		fmt.Println(err)
	}

	map1 := make(map[string]string)

	for key := range repoList {
		r1, _ := json.Marshal(repoList[key])

		if err := json.Unmarshal([]byte(string(r1)), &repo1); err != nil {
			fmt.Println(err)
		}

		name, _ := json.Marshal(repo1.Name)
		userid, _ := json.Marshal(repo1.Owner.Login)

		t1, err := strconv.Unquote(string(userid))
		t2, err := strconv.Unquote(string(name))

		comments, _, err := client.Repositories.ListComments(t1, t2, nil)

		if err != nil {
			fmt.Println(err)
		}

		comment, err := json.Marshal(comments)

		if err != nil {
			fmt.Println(err)
		}

		map1[string(name)] = string(comment)
	}

	fmt.Println(map1)

	fmt.Fprintf(w, string(repo))

}

func repoCommitsComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	session, err := store.Get(r, "github")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	at, ok := session.Values["access_token"]

	if !ok {
		fmt.Println(ok)
	}

	access_token, ok := at.(string)

	if !ok {
		fmt.Println(ok)
	}

	uid, ok := session.Values["userid"]

	if !ok {
		fmt.Println(ok)
	}

	userid, ok := uid.(string)

	if !ok {
		fmt.Println(ok)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	t1, err := strconv.Unquote(string(userid))
	t2 := string(ps.ByName("repo"))

	commits, _, err := client.Repositories.ListCommits(t1, t2, nil)

	if err != nil {
		fmt.Println(err)
	}

	commit, err := json.Marshal(commits)

	if err != nil {
		fmt.Println(err)
	}

	var repoCommits []RepositoryCommit
	var commit1 RepositoryCommit

	if err := json.Unmarshal([]byte(string(commit)), &repoCommits); err != nil {
		fmt.Println(err)
	}

	for key := range repoCommits {
		repoCommit1, _ := json.Marshal(repoCommits[key])

		if err := json.Unmarshal([]byte(string(repoCommit1)), &commit1); err != nil {
			fmt.Println(err)
		}

		sha, _ := json.Marshal(commit1.SHA)

		t3, err := strconv.Unquote(string(sha))

		comments, _, err := client.Repositories.ListCommitComments(t1, t2, t3, nil)

		if err != nil {
			fmt.Println(err)
		}

		comment, err := json.Marshal(comments)
		fmt.Println(string(comment))
	}

	fmt.Fprintf(w, string(commit))

}

func repoCommits(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	session, err := store.Get(r, "github")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	at, ok := session.Values["access_token"]

	if !ok {
		fmt.Println(ok)
	}

	access_token, ok := at.(string)

	if !ok {
		fmt.Println(ok)
	}

	uid, ok := session.Values["userid"]

	if !ok {
		fmt.Println(ok)
	}

	userid, ok := uid.(string)

	if !ok {
		fmt.Println(ok)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	t1, err := strconv.Unquote(string(userid))
	t2 := string(ps.ByName("repo"))

	commits, _, err := client.Repositories.ListCommits(t1, t2, nil)

	if err != nil {
		fmt.Println(err)
	}

	commit, err := json.Marshal(commits)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(commit))

}

func repoCommitComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	session, err := store.Get(r, "github")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	at, ok := session.Values["access_token"]

	if !ok {
		fmt.Println(ok)
	}

	access_token, ok := at.(string)

	if !ok {
		fmt.Println(ok)
	}

	uid, ok := session.Values["userid"]

	if !ok {
		fmt.Println(ok)
	}

	userid, ok := uid.(string)

	if !ok {
		fmt.Println(ok)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	t1, err := strconv.Unquote(string(userid))
	t2 := string(ps.ByName("repo"))
	t3 := string(ps.ByName("sha"))

	comments, _, err := client.Repositories.ListCommitComments(t1, t2, t3, nil)

	if err != nil {
		fmt.Println(err)
	}

	comment, err := json.Marshal(comments)

	fmt.Println(string(comment))

	//var cmt []Comments
	var repoComments []Comments
	//var repoComment Comments

	if err := json.Unmarshal([]byte(string(comment)), &repoComments); err != nil {
		fmt.Println(err)
	}

	/*	for key := range repoComments {
		repoCmt, _ := json.Marshal(repoComments[key])

		if err := json.Unmarshal([]byte(string(repoCmt)), &repoComment); err != nil {
			fmt.Println(err)
		}
		body, _ := json.Marshal(repoComment.Body)
		fmt.Println(string(body))
	}*/

	cmt, err := json.Marshal(repoComments)

	fmt.Println(string(cmt))

	fmt.Fprintf(w, string(cmt))

}

func reposCommitsComments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	session, err := store.Get(r, "github")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	at, ok := session.Values["access_token"]

	if !ok {
		fmt.Println(ok)
	}

	access_token, ok := at.(string)

	if !ok {
		fmt.Println(ok)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List("", nil)

	if err != nil {
		fmt.Println(err)
	}

	repo, err := json.Marshal(repos)

	if err != nil {
		fmt.Println(err)
	}

	var repoList []Repository
	var repo1 Repository

	if err := json.Unmarshal([]byte(string(repo)), &repoList); err != nil {
		fmt.Println(err)
	}

	map1 := make(map[string]string)

	for key := range repoList {
		r1, _ := json.Marshal(repoList[key])

		if err := json.Unmarshal([]byte(string(r1)), &repo1); err != nil {
			fmt.Println(err)
		}

		name, _ := json.Marshal(repo1.Name)
		userid, _ := json.Marshal(repo1.Owner.Login)
		//id, _ := json.Marshal(repo1.ID)

		t1, err := strconv.Unquote(string(userid))
		t2, err := strconv.Unquote(string(name))

		commits, _, err := client.Repositories.ListCommits(t1, t2, nil)

		if err != nil {
			fmt.Println(err)
		}

		commit, err := json.Marshal(commits)

		if err != nil {
			fmt.Println(err)
		}

		var repoCommits []RepositoryCommit
		var commit1 RepositoryCommit

		if err := json.Unmarshal([]byte(string(commit)), &repoCommits); err != nil {
			fmt.Println(err)
		}

		for key := range repoCommits {
			repoCommit1, _ := json.Marshal(repoCommits[key])

			if err := json.Unmarshal([]byte(string(repoCommit1)), &commit1); err != nil {
				fmt.Println(err)
			}

			sha, _ := json.Marshal(commit1.SHA)

			t3, err := strconv.Unquote(string(sha))

			comments, _, err := client.Repositories.ListCommitComments(t1, t2, t3, nil)

			if err != nil {
				fmt.Println(err)
			}

			comment, err := json.Marshal(comments)
			map1[string(sha)] = string(comment)

		}

	}

	fmt.Println(map1)

	fmt.Fprintf(w, string(repo))

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/auth/github", authGithub)
	router.GET("/auth/github/callback", authGithubCallback)
	router.GET("/repos", repos)
	router.GET("/repos/:repo/comments", repoComments)
	router.GET("/repos/:repo/commits", repoCommits)
	router.GET("/repos/:repo/commits/:sha/comments", repoCommitComments)
	log.Fatal(http.ListenAndServe("localhost:1337", router))
}
