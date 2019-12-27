package repository

import (
	"go_microservice/resources"
	"testing"
)

func TestAdd(t *testing.T) {
	article := New()
	article.Add(resources.New{1, "hey", "hello", 0})
	for _, a := range article.Articles {
		if a.Count < 0 {
			t.Errorf("Body count cannot be zero")
		}
	}
	if len(article.Articles) != 1 {
		t.Errorf("Article was not added")
	}
}

func TestGetAll(t *testing.T) {
	article := New()
	article.Add(resources.New{})
	results := article.Get()
	if len(results) != 1 {
		t.Errorf("List is empty")
	}
}
