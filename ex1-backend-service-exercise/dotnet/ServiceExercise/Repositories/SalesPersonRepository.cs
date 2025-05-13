using System.Reflection;
using Newtonsoft.Json;
using ServiceExercise.Models;

namespace ServiceExercise.Repositories;

public class SalesPersonRepository : ISalesPersonRepository
{
    private readonly List<SalesPerson> _salesPersons = new List<SalesPerson>();
    public SalesPersonRepository()
    {
        var assembly = Assembly.GetExecutingAssembly();
        var resourceName = "ServiceExercise.records.json";

        using Stream? stream = assembly.GetManifestResourceStream(resourceName);
        if (stream != null)
        {
            using (StreamReader reader = new StreamReader(stream))
            {
                string json = reader.ReadToEnd();
                _salesPersons = JsonConvert.DeserializeObject<List<SalesPerson>>(json) ?? [];
            }
        }
    }

    public IList<SalesPerson> GetSalesPersons()
    {
        return _salesPersons.ToList();
    }
}