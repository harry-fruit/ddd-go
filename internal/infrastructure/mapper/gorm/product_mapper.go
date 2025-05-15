package gormmapper

import (
	"github.com/harry-fruit/ddd-go/internal/domain/entity"
	gormmodel "github.com/harry-fruit/ddd-go/internal/infrastructure/model/gorm"
	sharedinfra "github.com/harry-fruit/ddd-go/internal/shared/infrastructure"
)

type ProductMapper struct{}

func NewProductMapper() sharedinfra.Mapper[entity.Product, gormmodel.ProductModel] {
	return &ProductMapper{}
}

func (pm *ProductMapper) ToDomain(model gormmodel.ProductModel) entity.Product {
	return entity.Product{
		ID:             model.ID,
		UniqueKey:      model.UniqueKey,
		Name:           model.Name,
		Description:    model.Description,
		Price:          model.Price,
		RemainingStock: model.RemainingStock,
	}
}

func (pm *ProductMapper) ToDB(entity entity.Product) gormmodel.ProductModel {
	return gormmodel.ProductModel{
		ID:             entity.ID,
		UniqueKey:      entity.UniqueKey,
		Name:           entity.Name,
		Description:    entity.Description,
		Price:          entity.Price,
		RemainingStock: entity.RemainingStock,
	}
}
