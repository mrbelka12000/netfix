CREATE TABLE IF NOT EXISTS general (
    ID  serial primary key,
    UserName text  not null,
    Email text  not null,
    Password text not null
)