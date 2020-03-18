package api

import (
	"encoding/json"
	"fmt"
	"github.com/berkansasmaz/kartaca-internship-application.git/auth"
	"net/http"
)

// UserJSON - json data expected for login/signup
type PostJSON struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}

// NewPost -
func (api *API) NewPost(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := PostJSON{}
	err := decoder.Decode(&jsondata)
	fmt.Println(err)
	fmt.Println(jsondata)
	fmt.Println(jsondata)
	// todo karakter kontrolü yap.
	if err != nil || jsondata.Text == "" || jsondata.Username == "" || len(jsondata.Text) >= 140 {
		http.Error(w, "Missing text or username", http.StatusBadRequest)
		return
	}

	api.posts.AddPost(jsondata.Text, jsondata.Username)
	user := api.users.FindUser(jsondata.Username)
	fmt.Println(user)
	jsontoken := auth.GetJSONToken(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsontoken))
}

// Get Posts -
func (api *API) GetPosts(w http.ResponseWriter, req *http.Request) {
	posts := api.posts.GetPosts()
	fmt.Println("getposts api", posts)
	js, _ := json.Marshal(posts)

	w.Write(js)
}
