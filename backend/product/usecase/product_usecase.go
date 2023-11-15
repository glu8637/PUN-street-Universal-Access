package usecase

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"
)

type productUsecase struct {
	productRepo domain.ProductRepo
}

func NewProductUsecase(productRepo domain.ProductRepo) domain.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (pu *productUsecase) GetByID(ctx context.Context, id int64) (*[]swagger.ProductInfo, error) {
	products, err := pu.productRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return products, nil
}

func (pu *productUsecase) AddByStoreId(ctx context.Context, id int64, product *swagger.ProductInfo) error {
	err := pu.productRepo.AddByStoreId(ctx, id, product)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (pu *productUsecase) UpdateById(ctx context.Context, StoreId int64, productId int64, product *swagger.ProductInfo) error {
	err := pu.productRepo.UpdateById(ctx, StoreId, productId, product)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
