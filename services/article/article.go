package services

// Article struct
type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Count int    `json:"int"`
}

type Data struct {
	Event string `json:"event"`
}

// Repo contains all the articles created
type Repo struct {
	Articles []Article
}

// New is a pointer to the Repo Struct
func New() *Repo {
	return &Repo{}
}

// Add will append new articles to the Repo
func (r *Repo) Add(article Article) {
	article.Count = len(article.Body)
	r.Articles = append(r.Articles, article)
}

// GetAll will return all the articles in the Repo
func (r *Repo) GetAll() []Article {
	return r.Articles
}
