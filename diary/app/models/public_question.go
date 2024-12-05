package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// think about how to show on post: use infinite scroll
func GetPublicQuestion(db *gorm.DB, lastID int, pageSize int) ([]PublicQuestion, error) {
	var p []PublicQuestion
	err := db.Where("id > ?", lastID).Limit(pageSize).Find(&p).Error
	return p, err
}

func CreatePublicQuestion(db *gorm.DB, p *PublicQuestion) error {
	p.CreateAt = time.Now()
	//p.UpdateAt = time.Now()
	if err := db.Create(p).Error; err != nil {
		return err
	}
	if p.ID == 0 {
		if err := db.Last(p).Error; err != nil {
			return err
		}
		log.Printf("Manual ID retrieval: New public question created with ID: %d", p.ID)
	}
	return nil
}

/*
func UpdatePublicQuestion(db *gorm.DB, p *PublicQuestion) (*PublicQuestion, error) {

	// Set the updated timestamp
	p.UpdateAt = time.Now()

	// Perform the update operation
	if err := db.Model(&PublicQuestion{}).
		Where("id = ?", p.ID).
		Updates(map[string]interface{}{
			"question":  p.Title,
			"content":   p.Content,
			"update_at": p.UpdateAt,
		}).Error; err != nil {
		return nil, err
	}

	// Return the updated PublicQuestion
	return p, nil
}
*/
