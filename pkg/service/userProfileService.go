package service

import "ozinshe/pkg/entity"

type UserProfileService interface {
	GetUserProfileByUserId(int) (entity.UserProfile, error)
	CreateUserProfile(*entity.UserProfile) error
	UpdateUserProfile(*entity.UserProfile) error
}

func (s *Service) GetUserProfileByUserId(userId int) (entity.UserProfile, error) {
	return s.repo.GetUserProfileByUserId(userId)
}

func (s *Service) CreateUserProfile(userProfile *entity.UserProfile) error {
	return s.repo.CreateUserProfile(userProfile)
}
func (s *Service) UpdateUserProfile(userProfile *entity.UserProfile) error {
	return s.repo.UpdateUserProfile(userProfile)
}

//func(s *Service) DoesUserProfileExist(userId int) bool {
//	userProfile,err := s.repo.GetUserProfileByUserId(userId)
//	if err.Error()  == entity.ErrNoRows {
//		return false
//	}
//}
