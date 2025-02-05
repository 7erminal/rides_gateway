package responses

type Item_images struct {
	ItemImageId int64
	ItemId      int64
	ImagePath   string
	IsDefault   int
}

type ItemImageOriResponseDTO struct {
	StatusCode int
	Value      *string
	StatusDesc string
}

type ItemImagesOriResponseDTO struct {
	StatusCode int
	ItemImages *[]Item_images
	StatusDesc string
}

type ItemImageResponseDTO struct {
	Success    bool
	Result     *string
	StatusDesc string
}

type ItemImagesResponseDTO struct {
	Success    bool
	Result     *[]Item_images
	StatusDesc string
}
