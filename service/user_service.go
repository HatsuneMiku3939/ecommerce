package service

import (
	"bitbucket.org/ashtishad/as_ti/domain"
	"bitbucket.org/ashtishad/as_ti/pkg/hashpassword"
)

type UserService interface {
	NewUser(request domain.NewUserRequestDTO) (*domain.UserResponseDTO, error)
	ExistingUser(request domain.ExistingUserRequestDTO) (*domain.UserResponseDTO, error)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}

// NewUser first generate a salt, hashedPassword, then creates a domain model from request dto,
// then Calls the repository to save(create/update) the new user, get the user model if everything okay, otherwise returns error
// Finally returns UserResponseDTO.
func (service DefaultUserService) NewUser(request domain.NewUserRequestDTO) (*domain.UserResponseDTO, error) {
	salt, err := hashpassword.GenerateSalt()
	if err != nil {
		return nil, err
	}

	hashedPassword := hashpassword.HashPassword(request.Password, salt)

	user := domain.User{
		Email:        request.Email,
		PasswordHash: hashedPassword,
		FullName:     request.FullName,
		Phone:        request.Phone,
		SignUpOption: request.SignUpOption,
		Status:       "active",
	}

	createdUser, err := service.repo.Save(user, salt)
	if err != nil {
		return nil, err
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

// ExistingUser calls the repository to save the new user, get the user model if everything okay, otherwise returns error
// Finally converts to UserResponseDTO.
func (service DefaultUserService) ExistingUser(request domain.ExistingUserRequestDTO) (*domain.UserResponseDTO, error) {
	existingUser, err := service.repo.FindExisting(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	userResponseDTO := &domain.UserResponseDTO{
		UserUUID:     existingUser.UserUUID,
		Email:        existingUser.Email,
		FullName:     existingUser.FullName,
		Phone:        existingUser.Phone,
		SignUpOption: existingUser.SignUpOption,
		Status:       existingUser.Status,
		CreatedAt:    existingUser.CreatedAt,
		UpdatedAt:    existingUser.UpdatedAt,
	}

	return userResponseDTO, nil
}
