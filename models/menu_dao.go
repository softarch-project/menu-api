package models

type ShortMenuDAO struct {
	Id                     string `db:"id"`
	Name                   string `db:"name"`
	ThumbnailImage         string `db:"thumbnailImage"`
	FullPrice              int64  `db:"fullPrice"`
	DiscountedPercent      int64  `db:"discountedPercent"`
	DiscountedTimePeriodId int64  `db:"discountedTimePeriodId"`
	Sold                   int64  `db:"sold"`
	TotalInStock           int64  `db:"totalInStock"`
}

type FullMenuDAO struct {
	Id                     string   `db:"id"`
	Name                   string   `db:"name"`
	ThumbnailImage         string   `db:"thumbnailImage"`
	FullPrice              int64    `db:"fullPrice"`
	DiscountedPercent      int64    `db:"discountedPercent"`
	DiscountedTimePeriodId int64    `db:"discountedTimePeriodId"`
	Sold                   int64    `db:"sold"`
	TotalInStock           int64    `db:"totalInStock"`
	LargeImage             string   `db:"largeImage"`
	OptionsId              []string `db:"optionsId"`
}

type DiscountedTimePeriodDAO struct {
	Id    string `db:"id"`
	Begin string `db:"begin"`
	End   string `db:"end"`
}

type OptionsDAO struct {
	Id        string   `db:"id"`
	Label     string   `db:"label"`
	ChoicesId []string `db:"choicesId"`
}

type ChoicesDAO struct {
	Id      string   `db:"id"`
	LabelId []string `db:"labelId"`
}

type Label struct {
	Id        string `db:"id"`
	LabelText string `db:"labelText"`
}
