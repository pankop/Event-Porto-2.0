package seeds

import (
	"encoding/json"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
	"io"
	"os"
)

func ListProvinceSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/province.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listProvince []entity.Province
	if err := json.Unmarshal(jsonData, &listProvince); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Province{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Province{}); err != nil {
			return err
		}
	}

	for _, data := range listProvince {
		var province entity.Province

		isData := db.Find(&province, "id = ?", data.ID).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
