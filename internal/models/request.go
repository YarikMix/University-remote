package models

type Request struct {
	AccessToken int64 `json:"access_token"`
	LessonId    int64 `json:"lesson_id"`
}
