CREATE DATABASE companies;
\c companies;
BEGIN;
    CREATE TABLE IF NOT EXISTS company (
        id VARCHAR(100) PRIMARY KEY,
        name VARCHAR(15) UNIQUE,
        description VARCHAR(3000),
        employees INTEGER,
        registered BOOLEAN,
        type VARCHAR(25),
        status VARCHAR(25)
    );

    CREATE TABLE IF NOT EXISTS users (
        id VARCHAR(100) PRIMARY KEY,
        name VARCHAR(200) UNIQUE,
        password VARCHAR(5000),
        email VARCHAR(200) UNIQUE,
        phone VARCHAR(20),
        role VARCHAR(20),
        status VARCHAR(100)
    );

    INSERT INTO users (id, name, password, email, phone, role, status) VALUES ('0032edb2-d2de-43f6-8e1f-5c0d2ecbad6b', 'Test', '$2a$14$urZTiJxuWgHjOjITez66leFKnE1OWDgx3.tS/8ng4hZtllYT9choy', 'test@test.com', '+91-1234567', 'admin', 'active');

COMMIT;
