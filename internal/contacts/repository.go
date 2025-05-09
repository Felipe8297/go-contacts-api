package contacts

import "database/sql"

type Repository interface {
	Create(contact *Contact) error
	FindAll() ([]*Contact, error)
	FindByID(id string) (*Contact, error)
	Update(contact *Contact) error
	Delete(id string) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(contact *Contact) error {
	query := `
		INSERT INTO contacts (name, email, phone, category_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id string
	err := r.db.QueryRow(query, contact.Name, contact.Email, contact.Phone, contact.CategoryID, contact.CreatedAt, contact.UpdatedAt).Scan(&id)
	if err != nil {
		return err
	}

	contact.ID = id
	return nil
}

func (r *PostgresRepository) FindAll() ([]*Contact, error) {
	query := `
		SELECT id, name, email, phone, category_id, created_at, updated_at
		FROM contacts
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	contacts := []*Contact{}

	for rows.Next() {
		contact := &Contact{}
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone, &contact.CategoryID, &contact.CreatedAt, &contact.UpdatedAt); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *PostgresRepository) FindByID(id string) (*Contact, error) {
	query := `
		SELECT id, name, email, phone, category_id, created_at, updated_at
		FROM contacts
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	contact := &Contact{}

	if err := row.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone, &contact.CategoryID, &contact.CreatedAt, &contact.UpdatedAt); err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *PostgresRepository) Update(contact *Contact) error {

	query := `
		UPDATE contacts
		SET name = $1, email = $2, phone = $3, category_id = $4, updated_at = $5
		WHERE id = $6
	`

	_, err := r.db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.CategoryID, contact.UpdatedAt, contact.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) Delete(id string) error {
	query := `
		DELETE FROM contacts
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
