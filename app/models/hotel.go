package models

import (
	"errors"

	"gorm.io/gorm"

	"github.com/st4rgaze/otaqu/config"
	"github.com/st4rgaze/otaqu/utils"
)

type Hotel struct {
	gorm.Model
	Name       string `gorm:"not null" json:"name"`
	Address    string `gorm:"not null" json:"address"`
	ImageUrl   string `gorm:"not null;unique" json:"image_url"`
	StarRating uint   `gorm:"not null" json:"star_rating"`
	Price      uint   `gorm:"not null" json:"price"`
}

type HotelFormatted struct {
	Name           string `gorm:"not null" json:"name"`
	Address        string `gorm:"not null" json:"address"`
	ImageUrl       string `gorm:"not null;unique" json:"image_url"`
	StarRating     uint   `gorm:"not null" json:"star_rating"`
	FormattedPrice string `gorm:"not null" json:"formatted_price"`
	Price          uint   `gorm:"not null" json:"price"`
}

// get all hotel
func GetAll() ([]HotelFormatted, error) {
	var hotels []Hotel
	err := config.DB.Find(&hotels).Error
	if err != nil {
		return nil, err
	}
	formattedHotels := make([]HotelFormatted, len(hotels))

	// Format the price
	for i, h := range hotels {
		formattedHotels[i] = HotelFormatted{
			Name:           h.Name,
			Address:        h.Address,
			ImageUrl:       h.ImageUrl,
			StarRating:     h.StarRating,
			FormattedPrice: utils.FormatPrice(h.Price),
			Price:          h.Price,
		}
	}

	return formattedHotels, nil
}

// create hotel
func (h *Hotel) Create() error {
	// check exist
	existingHotel, err := GetHotelByNameAndAddress(h.Name, h.Address)
	if err != nil {
		return err
	}

	if existingHotel != nil {
		// update if exist
		existingHotel.ImageUrl = h.ImageUrl
		existingHotel.StarRating = h.StarRating
		existingHotel.Price = h.Price
		err = config.DB.Save(existingHotel).Error
		if err != nil {
			return err
		}
	} else {
		// Create a new hotel if not
		err = config.DB.Create(h).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func GetHotelByNameAndAddress(name, address string) (*Hotel, error) {
	var hotel Hotel
	err := config.DB.Where("name = ? AND address = ?", name, address).First(&hotel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &hotel, nil
}

// get hotel by name
func GetHotelByName(name string) (*Hotel, error) {
	var hotel Hotel
	err := config.DB.Where("name = ?", name).First(&hotel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &hotel, nil
}

// get hotel by id
func GetHotelByID(id uint) (*Hotel, error) {
	var hotel Hotel
	err := config.DB.First(&hotel, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &hotel, nil
}

// Update hotel
func (h *Hotel) Update() error {
	err := config.DB.Save(h).Error
	if err != nil {
		return err
	}
	return nil
}
