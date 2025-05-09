package contacts

import (
	"time"
)

type Service interface {
	CreateNewContact(name, email, phone, categoryID string) (*Contact, error)
	GetAllContacts() ([]*Contact, error)
	GetContactByID(id string) (*Contact, error)
	UpdateContact(id, name, email, phone, categoryID string) (*Contact, error)
	DeleteContact(id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateNewContact(name, email, phone, categoryID string) (*Contact, error) {
	now := time.Now()

	contact := &Contact{
		Name:       name,
		Email:      email,
		Phone:      phone,
		CategoryID: categoryID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if err := s.repo.Create(contact); err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *service) GetAllContacts() ([]*Contact, error) {
	return s.repo.FindAll()
}

func (s *service) GetContactByID(id string) (*Contact, error) {
	return s.repo.FindByID(id)
}

func (s *service) UpdateContact(id, name, email, phone, categoryID string) (*Contact, error) {
	contact, err := s.GetContactByID(id)
	if err != nil {
		return nil, err
	}

	contact.Name = name
	contact.Email = email
	contact.Phone = phone
	contact.CategoryID = categoryID
	contact.UpdatedAt = time.Now()

	if err := s.repo.Update(contact); err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *service) DeleteContact(id string) error {
	return s.repo.Delete(id)
}
