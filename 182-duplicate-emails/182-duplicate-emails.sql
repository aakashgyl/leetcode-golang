# Write your MySQL query statement below
select email as Email
from Person
GROUP BY email
Having count(email) > 1;