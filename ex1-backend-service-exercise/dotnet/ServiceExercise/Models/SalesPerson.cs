namespace ServiceExercise.Models;

public class SalesPerson
{
    public required string Name { get; set; }
    public required IList<Sales> Sales { get; set; }
}