CREATE TABLE IF NOT EXISTS company (
ID  integer unique,
FOREiGN KEY(ID) REFERENCES general(ID),
WorkField text not null
)