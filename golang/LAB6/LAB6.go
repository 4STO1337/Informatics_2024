using System;

public struct Employee
{
    private string name;
    private DateTime hireDate;
    private string position;

    public Employee(string name, DateTime hireDate, string position)
    {
        this.name = name;
        this.hireDate = hireDate;
        this.position = position;
    }

    public string GetName()
    {
        return name;
    }

    public void SetName(string name)
    {
        this.name = name;
    }

    public DateTime GetHireDate()
    {
        return hireDate;
    }

    public void SetHireDate(DateTime hireDate)
    {
        this.hireDate = hireDate;
    }

    public string GetPosition()
    {
        return position;
    }

    public void SetPosition(string position)
    {
        this.position = position;
    }
}
public class Program
{
    public static void Main(string[] args)
    {
        Employee employee = new Employee("Иван Иванов", new DateTime(2020, 5, 1), "Менеджер");
        
        Console.WriteLine($"Имя: {employee.GetName()}, Дата приема: {employee.GetHireDate().ToShortDateString()}, Должность: {employee.GetPosition()}");
        
        employee.SetPosition("Старший менеджер");
        Console.WriteLine($"Обновленная должность: {employee.GetPosition()}");
    }
}