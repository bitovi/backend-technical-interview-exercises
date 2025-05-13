using Microsoft.AspNetCore.Mvc;
using ServiceExercise.Models;
using ServiceExercise.Repositories;

namespace ServiceExercise.Controllers;

[ApiController]
[Route("[controller]")]
public class SalesPersonController(ILogger<SalesPersonController> logger, ISalesPersonRepository repository)
    : ControllerBase
{
    private readonly ILogger<SalesPersonController> _logger = logger;
    private readonly ISalesPersonRepository _repository = repository;

    /*
     * GET /salesperson is for retrieving a list of sales people and some aggregate sales data for each month.
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
    [HttpGet()]
    public ActionResult<IList<SalesPerson>> Get()
    {
        /*
         * Step 1 - Add new a parameter for ignoring the months that are below a certain value.
         *          This value should be required and should have its type validated.
         *
         * Step 2 - Fetch all records from the SalesPerson repository and modify the monthly sales data based on the new parameter.
         *          Any sales months lower than the parameter value should have `amount` set to `0`.
         *
         * Step 3 - Return the modified data in the response.
         */
        IList<SalesPerson> salesPeople = _repository.GetSalesPersons();
        return Ok(salesPeople);
    }
}