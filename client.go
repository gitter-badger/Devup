package main 

import(
	"encoding/json"
	"fmt"
  "io/ioutil"
	"net/http"
  "strings"
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
)

type Payload struct{
    Stuff Data
}

type Data struct{
    Fruit Fruits
    Veggies Vegetables
}

type Token struct{
  Access_token string "json:access_token"
  Scope string "json:scope"
  Token_type string "json:token_type"
}

type Fruits map[string]int
type Vegetables map[string]int

type User struct{
  Login string "json:Login"
  ID int32 "json:ID"
  AvatarURL string "json:AvatarURL"
  HTMLURL string "json:HTMLURL"
  GravatarID string "json:GravatarID"
  Type string "json:Type"
  SiteAdmin bool "json:SiteAdmin"
  URL string "json:URL"
  EventsURL string "json:EventsURL"
  FollowingURL string "json:FollowingURL"
  FollowersURL string "json:FollowersURL"
  GistsURL string "json:GistsURL"
  OrganizationsURL string "json:OrganizationsURL"
  ReceivedEventsURL string "json:ReceivedEventsURL"
  ReposURL string "json:ReposURL"
  StarredURL string "json:StarredURL"
  SubscriptionsURL string "json:SubscriptionsURL"

}

type Owner struct{
  User User
}        

type CreatedAt struct{
  Timestamp string "json:Timestamp"
} 

type PushedAt struct{
  Timestamp string "json:Timestamp"
} 

type UpdatedAt struct{
  Timestamp string "json:Timestamp"
} 

type Permissions struct{
  admin bool "json:admin"
  push bool "json:push"
  pull bool "json:pull"
}

type  Repository struct{
  ID string "json:ID"
  Owner Owner
  Name string "json:Name"
  FullName string "json:FullName"
  Description string "json:Description"
  Homepage string "json:Homepage"
  DefaultBranch string "json:DefaultBranch"
  MasterBranch string "json:MasterBranch"
  CreatedAt CreatedAt
  PushedAt PushedAt
  UpdatedAt UpdatedAt
  HTMLURL string "json:HTMLURL"
  CloneURL string "json:CloneURL"
  GitURL string "json:GitURL"
  SSHURL string "json:SSHURL"
  SVNURL string "json:SVNURL"
  Language string "json:Language"
  Fork bool "json:Fork"
  ForksCount int32 "json:ForksCount"
  OpenIssuesCount int32 "json:OpenIssuesCount"
  StargazersCount int32 "json:StargazersCount"
  WatchersCount int32 "json:WatchersCount"
  Size int32 "json:Size"
  Permissions Permissions
  Private bool "json:Private"
  HasIssues bool "json:HasIssues"
  HasWiki bool "json:HasWiki"
  HasDownloads bool "json:HasDownloads"
  URL string "json:URL"
  ArchiveURL string "json:ArchiveURL"
  AssigneesURL string "json:AssigneesURL"
  BlobsURL string "json:BlobsURL"
  BranchesURL string "json:BranchesURL"
  CollaboratorsURL string "json:CollaboratorsURL"
  CommentsURL string "json:CommentsURL"
  CommitsURL string "json:CommitsURL"
  CompareURL string "json:CompareURL"
  ContentsURL string "json:ContentsURL"
  ContributorsURL string "json:ContributorsURL"
  DownloadsURL string "json:DownloadsURL"
  EventsURL string "json:EventsURL"
  ForksURL string "json:ForksURL"
  GitCommitsURL string "json:GitCommitsURL"
  GitRefsURL string "json:GitRefsURL"
  GitTagsURL string "json:GitTagsURL"
  HooksURL string "json:HooksURL"
  IssueCommentURL string "json: IssueCommentURL"
  IssueEventsURL string "json: IssueEventsURL"
  IssuesURL string "json: IssuesURL"
  KeysURL string "json: KeysURL"
  LabelsURL string "json: LabelsURL"
  LanguagesURL string "json: LanguagesURL"
  MergesURL string "json : MergesURL"
  MilestonesURL string "json: MilestonesURL"
  NotificationsURL string "json: NotificationsURL"
  PullsURL string "json : PullsURL"
  ReleasesURL string "json: ReleasesURL"
  StargazersURL string "json :StargazersURL"
  StatusesURL string "json: StatusesURL"
  SubscribersURL string "json : SubscribersURL"
  SubscriptionURL string "json: SubscriptionURL"
  TagsURL string "json: TagsURL"
  TreesURL string "json : TreesURL"
  TeamsURL string "json : TeamsURL"
}



func serveRest(w http.ResponseWriter, r *http.Request){
    response, err := getJsonResponse()
    if err != nil {
        panic(err);
    }

    fmt.Fprintf(w, string(response))
}

func serveRest1(w http.ResponseWriter, r *http.Request){
    http.Redirect(w, r, "https://github.com/login/oauth/authorize?client_id=c5376f36b92a55ac20e1", 301)
}

func serveRest2(w http.ResponseWriter, r *http.Request){
    code := r.URL.Query().Get("code")

    res, err := http.Get("https://github.com/login/oauth/access_token?client_id=c5376f36b92a55ac20e1&client_secret=a2933fb99800ed7b683b70a139f2691d5e2953e8&code="+code+"") 

    if err != nil{
        panic(err)
    }

    defer res.Body.Close()

    url, err := ioutil.ReadAll(res.Body)

    if err != nil{
      panic(err)
    }

    fmt.Println(string(url))

    s := strings.Split(string(url), "&")

    fmt.Println(s);

    at := s[0]

    fmt.Println(at)

    access_token := strings.Split(at, "=")

    fmt.Println(access_token[1])

    ts := oauth2.StaticTokenSource(
      &oauth2.Token{AccessToken: access_token[1]},
    )
    tc := oauth2.NewClient(oauth2.NoContext, ts)

    client := github.NewClient(tc)

    // list all repositories for the authenticated user
    repos, _, err := client.Repositories.List("", nil)

    if err != nil{
      panic(err)
    }
    
    //var p []Repository

    p, err := json.Marshal(repos)

    if err != nil{
      panic(err)
    }

    //p1 := json.MarshalIndent(p, "", " ")
    p1 := string(p)
    len1 := len(repos)
    fmt.Println(len1)
 /*   for i, repo := range repos {
      fmt.Println(i)
      fmt.Println(repo["github.Repository"])
      fmt.Println(string(repo.id))
      comments, _, err := client.Repositories.ListComments(string(repo.owner.id),string(repo.id), nil)
      fmt.Println(comments)
    }*/

    //fmt.Println(string(p)) 

    fmt.Fprintf(w, p1)
}


func main() {

	http.HandleFunc("/", serveRest)
  http.HandleFunc("/github", serveRest1)
  http.HandleFunc("/github/callback", serveRest2)
  http.ListenAndServe("localhost:1337", nil) 

}

func getJsonResponse() ([]byte, error){

    fruits := make(map[string]int)
    fruits["Banana"] = 23
    fruits["Apple"] = 4

    vegetables := make(map[string]int)
    vegetables["Carrets"] = 44
    vegetables["Goolse"] = 7

    d := Data{fruits, vegetables}
    p := Payload{d}

    return json.MarshalIndent(p, "", " ")

}