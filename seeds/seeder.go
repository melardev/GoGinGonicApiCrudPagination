package seeds

import (
	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	"github.com/melardev/GoGinGonicApiGormCrudPagination/models"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {

	return rand.Intn(max-min) + min
}
func Seed(db *gorm.DB) {
	fake.Seed(time.Now().Unix())
	var countTodos int
	db.Model(&models.Todo{}).Count(&countTodos)
	todosToSeed := 43
	todosToSeed -= countTodos
	completed := true
	if randomInt(0, 3)%2 == 0 {
		completed = false
	}

	for i := 0; i < todosToSeed; i++ {
		db.Create(&models.Todo{
			Title:       fake.WordsN(randomInt(2, 4)),
			Description: fake.SentencesN(randomInt(1, 3)),
			Completed:   completed,
		})
	}
}
