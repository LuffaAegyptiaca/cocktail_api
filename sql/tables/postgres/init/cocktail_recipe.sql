create table cocktail_recipe (
    id              serial   PRIMARY KEY,
    name            text     NOT NULL,
    recipe          text     NOT NULL,
    cocktail_style  text     NOT NULL,
    alcohol         smallint NOT NULL
);