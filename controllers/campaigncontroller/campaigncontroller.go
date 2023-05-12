package campaigncontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Dvizio/BackendAPI/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var campaigns []models.Campaign

	models.DB.Find(&campaigns)
	c.JSON(http.StatusOK, gin.H{"campaigns": campaigns})

}

func Show(c *gin.Context) {
	var campaign models.Campaign
	id := c.Param("id")

	if err := models.DB.First(&campaign, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Campaign": campaign})
}

func MaxClickThrough(c *gin.Context) {
	var campaign models.Campaign

	if err := models.DB.Order("click_through DESC").First(&campaign).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Campaign": campaign})
}

func MaxConversion(c *gin.Context) {
	var campaign models.Campaign

	if err := models.DB.Order("conversion DESC").First(&campaign).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Campaign": campaign})
}

func MaxNilaiAkhir(c *gin.Context) {
	var campaign models.Campaign

	if err := models.DB.Order("nilai_akhir DESC").First(&campaign).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Campaign": campaign})
}

func Create(c *gin.Context) {
	var campaign models.Campaign

	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&campaign)
	c.JSON(http.StatusOK, gin.H{"campaign": campaign})
}

// func CreateNilaiAkhir(c *gin.Context) {
// 	var campaign models.Campaign

// 	if err := c.ShouldBindJSON(&campaign); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	models.DB.Exec(`INSERT INTO campaigns(nilai_akhir)
// 	SELECT campaigns.click_through * campaigns.conversion
// 	FROM campaigns`)
// 	c.JSON(http.StatusOK, gin.H{"campaign": campaign})
// }

func Update(c *gin.Context) {
	var campaign models.Campaign
	id := c.Param("id")

	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&campaign).Where("id = ?", id).Updates(&campaign).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate Campaign"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var campaign models.Campaign

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&campaign, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus Campaign"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
