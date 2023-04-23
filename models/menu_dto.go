package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShortMenu struct {
	Id                   primitive.ObjectID   `json:"id,omitempty" bson:"_id"`
	Name                 string               `json:"name,omitempty" validate:"required"`
	ThumbnailImage       string               `json:"thumbnailImage,omitempty" validate:"required"`
	FullPrice            int64                `json:"fullPrice,omitempty" validate:"required"`
	DiscountedPercent    int64                `json:"discountedPercent,omitempty" validate:"required"`
	DiscountedTimePeriod DiscountedTimePeriod `json:"discountedTimePeriod,omitempty" validate:"required"`
	Sold                 int64                `json:"sold,omitempty" validate:"required"`
	TotalInStock         int64                `json:"totalInStock,omitempty" validate:"required"`
}

type FullMenu struct {
	Id                   primitive.ObjectID   `json:"id,omitempty" bson:"_id"`
	Name                 string               `json:"name,omitempty" validate:"required"`
	ThumbnailImage       string               `json:"thumbnailImage,omitempty" validate:"required"`
	FullPrice            int64                `json:"fullPrice,omitempty" validate:"required"`
	DiscountedPercent    int64                `json:"discountedPercent,omitempty" validate:"required"`
	DiscountedTimePeriod DiscountedTimePeriod `json:"discountedTimePeriod,omitempty" validate:"required"`
	Sold                 int64                `json:"sold,omitempty" validate:"required"`
	TotalInStock         int64                `json:"totalInStock,omitempty" validate:"required"`
	LargeImage           string               `json:"largeImage,omitempty" validate:"required"`
	Options              []Option             `json:"options,omitempty" validate:"required"`
}

type DiscountedTimePeriod struct {
	Begin string `json:"begin,omitempty" validate:"required"`
	End   string `json:"end,omitempty" validate:"required"`
}

type Option struct {
	Label   string   `json:"label,omitempty" validate:"required"`
	Choices []Choice `json:"choices,omitempty" validate:"required"`
}

type Choice struct {
	Label string `json:"label,omitempty" validate:"required"`
}
