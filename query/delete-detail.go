package query

import (
	"fmt"
	"github.com/seichewarning/mysql-restful-server/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDetail(c *gin.Context) {
	if AdminTableFilter(c) {
		deleteDetail(c)
	}
}

func deleteDetail(c *gin.Context) {
	table := c.Param("table")
	id := c.Param("id")
	sqlstring := "delete  from " + table + " where id = '" + id + "' ;"
	result, err := connection.GetConnection().Exec(sqlstring)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNoContent, gin.H{
			"message": fmt.Sprintf("%s not found", id),
		})
	} else {
		ids, erro := result.RowsAffected()
		if erro != nil || ids == 0 {
			fmt.Println(erro)
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Failed deleted user: %s ", id),
			})
		} else {

			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Successfully deleted user: %s", id),
			})
		}
	}

}
