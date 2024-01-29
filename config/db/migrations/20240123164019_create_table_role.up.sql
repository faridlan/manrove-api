CREATE TABLE
    role (
        uid VARCHAR DEFAULT
        REPLACE (
                uuid_generate_v4():: text,
                '-',
                ''
            ) NOT NULL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            created_at BIGINT NOT NULL,
            updated_at BIGINT NOT NULL,
            deleted_at TIMESTAMP NULL
    );