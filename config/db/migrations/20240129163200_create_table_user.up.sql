CREATE TABLE
    users (
        id VARCHAR DEFAULT
        REPLACE (
                uuid_generate_v4():: text,
                '-',
                ''
            ) NOT NULL PRIMARY KEY,
            email VARCHAR(100) NOT NULL,
            name VARCHAR(100) NOT NULL,
            password VARCHAR(100) NOT NULL,
            phone_number VARCHAR(20) NOT NULL,
            role_id VARCHAR NOT NULL,
            image_url VARCHAR(100),
            first_visit BOOLEAN DEFAULT TRUE,
            created_at BIGINT NOT NULL,
            updated_at BIGINT NOT NULL,
            deleted_at TIMESTAMP NULL,
            CONSTRAINT fk_user_role Foreign Key (role_id) REFERENCES role(uid)
    );