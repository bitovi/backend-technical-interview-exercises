const salesData = [
  {
    name: "Abdul Karim Khandker",
    sales: [117, 46, 82, 102, 119, 93, 123, 138, 91, 53, 55, 115]
  },
  {
    name: "Adam Cortez",
    sales: [52, 55, 91, 107, 53, 158, 114, 158, 73, 106, 58, 152]
  },
  {
    name: "Alycia Lin",
    sales: [125, 113, 128, 78, 156, 59, 102, 64, 76, 35, 103, 38]
  }
];

function Sales() {}

/*
 * Query the Sales data
 */
Sales.findAll = () => {
  return Promise.resolve(salesData);
};

module.exports = Sales;
