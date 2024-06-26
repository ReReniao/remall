package serializer

import "re-mall/model"

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductId,
		ImgPath:   item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImg []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImg = append(productImg, product)
	}
	return
}
