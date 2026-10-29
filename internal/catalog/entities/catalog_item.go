package entities

type CategoryItem struct {
	BaseEntity
	ShortDescription string
	FullDescription  string
	ImageURL         string
	Brand            *Brand
	Category         *Category
	Price            float64
}
