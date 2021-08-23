create database cocktail;
CREATE TABLE cocktail_ingredient (
    id    serial   PRIMARY KEY,
    name  text     NOT NULL,
    type  text     NOT NULL
);