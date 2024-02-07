CREATE TABLE
    finance (
        id VARCHAR DEFAULT
        REPLACE (
                uuid_generate_v4():: text,
                '-',
                ''
            ) NOT NULL PRIMARY KEY,
            date BIGINT NOT NULL,
            is_debit BOOLEAN NOT NULL,
            user_id VARCHAR NOT NULL,
            description TEXT NOT NULL,
            image_url VARCHAR(100),
            created_at BIGINT NOT NULL,
            updated_at BIGINT NOT NULL,
            deleted_at TIMESTAMP NULL,
            CONSTRAINT fk_user_finance Foreign Key (user_id) REFERENCES users(id)
    );