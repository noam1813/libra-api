package PostController

import (
	"encoding/json"
	"log"
	"net/http"

	DB "app/database"
	Models "app/models"
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
