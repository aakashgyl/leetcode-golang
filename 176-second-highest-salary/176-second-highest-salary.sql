# Write your MySQL query statement below
SELECT max(salary) as SecondHighestSalary
FROM Employee as e
WHERE salary not in (
    select max(salary) 
    from Employee
);
