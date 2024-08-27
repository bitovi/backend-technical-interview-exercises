const express = require("express");
const app = express();
const Sales = require("./models/sales");

app.get("/", (req, res) => {
  res.send("Status: OK");
});

/*
 * This endpoint is for retrieving sales data
 *     It should return an array of people and their sales for a given year
 *     Each person object will look like:
 *     {
 *         name: "Abdul Karim Khandker",
 *         sales: [117, 46, 82, 102, 119, 93, 123, 138, 91, 53, 55, 115]
 *     }
 */
app.get("/sales", (req, res) => {
  /*
   * Exercise 1 - Add a parameter for excluding the sales below a certain value
   *              This value should be required and should always be a Number
   */

  // parameter code goes here

  /*
   * Exercise 2 -  Retrieve all sales from the Sales model
   *              and filter them based on the parameter in Exercise 1.
   *              Any sales lower than the minimum should show up as 0
   *              in the array for that salesperson.
   */

  // sales retrieval and filter code goes here

  res.send("Sales")
});

app.listen(3002, () => {
  console.log("Exercise 1: server started, curl 'localhost:3002/sales'");
});
