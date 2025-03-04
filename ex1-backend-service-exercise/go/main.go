package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "\n")
	})

	/*
	 * GET /sales_person is for retrieving a list of sales people and some aggregate sales data for each month.
	 * Each person object will look like:
	 *
	 * {
	 *   "name": "Abdul Karim Khandker",
	 *   "sales": [
	 *     {
	 *       "month": "1",
	 *       "amount": 117
	 *     },
	 *     ...
	 *     {
	 *       "month": "12",
	 *       "amount": 117
	 *     },
	 * }
	 */
	r.GET("/sales_person", func(c *gin.Context) {
		/*
		 * Step 1 - Add new a parameter for ignoring the months that are below a certain value.
		 *          This value should be required and should have its type validated.
		 *
		 * Step 2 - Fetch all records from the SalesPerson model and modify the monthly sales data based on the new parameter.
		 *          Any sales months lower than the parameter value should have `amount` set to `0`.
		 *
		 * Step 3 - Return the modified data in the response.
		 */
	})

	port := ":3002"
	r.Run(port)
}
