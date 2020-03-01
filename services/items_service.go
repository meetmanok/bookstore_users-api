package services

var (
	ItemsService itemsService = itemsService{}
)

type itemsService struct {
}

type itemsServiceInterface interface {
	GetItem()
	SaveItem()
}

func (is *itemsService) GetItem() {

}

func (is *itemsService) SaveItem() {

}