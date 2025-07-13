package utils

import (
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FilterSearch(c *gin.Context, textField, dateField string) bson.M {
	filter := bson.M{}

	// Filtro por texto
	text := c.Query(textField)
	if text != "" {
		filter[textField] = bson.M{"$regex": text, "$options": "i"}
	}
	// Filtro por rango de fechas
	from := c.Query("from")
	to := c.Query("to")
	dateFilter := bson.M{}

	if from != "" {
        if fromTime, err := time.Parse("2006-01-02", from); err == nil {
            dateFilter["$gte"] = fromTime
        }
    }
    if to != "" {
        if toTime, err := time.Parse("2006-01-02", to); err == nil {
            dateFilter["$lte"] = toTime
        }
    }
    if len(dateFilter) > 0 {
        filter[dateField] = dateFilter
    }

    return filter

}