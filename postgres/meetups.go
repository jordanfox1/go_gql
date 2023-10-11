package postgres

import (
	"go_gql/graph/model"

	"github.com/go-pg/pg/v10"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups() ([]*model.Meetup, error) {
	var meetups []*model.Meetup

	err := m.DB.Model(&meetups).Order("id").Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}

func (m *MeetupsRepo) CreateMeetup(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()
	return meetup, err
}

func (m *MeetupsRepo) GetByID(id string) (*model.Meetup, error) {
	var meetup model.Meetup
	err := m.DB.Model(&meetup).Where("id = ?", id).First()
	return &meetup, err
}

func (m *MeetupsRepo) Update(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Update()
	return meetup, err
}

func (m *MeetupsRepo) Delete(meetup *model.Meetup) error {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Delete()
	return err
}

func (m *MeetupsRepo) GetMeetupsForUser(user *model.User) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := m.DB.Model(&meetups).Where("user_id = ?", user.ID).Order("id").Select()

	return meetups, err
}
