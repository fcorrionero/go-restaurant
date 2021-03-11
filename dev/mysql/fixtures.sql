CREATE DATABASE IF NOT EXISTS go_restaurant;

USE go_restaurant;

CREATE TABLE IF NOT EXISTS dishes (
    id BINARY(16) PRIMARY KEY,
    id_uuid VARCHAR(40) NOT NULL,
    disn_name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    ingredients JSON
);

CREATE TABLE IF NOT EXISTS dishes_ingredients (
    dish_id binary,
    ingredient_id binary,
    FOREIGN KEY (dish_id) REFERENCES dishes(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
);


CREATE TABLE IF NOT EXISTS ingredients (
    id BINARY(16) PRIMARY KEY,
    id_uuid VARCHAR(40) NOT NULL,
    ingredient_name VARCHAR(255) NOT NULL,
    allergens JSON
);

CREATE TABLE IF NOT EXISTS ingredients_allergens (
    allergen_id binary,
    ingredient_id binary,
    FOREIGN KEY (allergen_id) REFERENCES allergens(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
);

CREATE TABLE IF NOT EXISTS allergens (
    id BINARY(16) PRIMARY KEY ,
    id_uuid VARCHAR(40) NOT NULL,
    allergen_name VARCHAR(255) NOT NULL
);