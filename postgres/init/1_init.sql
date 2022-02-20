--DBの作成
CREATE DATABASE cocktail;
-- 利用DB切替
\c cocktail;

-- 各種テーブルを作成
CREATE TABLE cocktail_recipe (
    id              serial   PRIMARY KEY,
    name            text     NOT NULL,
    recipe          text     NOT NULL,
    cocktail_style  text     NOT NULL,
    alcohol         smallint NOT NULL
);
CREATE TABLE cocktail_ingredient (
    id    serial   PRIMARY KEY,
    name  text     NOT NULL,
    type  text     NOT NULL
);
CREATE TABLE cocktail_ingredient_quantity (
    cocktail_recipe_id     smallint REFERENCES cocktail_recipe,
    cocktail_ingredient_id smallint REFERENCES cocktail_ingredient,
    quantity               smallint NOT NULL,
    unit                   text     NOT NULL,
    base_flag              boolean  NOT NULL,
    PRIMARY KEY (cocktail_recipe_id, cocktail_ingredient_id)
);
-- サンプルデータ挿入
INSERT INTO cocktail_recipe (
    name,
    recipe,
    cocktail_style,
    alcohol
)
VALUES 
    ('ジントニック', 'いい感じに混ぜるとできる', 'ロング', 5),
    ('stmnオリジナルカクテル', '気分に合わせて度数高い酒を入れまくって混ぜる', 'ショート', 40)
;
INSERT INTO cocktail_ingredient (
    name,
    type
)
VALUES 
    ('ラム', 'アルコール'),
    ('ジン', 'アルコール'),
    ('トニックウォーター', 'ノンアルコール'),
    ('ウイスキー', 'アルコール'),
    ('テキーラ', 'アルコール')
;
INSERT INTO cocktail_ingredient_quantity
VALUES
    (1,2,45,'ml',true),
    (1,3,-1,'適量',true),
    (2,1,30,'ml',true),
    (2,2,30,'ml',true),
    (2,4,30,'ml',true),
    (2,5,30,'ml',true)
;