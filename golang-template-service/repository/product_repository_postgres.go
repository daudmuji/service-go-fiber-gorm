package repository

import (
	"golang-template-service/model"
	"gorm.io/gorm"
)

type ProductRepositoryPostgres struct {
	Conn *gorm.DB
}

func NewProductRepositoryPostgres(conn *gorm.DB) *ProductRepositoryPostgres {
	return &ProductRepositoryPostgres{Conn: conn}
}

func (p *ProductRepositoryPostgres) CreateProduct(product *model.Product) (err error, result int64) {
	tx := p.Conn.Table("tbl_product").Create(product)
	err = tx.Error

	if err != nil {
		return err, result
	}

	affected := tx.RowsAffected
	if affected == 0 {
		return err, affected
	}

	return nil, affected

}

func (p *ProductRepositoryPostgres) GetListProducts() (err error, products []model.ProductQuery) {
	rows, err := p.Conn.Table("tbl_product").Find(&products).Rows()
	if err != nil {
		return err, nil
	}

	defer rows.Close()

	listProducts := make([]model.ProductQuery, 0)
	for rows.Next() {
		var product model.ProductQuery
		err = rows.Scan(
			&product.ID,
			&product.NamaBarang,
			&product.JumlahStokBarang,
			&product.NomorSeriBarang,
			&product.AdditionalInfo,
			&product.GambarBarang,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return err, nil
		}
		listProducts = append(listProducts, product)
	}

	return nil, listProducts
}

func (p *ProductRepositoryPostgres) GetDetailProductByNomorSeri(nomorSeri int) (err error, product model.ProductQuery) {
	row := p.Conn.Table("tbl_product").Where("nomor_seri_barang = ?", nomorSeri).First(&product).Row()
	err = row.Err()
	if err != nil {
		return err, product
	}

	err = row.Scan(
		&product.ID,
		&product.NamaBarang,
		&product.JumlahStokBarang,
		&product.NomorSeriBarang,
		&product.AdditionalInfo,
		&product.GambarBarang,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return err, product
	}

	return nil, product
}

func (p *ProductRepositoryPostgres) UpdateProductByNomorSeri(nomorSeri int, product *model.Product) (err error, result int64) {
	tx := p.Conn.Table("tbl_product").Where("nomor_seri_barang = ?", nomorSeri).Updates(product)
	err = tx.Error
	if err != nil {
		return err, result
	}

	affected := tx.RowsAffected
	if affected == 0 {
		return err, affected
	}
	return nil, affected
}

func (p *ProductRepositoryPostgres) DeleteProductByNomorSeri(nomorSeri int, product *model.Product) (err error, result int64) {
	tx := p.Conn.Table("tbl_product").Where("nomor_seri_barang = ?", nomorSeri).Delete(&product)
	err = tx.Error
	if err != nil {
		return err, result
	}

	affected := tx.RowsAffected
	if affected == 0 {
		return err, affected
	}
	return nil, affected
}
