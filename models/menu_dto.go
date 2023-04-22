package models

type ShortMenu struct {
	Id                   string `json:"id" db:"id"`
	Name                 string `json:"name" db:"name"`
	ThumbnailImage       string `json:"thumbnailImage" db:"thumbnailImage"`
	FullPrice            int64  `json:"fullPrice" db:"fullPrice"`
	DiscountedPercent    int64  `json:"discountedPercent" db:"discountedPercent"`
	DiscountedTimePeriod struct {
		Begin string `json:"begin" db:"begin"`
		End   string `json:"end" db:"end"`
	} `json:"discountedTimePeriod" db:"discountedTimePeriod"`
	Sold         int64 `json:"sold" db:"sold"`
	TotalInStock int64 `json:"totalInStock" db:"totalInStock"`
}

type FullMenu struct {
	Id                   string `json:"id" db:"id"`
	Name                 string `json:"name" db:"name"`
	ThumbnailImage       string `json:"thumbnailImage" db:"thumbnailImage"`
	FullPrice            int64  `json:"fullPrice" db:"fullPrice"`
	DiscountedPercent    int64  `json:"discountedPercent" db:"discountedPercent"`
	DiscountedTimePeriod struct {
		Begin string `json:"begin" db:"begin"`
		End   string `json:"end" db:"end"`
	} `json:"discountedTimePeriod" db:"discountedTimePeriod"`
	Sold         int64  `json:"sold" db:"sold"`
	TotalInStock int64  `json:"totalInStock" db:"totalInStock"`
	LargeImage   string `json:"largeImage" db:"largeImage"`
	Options      []struct {
		Label   string `json:"label" db:"label"`
		Choices []struct {
			Label string `json:"label" db:"label"`
		} `json:"choices" db:"choices"`
	} `json:"options" db:"options"`
}
