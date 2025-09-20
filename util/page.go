package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// Paginator is the main pagination struct that handles all pagination logic
type Paginator struct {
	Page       int         `json:"page" form:"page" query:"page"`
	PerPage    int         `json:"per_page" form:"per_page" query:"per_page"`
	Sort       string      `json:"sort" form:"sort" query:"sort"`
	Order      string      `json:"order" form:"order" query:"order"`
	Search     string      `json:"search" form:"search" query:"search"`
	Filters    QueryFilter `json:"filters" form:"filters" query:"filters"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	FromRow    int         `json:"from_row"`
	ToRow      int         `json:"to_row"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_previous"`
}

// QueryFilter represents additional filters that can be applied
type QueryFilter map[string]interface{}

// PaginatedResponse represents the standardized API response with pagination
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination *Paginator  `json:"pagination"`
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
}

// Config holds pagination configuration
type Config struct {
	DefaultPerPage int      `json:"default_per_page"`
	MaxPerPage     int      `json:"max_per_page"`
	DefaultSort    string   `json:"default_sort"`
	DefaultOrder   string   `json:"default_order"`
	AllowedSorts   []string `json:"allowed_sorts"`
}

// DefaultConfig returns sensible defaults
func DefaultConfig() *Config {
	return &Config{
		DefaultPerPage: 20,
		MaxPerPage:     100,
		DefaultSort:    "created_at",
		DefaultOrder:   "desc",
		AllowedSorts:   []string{"id", "created_at", "updated_at", "name"},
	}
}

// New creates a new Paginator with default values
func New(config ...*Config) *Paginator {
	cfg := DefaultConfig()
	if len(config) > 0 && config[0] != nil {
		cfg = config[0]
	}

	return &Paginator{
		Page:    1,
		PerPage: cfg.DefaultPerPage,
		Sort:    cfg.DefaultSort,
		Order:   cfg.DefaultOrder,
		Filters: make(QueryFilter),
	}
}

// FromQuery parses pagination parameters from query string or form data
func (p *Paginator) FromQuery(params map[string][]string, config ...*Config) *Paginator {
	cfg := DefaultConfig()
	if len(config) > 0 && config[0] != nil {
		cfg = config[0]
	}

	// Parse page
	if pageStr, exists := params["page"]; exists && len(pageStr) > 0 {
		if page, err := strconv.Atoi(pageStr[0]); err == nil && page > 0 {
			p.Page = page
		}
	}

	// Parse per_page
	if perPageStr, exists := params["per_page"]; exists && len(perPageStr) > 0 {
		if perPage, err := strconv.Atoi(perPageStr[0]); err == nil && perPage > 0 {
			if perPage > cfg.MaxPerPage {
				p.PerPage = cfg.MaxPerPage
			} else {
				p.PerPage = perPage
			}
		}
	}

	// Parse sort
	if sortStr, exists := params["sort"]; exists && len(sortStr) > 0 {
		sort := strings.TrimSpace(sortStr[0])
		if p.isAllowedSort(sort, cfg.AllowedSorts) {
			p.Sort = sort
		}
	}

	// Parse order
	if orderStr, exists := params["order"]; exists && len(orderStr) > 0 {
		order := strings.ToLower(strings.TrimSpace(orderStr[0]))
		if order == "asc" || order == "desc" {
			p.Order = order
		}
	}

	// Parse search
	if searchStr, exists := params["search"]; exists && len(searchStr) > 0 {
		p.Search = strings.TrimSpace(searchStr[0])
	}

	return p
}

// isAllowedSort checks if the sort field is allowed
func (p *Paginator) isAllowedSort(sort string, allowedSorts []string) bool {
	for _, allowed := range allowedSorts {
		if sort == allowed {
			return true
		}
	}
	return false
}

