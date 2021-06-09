package main

import (
	"fmt"
)

type VideoEntity struct {
	name string
	url  string
}

type VideoI interface {
	Save(VideoEntity) VideoEntity
	FindAll() []VideoEntity
}

type VideoSvc struct {
	videos []VideoEntity
}

func New() VideoI {
	return &VideoSvc{
		videos: []VideoEntity{},
	}
}

func (video *VideoSvc) FindAll() []VideoEntity {
	return video.videos
}

func (video *VideoSvc) Save(v VideoEntity) VideoEntity {
	video.videos = append(video.videos, v)
	return v
}

func main() {
	fmt.Println("hello work")
	netflix := VideoEntity{
		name: "Minh film",
		url:  "minh.com",
	}

	var videoChildren VideoI = New()
	videoChildren.Save(netflix)
	getV := videoChildren.FindAll()
	fmt.Println(getV)
	fmt.Println(netflix)

}
