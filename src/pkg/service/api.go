package service

import (
	"ApsaraLive/pkg/models"
)

type LiveRoomManagerAPI interface {
	GetIMToken(env, userId, deviceId, deviceType string) (string, error)

	GetRoomList(pageSize int, pageNum int, status int, role string) (int64, []*models.RoomInfo, error)

	CreateRoom(title, notice, coverUrl string, Anchor string, extend string, mode int) (*models.RoomInfo, error)

	GetRoom(id, role string) (*models.RoomInfo, error)

	UpdateRoom(id string, title, notice string, extend string) (*models.RoomInfo, error)

	DeleteRoom(id string) (*models.RoomInfo, error)

	StartLive(id string, role string) (*models.RoomInfo, error)

	PauseLive(id string, role string) (*models.RoomInfo, error)

	StopLive(id string, role string) (*models.RoomInfo, error)

	UpdateMeetingInfo(id string, members []models.MeetingMember) (*models.MeetingInfo, error)

	GetMeetingInfo(id string) (*models.MeetingInfo, error)
}
