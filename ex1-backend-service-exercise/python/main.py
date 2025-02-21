from fastapi import FastAPI
import json

with open('records.json', 'r') as file:
    records = json.load(file)

app = FastAPI()

# GET /sales_person is for retrieving a list of sales people and some aggregate sales data for each month.
# Each person object will look like:
#
# {
#   "name": "Abdul Karim Khandker",
#   "sales": [
#     {
#       "month": "1",
#       "amount": 117
#     },
#     ...
#     {
#       "month": "12",
#       "amount": 117
#     },
# }
@app.get("/sales_person")
def get_sales_people():
#     Step 1 - Add new a parameter for ignoring the months that are below a certain value.
#              This value should be required and should have its type validated.
#
#     Step 2 - Fetch all records from the SalesPerson model and modify the monthly sales data based on the new parameter.
#              Any sales months lower than the parameter value should have `amount` set to `0`.
#
#     Step 3 - Return the modified data in the response.
    return []