// Paginate applies pagination to a GORM query
func (p *Paginator) Paginate(value interface{}, db *gorm.DB, result interface{}) *PaginatedResponse {
	// Count total records
	var totalRows int64
	countDB := db.Model(value)

	// Apply search if provided
	if p.Search != "" {
		countDB = p.applySearch(countDB)
	}

	// Apply filters if provided
	if len(p.Filters) > 0 {
		countDB = p.applyFilters(countDB)
	}

	countDB.Count(&totalRows)

	// Calculate pagination metadata
	p.TotalRows = totalRows
	p.TotalPages = int(math.Ceil(float64(totalRows) / float64(p.PerPage)))

	if p.Page > p.TotalPages {
		p.Page = p.TotalPages
	}
	if p.Page < 1 {
		p.Page = 1
	}

	p.FromRow = (p.Page-1)*p.PerPage + 1
	p.ToRow = p.Page * p.PerPage
	if p.ToRow > int(totalRows) {
		p.ToRow = int(totalRows)
	}

	p.HasNext = p.Page < p.TotalPages
	p.HasPrev = p.Page > 1

	// Apply pagination to the main query
	offset := (p.Page - 1) * p.PerPage
	queryDB := db.Model(value).Offset(offset).Limit(p.PerPage)

	// Apply search
	if p.Search != "" {
		queryDB = p.applySearch(queryDB)
	}

	// Apply filters
	if len(p.Filters) > 0 {
		queryDB = p.applyFilters(queryDB)
	}

	// Apply sorting
	if p.Sort != "" {
		orderBy := fmt.Sprintf("%s %s", p.Sort, strings.ToUpper(p.Order))
		queryDB = queryDB.Order(orderBy)
	}

	// Execute query
	queryDB.Find(result)

	return &PaginatedResponse{
		Data:       result,
		Pagination: p,
		Success:    true,
	}
}

// applySearch applies search functionality (customize based on your needs)
func (p *Paginator) applySearch(db *gorm.DB) *gorm.DB {
	if p.Search == "" {
		return db
	}

	// Example: search in name and description fields
	// Modify this based on your model's searchable fields
	searchTerm := fmt.Sprintf("%%%s%%", p.Search)
	return db.Where("name ILIKE ? OR description ILIKE ?", searchTerm, searchTerm)
}

// applyFilters applies additional filters
func (p *Paginator) applyFilters(db *gorm.DB) *gorm.DB {
	for key, value := range p.Filters {
		switch key {
		case "status":
			db = db.Where("status = ?", value)
		case "category_id":
			db = db.Where("category_id = ?", value)
		case "created_after":
			db = db.Where("created_at >= ?", value)
		case "created_before":
			db = db.Where("created_at <= ?", value)
		// Add more filter cases as needed
		default:
			// Generic filter - be careful with this in production
			db = db.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	return db
}

// AddFilter adds a filter to the paginator
func (p *Paginator) AddFilter(key string, value interface{}) *Paginator {
	if p.Filters == nil {
		p.Filters = make(QueryFilter)
	}
	p.Filters[key] = value
	return p
}

// Response creates a standardized paginated response
func (p *Paginator) Response(data interface{}, message ...string) *PaginatedResponse {
	resp := &PaginatedResponse{
		Data:       data,
		Pagination: p,
		Success:    true,
	}

	if len(message) > 0 {
		resp.Message = message[0]
	}

	return resp
}

// Example usage functions

// User represents a sample user model
type User struct {
	ID         uint   `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Status     string `json:"status"`
	CategoryID uint   `json:"category_id"`
	CreatedAt  int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  int64  `json:"updated_at" gorm:"autoUpdateTime"`
}

// UserService demonstrates how to use the paginator in a service
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new user service
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GetUsers demonstrates pagination usage in a service method
func (s *UserService) GetUsers(params map[string][]string) *PaginatedResponse {
	// Create paginator with custom config
	config := &Config{
		DefaultPerPage: 15,
		MaxPerPage:     50,
		DefaultSort:    "created_at",
		DefaultOrder:   "desc",
		AllowedSorts:   []string{"id", "name", "email", "created_at", "updated_at"},
	}

	paginator := New(config).FromQuery(params, config)

	// Add custom filters based on query parameters
	if status, exists := params["status"]; exists && len(status) > 0 {
		paginator.AddFilter("status", status[0])
	}

	if categoryID, exists := params["category_id"]; exists && len(categoryID) > 0 {
		paginator.AddFilter("category_id", categoryID[0])
	}

	var users []User
	return paginator.Paginate(&User{}, s.db, &users)
}

// Example HTTP handler using Gin
/*
func (s *UserService) HandleGetUsers(c *gin.Context) {
	// Parse query parameters
	params := make(map[string][]string)
	for key, values := range c.Request.URL.Query() {
		params[key] = values
	}

	// Get paginated users
	response := s.GetUsers(params)

	c.JSON(200, response)
}
*/

// Example response structure:
/*
{
	"data": [
		{
			"id": 1,
			"name": "John Doe",
			"email": "john@example.com",
			"status": "active",
			"category_id": 1,
			"created_at": 1640995200,
			"updated_at": 1640995200
		}
	],
	"pagination": {
		"page": 1,
		"per_page": 15,
		"sort": "created_at",
		"order": "desc",
		"search": "",
		"total_rows": 150,
		"total_pages": 10,
		"from_row": 1,
		"to_row": 15,
		"has_next": true,
		"has_previous": false
	},
	"success": true
}
*/
