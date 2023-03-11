-- drop table posts CASCADE if exists posts;    失效语法
-- drop table comments CASCADE if exists;

-- 代码清单6-13 创建两个相关联的表
drop table if exists posts CASCADE;     -- 正确语法
drop table if exists comments CASCADE;  -- 正确语法

create table posts (
    id      serial primary key,
    content text,
    author  varchar(255)
);

create table comments (
    id      serial primary key,
    content text,
    author  varchar(255),
    post_id integer references posts(id)
);