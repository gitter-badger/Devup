package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
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
	Login             string "json:login"
	ID                int64  "json:id"
	AvatarURL         string "json:avatar_url"
	HTMLURL           string "json:html_url"
	GravatarID        string "json:gravatar_id"
	Type              string "json:type"
	SiteAdmin         bool   "json:site_admin"
	URL               string "json:url"
	EventsURL         string "json:events_url"
	FollowingURL      string "json:following_url"
	FollowersURL      string "json:followers_url"
	GistsURL          string "json:gists_url"
	OrganizationsURL  string "json:organizations_url"
	ReceivedEventsURL string "json:received_events_url"
	ReposURL          string "json:repos_url"
	StarredURL        string "json:starred_url"
	SubscriptionsURL  string "json:subscriptions_url"
}

type Permissions struct {
	admin bool "json:admin"
	push  bool "json:push"
	pull  bool "json:pull"
}

type Repository struct {
	ID               int64 "json:id"
	Owner            User
	Name             string "json:name"
	FullName         string "json:fullName"
	Description      string "json:description"
	Homepage         string "json:homepage"
	DefaultBranch    string "json:default_branch"
	MasterBranch     string "json:master_branch"
	CreatedAt        string "json:Timestamp"
	PushedAt         string "json:Timestamp"
	UpdatedAt        string "json:Timestamp"
	HTMLURL          string "json:html_url"
	CloneURL         string "json:clone_url"
	GitURL           string "json:git_url"
	SSHURL           string "json:ssh_url"
	SVNURL           string "json:svn_url"
	Language         string "json:language"
	Fork             bool   "json:fork"
	ForksCount       string "json:forks_count"
	OpenIssuesCount  int64  "json:open_issues_count"
	StargazersCount  int64  "json:stargazers_count"
	WatchersCount    int64  "json:watchers_count"
	Size             int64  "json:size"
	Permissions      Permissions
	Private          bool   "json:private"
	HasIssues        bool   "json:has_issues"
	HasWiki          bool   "json:has_wiki"
	HasDownloads     bool   "json:has_downloads"
	URL              string "json:url"
	ArchiveURL       string "json:archive_url"
	AssigneesURL     string "json:assignees_url"
	BlobsURL         string "json:blobs_url"
	BranchesURL      string "json:branches_url"
	CollaboratorsURL string "json:collaborators_url"
	CommentsURL      string "json:comments_url"
	CommitsURL       string "json:commits_url"
	CompareURL       string "json:compare_url"
	ContentsURL      string "json:contents_url"
	ContributorsURL  string "json:contributors_url"
	DownloadsURL     string "json:downloads_url"
	EventsURL        string "json:events_url"
	ForksURL         string "json:forks_url"
	GitCommitsURL    string "json:git_commits_url"
	GitRefsURL       string "json:git_refs_url"
	GitTagsURL       string "json:git_tags_url"
	HooksURL         string "json:hooks_url"
	IssueCommentURL  string "json: issue_comment_url"
	IssueEventsURL   string "json: issue_events_url"
	IssuesURL        string "json: issues_url"
	KeysURL          string "json: keys_url"
	LabelsURL        string "json: labels_url"
	LanguagesURL     string "json: languages_url"
	MergesURL        string "json : merges_url"
	MilestonesURL    string "json: milestones_url"
	NotificationsURL string "json: notifications_url"
	PullsURL         string "json : pulls_url"
	ReleasesURL      string "json: releases_url"
	StargazersURL    string "json :stargazers_url"
	StatusesURL      string "json: statuses_url"
	SubscribersURL   string "json : subscribers_url"
	SubscriptionURL  string "json: subscription_url"
	TagsURL          string "json: tags_url"
	TreesURL         string "json : trees_url"
	TeamsURL         string "json : teams_url"
}

