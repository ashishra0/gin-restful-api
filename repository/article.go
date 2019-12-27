package repository

import "go_microservice/resources"

// Repo contains all the articles created
type Repo struct {
	Articles []resources.New
}

// New is a pointer to the Repo Struct
func New() *Repo {
	return &Repo{}
}

// Add will append new articles to the Repo
func (r *Repo) Add(article resources.New) map[string]interface{} {
	article.Count = len(article.Body)
	r.Articles = append(r.Articles, article)
	return map[string]interface{}{
		"id":    article.ID,
		"title": article.Title,
		"body":  article.Body,
		"count": len(article.Body),
	}
}

// Get will return all the articles in the Repo
func (r *Repo) Get() []resources.New {
	return r.Articles
}
