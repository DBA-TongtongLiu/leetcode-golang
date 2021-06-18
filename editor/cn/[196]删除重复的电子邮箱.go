//编写一个 SQL 查询，来删除 Person 表中所有重复的电子邮箱，重复的邮箱里只保留 Id 最小 的那个。 
//
// +----+------------------+
//| Id | Email            |
//+----+------------------+
//| 1  | john@example.com |
//| 2  | bob@example.com  |
//| 3  | john@example.com |
//+----+------------------+
//Id 是这个表的主键。
// 
//
// 例如，在运行你的查询语句之后，上面的 Person 表应返回以下几行: 
//
// +----+------------------+
//| Id | Email            |
//+----+------------------+
//| 1  | john@example.com |
//| 2  | bob@example.com  |
//+----+------------------+
// 
//
// 
//
// 提示： 
//
// 
// 执行 SQL 之后，输出是整个 Person 表。 
// 使用 delete 语句。 
// 
// 👍 368 👎 0


//There is no code of Go type for this problem

delete from Person where Id not in
(select min(Id) as Id from Person group by Email having count(1) > 1

union all

select Id from Person group by Email having count(1) = 1)