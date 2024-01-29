# Task Management System


# Database Used 
MySQL

# Database Scema

taskDB -- accessible to task microservice
```
CREATE SCHEMA `taskDB` ;

CREATE TABLE `tasks` (
  `task_id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(45) NOT NULL,
  `description` varchar(45) DEFAULT NULL,
  `priority` int NOT NULL,
  `completed` tinyint NOT NULL DEFAULT '0',
  `assigned_to` int NOT NULL,
  `assigned_by` int NOT NULL,
  `due_date` date NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`task_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


```

userDB -- accessible to user microservice
```
CREATE SCHEMA `userDB` ;

CREATE TABLE `users` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `type` varchar(45) NOT NULL,
  `created_by` varchar(45) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

```

# Messaging Service Used for processing Events
# RabbitMQ

# Admin's API

# API to create User
```
curl --location 'localhost:8080/user/admin/create' \
--header 'Content-Type: application/json' \
--data '{
    "username":"user5",
    "email":"abhishek.amt",
    "type":"default",
    "createdBy":1
}'

```

# API to Delete User
```
curl --location --request DELETE 'localhost:8080/user/admin/delete?userId=3'

```

# API to Create Task
```
curl --location 'localhost:8081/task/create' \
--header 'Content-Type: application/json' \
--data '{
    "title":"abc4",
    "description":"dfcaf",
    "priority":3,
    "dueDate":"2024-02-12",
    "assignedTo":4,
    "assignedBy":1 

}'

```

# API to Update Task
```
curl --location --request PUT 'localhost:8081/task/admin/update' \
--header 'Content-Type: application/json' \
--data '{
    "title":"abcde",
    "dueDate": "2024-02-16",
    "taskId":3,
    "priority":2
}'

```

# Default User's API

# API to get User Detail

```
curl --location 'localhost:8080/user/get?userId=1'

```


# API to mark task completed 
```
curl --location --request PUT 'localhost:8081/task/complete?taskId=1' \
--data ''

```

# API to get all the tasks related to admin or user 
Description : If isAdmin is true then this api will return all the task that is assigned by admin in descending sorting order by sortBy key else it will return all the task which is assigned to user with the requested user id.

```
curl --location 'localhost:8081/task/get?isAdmin=true&userId=1&sortBy=dueDate'

```




