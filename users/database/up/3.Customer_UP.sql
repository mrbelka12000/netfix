CREATE TABLE IF NOT EXISTS customer (
    ID  integer unique,
    FOREIGN KEY(ID) REFERENCES general(ID),
    Birth DATE not null
)