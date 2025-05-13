using System.Diagnostics;
using ServiceExercise.Repositories;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddSingleton<ISalesPersonRepository, SalesPersonRepository>();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseAuthorization();

app.MapControllers();

app.Logger.Log(LogLevel.Information,"Exercise 1 server started => curl http://localhost:5254/salesperson");
app.Run();