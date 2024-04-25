CREATE TABLE owners (
     id SERIAL PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
     surname VARCHAR(255) NOT NULL,
     patronymic VARCHAR(255)
);

CREATE TABLE cars (
     id SERIAL PRIMARY KEY,
     reg_num VARCHAR(255) UNIQUE NOT NULL,
     mark VARCHAR(255) NOT NULL,
     model VARCHAR(255) NOT NULL,
     year INT NOT NULL,
     owner_id INTEGER REFERENCES owners(id)
);
