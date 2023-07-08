-- golang_dbをアクティブ
DROP TABLE IF EXISTS users;
-- usersテーブルを作成。名前とパスワード
CREATE TABLE posts (
    id int AUTO_INCREMENT,
    sentence text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id)
);
-- usersテーブルに２つレコードを追加
INSERT INTO posts (sentence)
VALUES ("gophar"),("hello"),("good bye");