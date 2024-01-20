Drop TABLE IF EXISTS content; 
CREATE TABLE content (
    id INT AUTO_INCREMENT NOT null,
    post varchar(255),
    author varchar(255),
    primary key(id)
);

insert into content 
(post , author)
values
("This is demo post details" , "joy"),
("This is demo post details 2 " , "joy2"),
("This is demo post details 3 " , "joy23");

