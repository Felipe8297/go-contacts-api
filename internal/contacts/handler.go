// @title           Contacts API
// @version         1.0
// @description     API para gerenciamento de contatos
// @host            localhost:8080
// @BasePath        /

package contacts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

// @Description Dados para criação de um contato
type CreateContactRequest struct {
	Name       string `json:"name" binding:"required" example:"João Silva"`               // Nome do contato
	Email      string `json:"email" binding:"required,email" example:"joao@example.com"`  // Email do contato
	Phone      string `json:"phone" example:"11999998888"`                                // Telefone do contato
	CategoryID string `json:"category_id" example:"123e4567-e89b-12d3-a456-426614174000"` // ID da categoria
}

// @Description Dados para atualização de um contato
type UpdateContactRequest struct {
	Name       string `json:"name" example:"João Silva Atualizado"`                       // Nome do contato
	Email      string `json:"email" example:"joao.novo@example.com"`                      // Email do contato
	Phone      string `json:"phone" example:"11999997777"`                                // Telefone do contato
	CategoryID string `json:"category_id" example:"123e4567-e89b-12d3-a456-426614174999"` // ID da categoria
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	contacts := router.Group("/contacts")
	{
		contacts.POST("", h.CreateContact)
		contacts.GET("", h.GetAllContacts)
		contacts.GET("/:id", h.GetContactByID)
		contacts.PUT("/:id", h.UpdateContact)
		contacts.DELETE("/:id", h.DeleteContact)
	}
}

// @Summary     Criar um novo contato
// @Description Cria um novo contato com as informações fornecidas
// @Tags        contacts
// @Accept      json
// @Produce     json
// @Param       request body CreateContactRequest true "Dados do contato"
// @Success     201 {object} Contact
// @Failure     400 {object} ErrorResponse "Erro de validação dos dados"
// @Failure     500 {object} ErrorResponse "Erro interno do servidor"
// @Router      /contacts [post]
func (h *Handler) CreateContact(c *gin.Context) {
	var req CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := h.service.CreateNewContact(req.Name, req.Email, req.Phone, req.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, contact)
}

// @Summary     Listar todos os contatos
// @Description Retorna uma lista de todos os contatos cadastrados
// @Tags        contacts
// @Accept      json
// @Produce     json
// @Success     200 {array} Contact
// @Failure     500 {object} ErrorResponse "Erro interno do servidor"
// @Router      /contacts [get]
func (h *Handler) GetAllContacts(c *gin.Context) {
	contacts, err := h.service.GetAllContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contacts)
}

// @Summary     Buscar contato por ID
// @Description Retorna um contato específico com base no ID fornecido
// @Tags        contacts
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do contato"
// @Success     200 {object} Contact
// @Failure     404 {object} ErrorResponse "Contato não encontrado"
// @Router      /contacts/{id} [get]
func (h *Handler) GetContactByID(c *gin.Context) {
	id := c.Param("id")

	contact, err := h.service.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contato não encontrado"})
		return
	}
	c.JSON(http.StatusOK, contact)
}

// @Summary     Atualizar contato
// @Description Atualiza os dados de um contato existente
// @Tags        contacts
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do contato"
// @Param       request body UpdateContactRequest true "Dados atualizados do contato"
// @Success     200 {object} Contact
// @Failure     400 {object} ErrorResponse "Erro de validação dos dados"
// @Failure     404 {object} ErrorResponse "Contato não encontrado"
// @Failure     500 {object} ErrorResponse "Erro interno do servidor"
// @Router      /contacts/{id} [put]
func (h *Handler) UpdateContact(c *gin.Context) {
	id := c.Param("id")

	var req UpdateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := h.service.UpdateContact(id, req.Name, req.Email, req.Phone, req.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contact)
}

// @Summary     Excluir contato
// @Description Remove um contato existente da base de dados
// @Tags        contacts
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do contato"
// @Success     204 "Contato removido com sucesso"
// @Failure     404 {object} ErrorResponse "Contato não encontrado"
// @Failure     500 {object} ErrorResponse "Erro interno do servidor"
// @Router      /contacts/{id} [delete]
func (h *Handler) DeleteContact(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteContact(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// ErrorResponse representa uma resposta de erro da API
// @Description Estrutura padrão para respostas de erro
type ErrorResponse struct {
	Error string `json:"error" example:"Mensagem de erro"` // Mensagem de erro
}
