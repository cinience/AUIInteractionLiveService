package service

import (
	"ApsaraLive/pkg/models"
)

type LiveRoomManagerAPI interface {
	GetIMToken(env, userId, deviceId, deviceType string) (string, error)

	GetRoomList(pageSize int, pageNum int, role string) ([]*models.RoomInfo, error)

	CreateRoom(title, notice, Anchor string, extend string, mode int) (*models.RoomInfo, error)

	GetRoom(id, role string) (*models.RoomInfo, error)

	UpdateRoom(id string, title, notice string, extend string) (*models.RoomInfo, error)

	StartRoom(id string, role string) (*models.RoomInfo, error)

	PauseRoom(id string, role string) (*models.RoomInfo, error)

	StopRoom(id string, role string) (*models.RoomInfo, error)

	DeleteRoom(id string) (*models.RoomInfo, error)

	AddMeetingUser(id string, members []models.MeetingMember) (*models.MeetingInfo, error)

	UpdateMeetingUser(id string, members []models.MeetingMember) (*models.MeetingInfo, error)

	DelMeetingUser(id string, members []models.MeetingMember) (*models.MeetingInfo, error)

	GetMeetingInfo(id string) (*models.MeetingInfo, error)
}
