package service

import (
	"bitbucket.org/ashtishad/as_ti/domain"
	"fmt"
)

type UserService interface {
	NewUser(request domain.NewUserRequestDTO) (*domain.UserResponseDTO, error)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}

// NewUser first converts Convert request DTO to a user domain model
// then Calls the repository to save the new user, get the user model if everything okay, otherwise returns error
// Finally converts to UserResponseDTO.
func (service DefaultUserService) NewUser(request domain.NewUserRequestDTO) (*domain.UserResponseDTO, error) {
	user := domain.User{
		Email:        request.Email,
		PasswordHash: request.Password, // You must hash this before saving
		FullName:     request.FullName,
		Phone:        request.Phone,
		SignUpOption: request.SignUpOption,
		Status:       "active", // Set initial status or based on request
	}

	createdUser, err := service.repo.Save(user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	userResponseDTO := &domain.UserResponseDTO{
		UserUUID:     createdUser.UserUUID,
		Email:        createdUser.Email,
		FullName:     createdUser.FullName,
		Phone:        createdUser.Phone,
		SignUpOption: createdUser.SignUpOption,
		Status:       createdUser.Status,
		CreatedAt:    createdUser.CreatedAt,
		UpdatedAt:    createdUser.UpdatedAt,
	}

	return userResponseDTO, nil
}
