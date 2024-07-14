CREATE TABLE identification_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO identification_types (name) VALUES 
('Cédula de ciudadanía'), 
('Cédula de extranjería'), 
('Pasaporte');

DROP TABLE IF EXISTS motorcycles;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    identification_type_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    password VARCHAR(255) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (identification_type_id) REFERENCES identification_types(id)
);

-- Crear la tabla con los cambios solicitados
CREATE TABLE motorcycles (
                             id SERIAL PRIMARY KEY,
                             user_id VARCHAR(50) NOT NULL,
                             plate VARCHAR(50) NOT NULL UNIQUE,
                             brand VARCHAR(255) NOT NULL,
                             model VARCHAR(255) NOT NULL,
                             year INT NOT NULL,
                             soat_file VARCHAR(255),
                             photo_file VARCHAR(255),
                             identification_file VARCHAR(255),
                             mechanical_technician_file VARCHAR(255),
                             created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                             FOREIGN KEY (user_id) REFERENCES users(id)
);

select * from users;
select * from motorcycles;