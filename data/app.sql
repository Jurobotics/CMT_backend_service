CREATE TABLE
IF NOT EXISTS recipe
(
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT,
    beschreibung TEXT,
    zubereitung TEXT,
    bild VARCHAR(255)
);

-- CREATE TABLE
-- IF NOT EXISTS module
-- (
--     id INTEGER NOT NULL PRIMARY KEY,
--     name TEXT,
--     dosierer int,
--     stand int
-- );

CREATE TABLE
IF NOT EXISTS servos
(
	id INTEGER NOT NULL PRIMARY KEY,
    arduino INTEGER NOT NULL
);

CREATE TABLE
IF NOT EXISTS ingredient
(
	id INTEGER NOT NULL PRIMARY KEY,
	name varchar(255) not null,
	-- unit varchar(255),
	amount int not null, -- Dosierer in L angeben
    servo_id int,
    uses int,
    FOREIGN KEY (servo_id) REFERENCES servos(id)
);


CREATE TABLE
IF NOT EXISTS ingredients
(
    id INTEGER NOT NULL PRIMARY KEY,
    recipe_id int,
    ingredient_id int,
    amount int,
    FOREIGN KEY (recipe_id) REFERENCES recipe(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredient(id)
)
