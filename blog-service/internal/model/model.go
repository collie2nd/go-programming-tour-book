package model

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         uint32 `json:"id"`          // id
	CreatedOn  uint32 `json:"created_on"`  // 创建时间
	CreatedBy  string `json:"created_by"`  // 创建人
	ModifiedOn uint32 `json:"modified_on"` // 修改时间
	ModifiedBy string `json:"modified_by"` // 修改人
	DeletedOn  uint32 `json:"deleted_on"`  // 删除时间
	DeletedBy  string `json:"deleted_by"`  // 删除人
	IsDel      uint8  `json:"is_del"`      // 是否删除，0为未删除，1为已删除
}

type Tag struct {
	*Model
	Name  string `json:"name"`  // 标签名称
	State uint8  `json:"state"` // 状态0为禁用，1为启用
}

func (model Tag) TableName() string {
	return "blog_tag"
}

type Article struct {
	*Model
	Title         string `json:"title"`           // 文章标题
	Desc          string `json:"desc"`            // 文章简述
	CoverImageUrl string `json:"cover_image_url"` // 封面图片地址
	Content       string `json:"content"`         // 文章内容
	State         uint8  `json:"state"`           // 状态0为禁用，1为启用
}

func (model Article) TableName() string {
	return "blog_article"
}

type ArticleTag struct {
	*Model
	ArticleId uint32 `json:"article_id"` // 文章ID
	TagId     uint32 `json:"tag_id"`     // 标签ID
}

func (model ArticleTag) TableName() string {
	return "blog_article_tag"
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
