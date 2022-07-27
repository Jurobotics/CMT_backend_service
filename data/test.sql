CREATE TABLE
IF NOT EXISTS rezepte
(id INTEGER PRIMARY KEY, name TEXT, beschreibung TEXT, zutaten TEXT, zubereitung TEXT);

CREATE TABLE
IF NOT EXISTS module
(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, dosierer int);

