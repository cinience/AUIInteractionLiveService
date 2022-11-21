package _default

import (
	"ApsaraLive/pkg/models"
	"fmt"
)

func (l *LiveRoomManager) AddMeetingUser(id string, members []models.MeetingMember) (*models.MeetingInfo, error) {
	record, err := l.sa.GetRoom(id)
	if err != nil {
		return nil, err
	}

	meetingInfo := &models.MeetingInfo{}
	err = meetingInfo.ParseJsonString(record.MeetingInfo)
	if err != nil {
		return nil, err
	}

	meetingInfo.Members = append(meetingInfo.Members, members...)
	record.MeetingInfo, err = meetingInfo.ToJsonString()
	if err != nil {
		return nil, err
	}

	err = l.sa.UpdateRoom(id, record)
	if err != nil {
		return nil, err
	}

	return meetingInfo, nil
}

func (l *LiveRoomManager) UpdateMeetingUser(id string, members []models.MeetingMember) (*models.MeetingInfo, error) {
	record, err := l.sa.GetRoom(id)
	if err != nil {
		return nil, err
	}

	meetingInfo := &models.MeetingInfo{}
	err = meetingInfo.ParseJsonString(record.MeetingInfo)
	if err != nil {
		return nil, err
	}

	var newMembers []models.MeetingMember
	for _, oldItem := range meetingInfo.Members {
		for _, item := range members {
			if item.UserId == oldItem.UserId {
				oldItem.CameraOpened = item.CameraOpened
				oldItem.MicOpened = item.MicOpened
				break
			}
		}
		newMembers = append(newMembers, oldItem)
	}

	meetingInfo.Members = newMembers
	record.MeetingInfo, err = meetingInfo.ToJsonString()
	if err != nil {
		return nil, err
	}

	err = l.sa.UpdateRoom(id, record)
	if err != nil {
		return nil, nil
	}

	return meetingInfo, nil
}

func (l *LiveRoomManager) DelMeetingUser(id string, members []models.MeetingMember) (*models.MeetingInfo, error) {
	record, err := l.sa.GetRoom(id)
	if err != nil {
		return nil, err
	}

	meetingInfo := &models.MeetingInfo{}
	err = meetingInfo.ParseJsonString(record.MeetingInfo)
	if err != nil {
		return nil, err
	}

	var newMembers []models.MeetingMember
	for _, oldItem := range meetingInfo.Members {
		isExist := false
		for _, item := range members {
			if item.UserId == oldItem.UserId {
				isExist = true
				break
			}
		}
		if !isExist {
			newMembers = append(newMembers, oldItem)
		}
	}
	meetingInfo.Members = newMembers
	record.MeetingInfo, err = meetingInfo.ToJsonString()
	if err != nil {
		return nil, err
	}

	err = l.sa.UpdateRoom(id, record)
	if err != nil {
		return nil, nil
	}

	return meetingInfo, nil
}

func (l *LiveRoomManager) UpdateMeetingInfo(id string, members []models.MeetingMember) (*models.MeetingInfo, error) {
	record, err := l.sa.GetRoom(id)
	if err != nil {
		return nil, err
	}

	if record.Mode != models.LiveModeLink {
		return nil, fmt.Errorf("unsupported live mode: %v", record.Mode)
	}

	meetingInfo := &models.MeetingInfo{}
	err = meetingInfo.ParseJsonString(record.MeetingInfo)
	if err != nil {
		return nil, err
	}

	meetingInfo.Members = members
	record.MeetingInfo, err = meetingInfo.ToJsonString()
	if err != nil {
		return nil, err
	}

	err = l.sa.UpdateRoom(id, record)
	if err != nil {
		return nil, nil
	}

	return meetingInfo, nil
}

func (l *LiveRoomManager) GetMeetingInfo(id string) (*models.MeetingInfo, error) {
	record, err := l.sa.GetRoom(id)
	if err != nil {
		return nil, err
	}

	if record.Mode != models.LiveModeLink {
		return nil, fmt.Errorf("unsupported live mode: %v", record.Mode)
	}
	
	meetingInfo := &models.MeetingInfo{}
	err = meetingInfo.ParseJsonString(record.MeetingInfo)
	if err != nil {
		return nil, err
	}

	return meetingInfo, nil
}
