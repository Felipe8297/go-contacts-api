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

// NewService cria uma nova instância do serviço de contatos
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// CreateNewContact cria um novo contato
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

// GetAllContacts retorna todos os contatos
func (s *service) GetAllContacts() ([]*Contact, error) {
	return s.repo.FindAll()
}

// GetContactByID busca um contato pelo ID
func (s *service) GetContactByID(id string) (*Contact, error) {
	return s.repo.FindByID(id)
}

// UpdateContact atualiza um contato existente
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

// DeleteContact remove um contato pelo ID
func (s *service) DeleteContact(id string) error {
	return s.repo.Delete(id)
}
