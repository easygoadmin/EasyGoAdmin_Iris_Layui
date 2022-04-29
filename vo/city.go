package vo

import "easygoadmin/model"

type CityInfoVo struct {
	model.City
	HaveChild bool `json:"haveChild"`
}
