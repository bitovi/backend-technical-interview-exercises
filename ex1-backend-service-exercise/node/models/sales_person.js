const records = require("./records.json")

function SalesPeople() {}

/*
 * Query the SalesPeople records
 */
SalesPeople.findAll = () => {
  return Promise.resolve(records);
};

module.exports = SalesPeople;
