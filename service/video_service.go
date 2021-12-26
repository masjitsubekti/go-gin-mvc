package service

import "github.com/masjitsubekti/go-gin-mvc/model"

// Class Abstract Interface
type VideoService interface {
	Save(model.Video) model.Video
	FindAll() []model.Video
}

type videoService struct {
	videos []model.Video
}

// Constructor
func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video model.Video) model.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []model.Video {
	return service.videos
}
