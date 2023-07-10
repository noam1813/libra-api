package PostController

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	DB "app/database"
	Models "app/models"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {

	db := DB.SqlConnect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var posts []Models.Post
	for rows.Next() {
		var post Models.Post
		err := rows.Scan(&post.ID, &post.Sentence, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}

	// レスポンスを設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := DB.SqlConnect()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		panic("Error!")
	}

	row, err := db.Query("SELECT * FROM posts WHERE id=? ORDER BY updated_at DESC", id)
	if err != nil {
		panic(err.Error())
	}
	defer row.Close()

	post := Models.Post{}
	for row.Next() {
		err = row.Scan(&post.ID, &post.Sentence, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
		if err != nil {
			panic(err.Error())
		}
	}

	jsonData, err := json.Marshal(post)
	if err != nil {
		log.Fatal(err)
	}

	// レスポンスを設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func Create(w http.ResponseWriter, r *http.Request) {
	db := DB.SqlConnect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}

	sentence := r.FormValue("sentence")

	insert, err := db.Prepare(fmt.Sprintf("INSERT INTO posts(sentence) VALUES (?)"))
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(sentence)
	w.WriteHeader(http.StatusOK)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := DB.SqlConnect()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}

	sentence := r.FormValue("sentence")

	update, err := db.Prepare("UPDATE posts SET sentence=? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(sentence, id)
	w.WriteHeader(http.StatusOK)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := DB.SqlConnect()
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	delete, err := db.Prepare("DELETE from posts WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	w.WriteHeader(http.StatusOK)
}
