go.buildquery
=======================

Introduction
------------
Build query is a go module designed to allow users to create MySQL/MariaDb
SQL query strings programmatically.  

Installation
------------
To include this go module in your project simply: 

    go get https://github.com/dmaworsley/buildquery

Usuage Example
--------------

Below is a basic example based around a simple 2 table inner join query example sql if needed is below

**sql**

    CREATE DATABASE example;

    CREATE TABLE users (
        user_id INT NOT NULL AUTO_INCREMENT,
        username VARCHAR(64) NOT NULL UNIQUE,
        join_date DATE NOT NULL,
        status TINYINT NOT NULL,
        PRIMARY KEY(user_id)
    );
    
    INSERT INTO users VALUES(NULL, 'user_adam', DATE ( NOW() ), 1) ; 
    INSERT INTO users VALUES(NULL, 'user_bob', DATE ( NOW() ), 1) ;
    INSERT INTO users VALUES(NULL, 'user_charles', DATE ( NOW() ), 1) ;
    INSERT INTO users VALUES(NULL, 'user_david', DATE ( NOW() ), 1) ;
    INSERT INTO users VALUES(NULL, 'user_edward', DATE ( NOW() ), 1) ;
    INSERT INTO users VALUES(NULL, 'user_frederick', DATE ( NOW() ), 1) ;
    INSERT INTO users VALUES(NULL, 'user_gareth', DATE ( NOW() ), 1) ;
    INSERT INTO users VALUES(NULL, 'user_harold', DATE ( NOW() ), 1) ;
    INSERT INTO users VALUES(NULL, 'user_iain', DATE ( NOW() ), 1) ;
    
    CREATE TABLE user_details (
        user_id INT NOT NULL,
        `key` VARCHAR(64) NOT NULL,
        `value` TEXT NOT NULL,
        PRIMARY KEY(user_id, `key`)
    );
     
     INSERT INTO user_details VALUES(1, 'email', 'adam@email.com');
     INSERT INTO user_details VALUES(2, 'email', 'bob@email.com');
     INSERT INTO user_details VALUES(3, 'email', 'charles@email.com');
     INSERT INTO user_details VALUES(4, 'email', 'david@email.com');
     INSERT INTO user_details VALUES(5, 'email', 'edward@email.com');
     INSERT INTO user_details VALUES(6, 'email', 'frederick@email.com');
     INSERT INTO user_details VALUES(7, 'email', 'gareth@email.com');
     INSERT INTO user_details VALUES(8, 'email', 'harold@email.com');
     INSERT INTO user_details VALUES(9, 'email', 'iain@email.com');

**buildquery**

    //create a new Select
    query := Select()
    
    //which table in the from clause
    from_table := make(AliasItem)
    //AliasItem = map[string]string
    from_table["u"] = "users"
    
    query.From(from_table)
    
    columns := make(AliasItem)
    columns["user_id"] = "`u`." +  EscapeField("user_id")
    columns["username"] = "username"
    columns["join_date"] = "join_date"
    columns["status"] = EscapeField("status")
    
    query.Columns(columns)
    
    //lets join something
    
    join_table := make(AliasItem)
    join_table["ud"] = "user_details"
    
    join_columns := make(AliasItem)
    join_columns["email"] = "`ud`." +  EscapeField("value")
    
    
    onClause := "ud.user_id = u.user_id and ud.`key` = 'email'"
    joinType := "INNER"
    
    query.Join(join_table, onClause, join_columns, joinType)
    
    //lets add a less Or Equal than clause to the query 
    lessThanOrEqual := make(AliasItem)
    lessThanOrEqual["u.user_id"] = "5"
    
    query.LessThanEqualTo(lessThanOrEqual)
    
    //get our sql string 
    sql := query.BuildSelect()
    
    //assuming you include fmt
    fmt.Printf("%v\n", sql)
    
**Output**
    
    MariaDB [example]> SELECT username as `username`, join_date as `join_date`, `status` as `status`, `u`.`user_id` as `user_id`, `ud`.`value` as `email` FROM users as `u` INNER JOIN user_details as `ud` ON ud.user_id = u.user_id and ud.`key` = 'email' WHERE u.user_id <= '5' ;
    
    +--------------+------------+--------+---------+-------------------+
    | username     | join_date  | status | user_id | email             |
    +--------------+------------+--------+---------+-------------------+
    | user_adam    | 2020-10-23 |      1 |       1 | adam@email.com    |
    | user_bob     | 2020-10-23 |      1 |       2 | bob@email.com     |
    | user_charles | 2020-10-23 |      1 |       3 | charles@email.com |
    | user_david   | 2020-10-23 |      1 |       4 | david@email.com   |
    | user_edward  | 2020-10-23 |      1 |       5 | edward@email.com  |
    +--------------+------------+--------+---------+-------------------+
    5 rows in set (0.003 sec)
    
    