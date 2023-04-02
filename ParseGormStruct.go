package csy

import (
	"github.com/front-ck996/csy/gorm_reorder"
	"gorm.io/gorm/schema"
)

func ParseGormStructField(s interface{}) (string, error) {
	cfg := gorm_reorder.Config{
		AutoAdd:       true,
		TablePrefix:   "",
		SingularTable: true,
	}
	reorder := gorm_reorder.NewReorder(cfg).AddModel([]interface{}{s}).Parser()
	b, err := gorm_reorder.MarshalSchema(reorder.GetSchemas())
	if err != nil {
		return "", err
	}
	return string(b), nil
}
func ParseGormStructSchema(s interface{}) []*schema.Schema {
	cfg := gorm_reorder.Config{
		AutoAdd:       true,
		TablePrefix:   "",
		SingularTable: true,
	}
	reorder := gorm_reorder.NewReorder(cfg).AddModel([]interface{}{s}).Parser()
	return reorder.GetSchemas()
}