/*type Login struct {
  Login string "json: login"
}
type ID struct {
  ID int64 "json:id"
}
type AvatarURL struct {
  AvatarURL string "json:avatar_url"
}
type HTMLURL struct{
  HTMLURL string "json:html_url"
}
type GravatarID struct {
  GravatarID string "json:gravatar_id"
}
type Type struct {
  Type string "json:type"
}
type SiteAdmin struct {
  SiteAdmin bool "json:site_admin"
}
type URL struct {
  URL string "json:url"
}
type EventsURL struct {
  EventsURL string "json:events_url"
}
type FollowingURL struct {
  FollowingURL string "json:following_url"
}
type FollowersURL struct {
  FollowersURL string "json:followers_url"
}
type GistsURL struct {
  GistsURL string "json:gists_url"
}
type OrganizationsURL struct {
  OrganizationsURL string "json:organizations_url"
}
type ReceivedEventsURL struct {
  ReceivedEventsURL string "json:received_events_url"
}
type ReposURL struct {
  ReposURL string "json:repos_url"
}
type StarredURL struct {
  StarredURL string "json:starred_url"
}
type SubscriptionsURL struct {
  SubscriptionsURL string "json:subscriptions_url"
}
type admin struct {
  admin bool "json:admin"
}
type push struct {
  push bool "json:push"
}
type pull struct {
  pull bool "json:pull"
}
type Name struct {
  Name string "json:name"
}
type FullName struct {
  FullName string "json:fullName"
}
type Description struct {
  Description string "json:description"
}
type Homepage struct {
  Homepage string "json:homepage"
}
type DefaultBranch struct {
  DefaultBranch string "json:default_branch"
}
type MasterBranch struct {
  MasterBranch string "json:master_branch"
}
type CreatedAt struct {
  CreatedAt string "json:created_at"
}
type PushedAt struct {
  PushedAt string "json:pushed_at"
}
type UpdatedAt struct {
  UpdatedAt string "json:updated_at"
}
type CloneURL struct {
  CloneURL string "json:clone_url"
}
type GitURL struct {
  GitURL string "json:git_url"
}
type SSHURL struct {
  SSHURL string "json:ssh_url"
}
type SVNURL struct {
  SVNURL string "json:svn_url"
}
type Language struct {
  Language string "json:language"
}
type Fork struct {
  Fork bool "json:fork"
}

type ForksCount struct {
  ForksCount string "json:forks_count"
}
type OpenIssuesCount struct {
  OpenIssuesCount int64 "json:open_issues_count"
}
type StargazersCount struct {
  StargazersCount int64 "json:stargazers_count"
}
type WatchersCount struct {
  WatchersCount int64 "json:watchers_count"
}
type Size struct {
  Size int64 "json:size"
}
type Private struct {
  Private bool "json:private"
}
type HasIssues struct {
  HasIssues bool "json:has_issues"
}
type HasWiki struct {
  HasWiki bool "json:has_wiki"
}
type HasDownloads struct {
  HasDownloads bool "json:has_downloads"
}
type ArchiveURL struct {
  ArchiveURL string "json:archive_url"
}
type AssigneesURL struct {
  AssigneesURL string "json:assignees_url"
}
type BlobsURL struct {
  BlobsURL string "json:blobs_url"
}
type BranchesURL struct {
  BranchesURL string "json:branches_url"
}
type CollaboratorsURL struct {
   CollaboratorsURL string "json:collaborators_url"
}
type CommentsURL struct {
  CommentsURL string "json:comments_url"
}
type CommitsURL struct {
  CommitsURL string "json:commits_url"
}
type CompareURL struct {
  CompareURL string "json:compare_url"
}
type ContentsURL struct {
  ContentsURL string "json:contents_url"
}
type ContributorsURL struct {
  ContributorsURL string "json:contributors_url"
}
type DownloadsURL struct {
  DownloadsURL string "json:downloads_url"
}
type ForksURL struct {
  ForksURL string "json:forks_url"
}
type TeamsURL struct {
  TeamsURL string "json : teams_url"
}
type TreesURL struct {
  TreesURL string "json : trees_url"
}
type TagsURL struct {
  TagsURL string "json: tags_url"
}
type SubscriptionURL struct {
  SubscriptionURL string "json: subscription_url"
}
type SubscribersURL struct {
  SubscribersURL string "json : subscribers_url"
}
type StatusesURL struct {
  StatusesURL string "json: statuses_url"
}
type StargazersURL struct {
  StargazersURL string "json :stargazers_url"
}
type ReleasesURL struct {
 ReleasesURL string "json: releases_url"
}
type PullsURL struct {
  PullsURL string "json : pulls_url"
}
type NotificationsURL struct {
  NotificationsURL string "json: notifications_url"
}
type MilestonesURL struct {
  MilestonesURL string "json: milestones_url"
}
type MergesURL struct {
  MergesURL string "json : merges_url"
}
type LanguagesURL struct {
  LanguagesURL string "json: languages_url"
}
type LabelsURL struct {
  LabelsURL string "json: labels_url"
}
type KeysURL struct {
  KeysURL string "json: keys_url"
}
type IssuesURL struct {
  IssuesURL string "json: issues_url"
}
type IssueEventsURL struct {
  IssueEventsURL string "json: issue_events_url"
}
type IssueCommentURL struct {
  IssueCommentURL string "json: issue_comment_url"
}
type HooksURL struct {
  HooksURL string "json:hooks_url"
}
type GitTagsURL struct {
  GitTagsURL string "json:git_tags_url"
}
type GitRefsURL struct {
  GitRefsURL string "json:git_refs_url"
}
type  GitCommitsURL struct{
  GitCommitsURL string "json:git_commits_url"
}
type Owner struct{
  Login Login
  ID ID
  AvatarURL AvatarURL
  HTMLURL HTMLURL
  GravatarID GravatarID
  Type Type
  SiteAdmin SiteAdmin
  URL URL
  EventsURL EventsURL
  FollowingURL FollowingURL
  FollowersURL FollowersURL
  GistsURL GistsURL
  OrganizationsURL OrganizationsURL
  ReceivedEventsURL ReceivedEventsURL
  ReposURL ReposURL
  StarredURL StarredURL
  SubscriptionsURL SubscriptionsURL
}
type Permissions struct{
  admin admin
  push push
  pull pull
}
type  Repository struct{
  ID ID
  Owner Owner
  Name Name
  FullName FullName
  Description Description
  Homepage Homepage
  DefaultBranch DefaultBranch
  MasterBranch MasterBranch
  CreatedAt CreatedAt
  PushedAt PushedAt
  UpdatedAt UpdatedAt
  HTMLURL HTMLURL
  CloneURL CloneURL
  GitURL GitURL
  SSHURL SSHURL
  SVNURL SVNURL
  Language Language
  Fork Fork
  ForksCount ForksCount
  OpenIssuesCount OpenIssuesCount
  StargazersCount StargazersCount
  WatchersCount WatchersCount
  Size Size
  Permissions Permissions
  Private Private
  HasIssues HasIssues
  HasWiki HasWiki
  HasDownloads HasDownloads
  URL URL
  ArchiveURL ArchiveURL
  AssigneesURL AssigneesURL
  BlobsURL BlobsURL
  BranchesURL BranchesURL
  CollaboratorsURL CollaboratorsURL
  CommentsURL CommentsURL
  CommitsURL CommitsURL
  CompareURL CompareURL
  ContentsURL ContentsURL
  ContributorsURL ContributorsURL
  DownloadsURL DownloadsURL
  EventsURL EventsURL
  ForksURL ForksURL
  GitCommitsURL GitCommitsURL
  GitRefsURL GitRefsURL
  GitTagsURL GitTagsURL
  HooksURL HooksURL
  IssueCommentURL IssueCommentURL
  IssueEventsURL IssueEventsURL
  IssuesURL IssuesURL
  KeysURL KeysURL
  LabelsURL LabelsURL
  LanguagesURL LanguagesURL
  MergesURL MergesURL
  MilestonesURL MilestonesURL
  NotificationsURL NotificationsURL
  PullsURL PullsURL
  ReleasesURL ReleasesURL
  StargazersURL StargazersURL
  StatusesURL StatusesURL
  SubscribersURL SubscribersURL
  SubscriptionURL SubscriptionURL
  TagsURL TagsURL
  TreesURL TreesURL
  TeamsURL TeamsURL
}*/

func serveRest(w http.ResponseWriter, r *http.Request) {

	/*	response, err := getJsonResponse()
		if err != nil {
			panic(err)
		}*/
	url, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(url))
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

	fmt.Println(string(url))

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

	//fmt.Println(repos)

	repo, err := json.Marshal(repos)

	if err != nil {
		panic(err)
	}

	var dat []Repository

	if err := json.Unmarshal(repo, &dat); err != nil {
		fmt.Println(err)
	}
	/*
		for key := range dat {
			//fmt.Println(key, dat[key])

			var rep Repository

			if err := json.Unmarshal(dat[key], &rep); err != nil {
				fmt.Println(err)
			}
		}*/
	//fmt.Println(string(repo))
	fmt.Println(repo)
	fmt.Fprintf(w, string(repo))
	//http.Redirect(w, r, "http://localhost:1337/", 301)

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
