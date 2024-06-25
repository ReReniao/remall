package serializer

import (
	"context"
	"re-mall/conf"
	"re-mall/dao"
	"re-mall/model"
)

type Cart struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_Id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created__at"`
	Num           int    `json:"num"`
	Name          string `json:"name"`
	MaxNum        int    `json:"max_num"`
	ImgPath       string `json:"img_path"`
	Check         bool   `json:"check"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		Id:            cart.ID,
		UserId:        cart.UserId,
		ProductId:     cart.ProductId,
		CreatedAt:     cart.CreatedAt.Unix(),
		Num:           int(cart.Num),
		MaxNum:        int(cart.MaxNum),
		ImgPath:       conf.Host + conf.HttpPort + conf.AvatarPath + product.ImgPath,
		Check:         cart.Check,
		DiscountPrice: product.DiscountPrice,
		Name:          product.Name,
		Price:         product.Price,
		BossId:        boss.ID,
		BossName:      boss.UserName,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (carts []Cart) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.BossId)
		if err != nil {
			continue
		}
		cart := BuildCart(item, product, boss)
		carts = append(carts, cart)
	}
	return
}
