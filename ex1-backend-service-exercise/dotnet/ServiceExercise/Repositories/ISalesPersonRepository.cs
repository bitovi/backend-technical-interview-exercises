using ServiceExercise.Models;

namespace ServiceExercise.Repositories;

public interface ISalesPersonRepository
{
    public IList<SalesPerson> GetSalesPersons();
}