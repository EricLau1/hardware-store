package repository

import (
	"net/http"
	"strconv"
)

type PaginationBuilderRepository interface {
	BuildProductsMetadata(*http.Request) (*Metadata, error)
}

type paginationBuilderRepositoryImpl struct {
	productsRepository ProductsRepository
}

type Pagination struct {
	Elements interface{} `json:"elements"`
	Metadata *Metadata   `json:"metadata"`
}

type Metadata struct {
	Total  int64 `json:"total"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Page   int   `json:"page"`
	Pages  int   `json:"pages"`
}

func NewPaginationBuilderRepository(productsRepository ProductsRepository) *paginationBuilderRepositoryImpl {
	return &paginationBuilderRepositoryImpl{productsRepository}
}

const (
	p_page  = "page"
	p_limit = "limit"

	default_page  = 1
	default_limit = 10
)

// return page, limit, offset,
func (p *paginationBuilderRepositoryImpl) BuildPagination(r *http.Request) (int, int, int) {
	params := r.URL.Query()

	page, _ := strconv.Atoi(params.Get(p_page))
	if page < 1 {
		page = default_page
	}

	limit, _ := strconv.Atoi(params.Get(p_limit))
	if limit < 1 {
		limit = default_limit
	}

	offset := (limit * page) - limit

	return page, limit, offset
}

func (p *paginationBuilderRepositoryImpl) BuildPages(elems int64, limit int) int {
	var pages int
	total := int(elems)
	if pages = (total / limit); (total % limit) != 0 {
		pages++
	}
	return pages
}

func (p *paginationBuilderRepositoryImpl) BuildProductsMetadata(r *http.Request) (*Metadata, error) {
	page, limit, offset := p.BuildPagination(r)

	total, err := p.productsRepository.Count()
	if err != nil {
		return nil, err
	}

	return &Metadata{
		Total:  total,
		Limit:  limit,
		Offset: offset,
		Page:   page,
		Pages:  p.BuildPages(total, limit),
	}, nil
}
