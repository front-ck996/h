package csy_test

import (
	"github.com/front-ck996/csy"
	"testing"
)

// ThemeStartTemplate 默认的主题文件列表
type ThemeStartTemplate struct {
	//Model
	// 上传的文件保存位置
	File    string `json:"file" gorm:"column:file;size:255;"`
	OldFile string `json:"old_file" gorm:"column:old_file;size:255;"`
	// 上传的文件名
	Name string `json:"name" gorm:"size:255;comment:上传的文件名"`
	// 备注
	Remark       string `json:"remark" gorm:"size:0;comment:备注"`
	TestField    string `json:"test_field"`
	Extractor    int    `json:"extractor" gorm:"type:tinyint;comment:解压完成"`
	ExtractorErr string `json:"extractor_err" gorm:"size:0;comment:解压错误信息"`
	// 单个文件详情
	//ThemeDefaultFileListItem []ThemeStartTemplateItem `gorm:"foreignKey:f_id"`
	//ThemeProject             []*ThemeProject          `json:"theme_start_template" gorm:"many2many:theme_m2m_project_start;comment:默认主题"`
	//ControlBy
	//ModelTime
}

func (ThemeStartTemplate) TableName() string {
	return "theme_start_template"
}

func TestParseGormStruct(t *testing.T) {
	csy.ParseGormStructSchema(ThemeStartTemplate{})
	csy.ParseGormStructField(ThemeStartTemplate{})
}
