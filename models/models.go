package models

import (
	"encoding/json"
	"strings"
)

type Response struct {
	Status  interface{} `json:"status"`
	Message string      `json:"message"`
	Time    interface{} `json:"time"`
	Data    interface{} `json:"data"`
}

type Config struct {
	Mysqlconnstring string `yaml:"Mysqlconnstring"`
	Server_port     string `yaml:"Server_port"`
}

type Division struct {
	ID     uint   `json:"_id"`
	Name   string `json:"name" gorm:"column:name"`
	BnName string `json:"bn_name" gorm:"column:bn_name"`
	Url    string `json:"url" gorm:"column:url"`
}

func (d *Division) MarshalJSON() ([]byte, error) {
	type Alias Division
	return json.Marshal(&struct {
		ID string `json:"_id"`
		*Alias
	}{
		ID:    strings.ToLower(d.Name),
		Alias: (*Alias)(d),
	})
}

type District struct {
	ID uint `json:"_id" gorm:"primary_key"`
	//DivisionID int    `json:"division_id" gorm:"column:division_id"`
	Name   string `json:"name" gorm:"column:name"`
	BnName string `json:"bn_name" gorm:"column:bn_name"`
	Lat    string `json:"lat" gorm:"column:lat"`
	Lon    string `json:"lon" gorm:"column:lon"`
	Url    string `json:"url" gorm:"column:url"`
}

func (d *District) MarshalJSON() ([]byte, error) {
	type Alias District
	return json.Marshal(&struct {
		ID string `json:"_id"`
		*Alias
	}{
		ID:    strings.ToLower(d.Name),
		Alias: (*Alias)(d),
	})
}

type Upazila struct {
	ID uint `json:"_id" gorm:"primary_key"`
	//	DistrictID int    `json:"district_id" gorm:"column:district_id"`
	Name   string `json:"name" gorm:"column:name"`
	BnName string `json:"bn_name" gorm:"column:bn_name"`
	Url    string `json:"url" gorm:"column:url"`
}

func (d *Upazila) MarshalJSON() ([]byte, error) {
	type Alias Upazila
	return json.Marshal(&struct {
		ID string `json:"_id"`
		*Alias
	}{
		ID:    strings.ToLower(d.Name),
		Alias: (*Alias)(d),
	})
}

type Union struct {
	ID uint `json:"_id" gorm:"primary_key"`
	//	UpazillaID int    `json:"upazilla_id" gorm:"column:upazilla_id"`
	Name   string `json:"name" gorm:"column:name"`
	BnName string `json:"bn_name" gorm:"column:bn_name"`
	Url    string `json:"url" gorm:"column:url"`
}

func (d *Union) MarshalJSON() ([]byte, error) {
	type Alias Union
	return json.Marshal(&struct {
		ID string `json:"_id"`
		*Alias
	}{
		ID:    strings.ToLower(d.Name),
		Alias: (*Alias)(d),
	})
}
