use todo_list;

CREATE TABLE todo_items (
    id VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    image JSON,
    description TEXT,
    status ENUM('doing', 'done', 'deleted') DEFAULT 'doing',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO todo_items (id, title, image, description, status, created_at, updated_at) 
VALUES 
('1', 'Buy groceries', NULL, 'Buy milk, eggs, and bread', 'doing', '2024-06-03 10:00:00', '2024-06-03 10:00:00'),
('2', 'Read book', NULL, 'Read "To Kill a Mockingbird"', 'done', '2024-06-01 15:30:00', '2024-06-02 09:00:00'),
('3', 'Go for a run', NULL, 'Run 5 kilometers in the park', 'doing', '2024-06-02 08:00:00', '2024-06-03 08:30:00'),
('4', 'Call mom', NULL, 'Call mom to check in and say hello', 'done', '2024-06-01 10:45:00', '2024-06-01 11:00:00'),
('5', 'COding', NULL, 'Golang, react, db', 'deleted', '2024-06-01 10:45:00', '2024-06-01 11:00:00');

explain
select * from todo_items where status = 'done';

CREATE INDEX status_index ON todo_items (status);

CREATE TABLE todo_user_like_items (
  user_id VARCHAR(50) NOT NULL,
  item_id VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, item_id)
);

INSERT INTO todo_user_like_items (user_id, item_id)
VALUES
  ('1', '1'),  
  ('1', '2'),
  ('1', '3'),  
  ('2', '1'),  
  ('3', '1');

 explain
 select * from todo_user_like_items WHERE user_id = 1 and item_id = 1;
 