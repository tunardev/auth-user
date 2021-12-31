package entity

import "errors"

type Task struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId int64  `json:"user_id"`
	Status bool `json:"status"`
}

func (task Task) IsValidCreate() error {

	if task.Title == "" {
		return errors.New("title is required")
	}

	if len(task.Title) < 5 {
		return errors.New("title must be at least 5 character")
	}

	if len(task.Title) > 50 {
		return errors.New("title is too long")
	}

	if task.Content == "" {
		return errors.New("content is required")
	}

	if len(task.Content) < 30 {
		return errors.New("content must be at least 30 character")
	}

	if len(task.Content) > 500 {
		return errors.New("content is too long")
	}

	return nil
}

func (task Task) IsValidUpdate() error {
	if task.Title == "" {
		return errors.New("title is required")
	}

	if len(task.Title) < 5 {
		return errors.New("title must be at least 5 character")
	}

	if len(task.Title) > 50 {
		return errors.New("title is too long")
	}

	if task.Content == "" {
		return errors.New("content is required")
	}

	if len(task.Content) < 30 {
		return errors.New("content must be at least 30 character")
	}

	if len(task.Content) > 500 {
		return errors.New("content is too long")
	}

	return nil
}