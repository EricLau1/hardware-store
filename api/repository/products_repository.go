package repository

import (
	"hardware-store/api/models"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductsRepository interface {
	Save(*models.Product) (*models.Product, error)
	Find(uint64) (*models.Product, error)
	FindAll() ([]*models.Product, error)
	Update(*models.Product) error
	Delete(uint64) error
	Count() (int64, error)
	Paginate(*Metadata) (*Pagination, error)
	Search(string) ([]*models.Product, error)
}

type productsRepositoryImpl struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *productsRepositoryImpl {
	return &productsRepositoryImpl{db}
}

func (r *productsRepositoryImpl) Save(product *models.Product) (*models.Product, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Product{}).Create(product).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return product, tx.Commit().Error
}

func (r *productsRepositoryImpl) Find(product_id uint64) (*models.Product, error) {
	product := &models.Product{}
	err := r.db.Debug().Model(&models.Product{}).Where("id = ?", product_id).Find(product).Error
	return product, err
}

func (r *productsRepositoryImpl) FindAll() ([]*models.Product, error) {
	products := []*models.Product{}
	err := r.db.Debug().Model(&models.Product{}).Find(&products).Error
	return products, err
}

func (r *productsRepositoryImpl) Update(product *models.Product) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"name":        product.Name,
		"price":       product.Price,
		"quantity":    product.Quantity,
		"status":      product.Status,
		"category_id": product.CategoryID,
		"updated_at":  time.Now(),
	}

	err := tx.Debug().Model(&models.Product{}).Where("id = ?", product.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *productsRepositoryImpl) Delete(product_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Product{}).Where("id = ?", product_id).Delete(&models.Product{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *productsRepositoryImpl) Count() (int64, error) {
	var c int64
	err := r.db.Debug().Model(&models.Product{}).Count(&c).Error
	return c, err
}

func (r *productsRepositoryImpl) Paginate(meta *Metadata) (*Pagination, error) {
	products := []*models.Product{}

	err := r.db.Debug().
		Model(&models.Product{}).
		Offset(meta.Offset).
		Limit(meta.Limit).
		Find(&products).Error

	return &Pagination{
		Elements: products,
		Metadata: meta,
	}, err
}

func (r *productsRepositoryImpl) Search(search string) ([]*models.Product, error) {
	products := []*models.Product{}

	err := r.db.Debug().
		Model(&models.Product{}).
		Where("name like ?", "%"+search+"%").
		Find(&products).Error

	return products, err
}
