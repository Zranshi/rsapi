create table auth_user
(
    id        bigint auto_increment comment 'user table id',
    email     varchar(32) default ''      not null comment 'user email, unique for visible user',
    `key`     varchar(64) default ''      not null comment 'hashed key with private salt',
    name      varchar(32) default ''      not null comment 'user display name',
    create_at datetime    default (now()) not null comment 'create time',
    update_at datetime    default (now()) not null on update CURRENT_TIMESTAMP comment 'update time',
    primary key (id),
    constraint unique_email
        unique (email)
) comment 'user table';


create table auth_role
(
    id        bigint auto_increment comment 'role table id',
    name      varchar(32) default ''      not null comment 'role name',
    create_at datetime    default (now()) not null comment 'create time',
    update_at datetime    default (now()) not null on update CURRENT_TIMESTAMP comment 'update time',
    primary key (id),
    constraint unique_name
        unique (name)
) comment 'role table';


create table mall_goods
(
    id          bigint auto_increment,
    name        varchar(255)  default ''      not null,
    images      varchar(1023) default '{}'    not null,
    description text                          null,
    create_at   datetime      default (now()) not null,
    update_at   datetime      default (now()) not null on update CURRENT_TIMESTAMP,
    primary key (id)
) comment 'user table';