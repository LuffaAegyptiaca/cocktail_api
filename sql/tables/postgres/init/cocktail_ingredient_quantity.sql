create database cocktail;
create table cocktail_recipe (
    cocktail_recipe_id     smallint REFERENCES cocktail_recipe,
    cocktail_ingredient_id smallint REFERENCES cocktail_ingredient,
    name                   text     NOT NULL,
    quantity               smallint NOT NULL,
    unit                   text     NOT NULL,
    base_flag              boolean  NOT NULL,
    PRIMARY KEY (cocktail_recipe_id, cocktail_ingredient_id)
);