CREATE TABLE users (
    id          uuid        PRIMARY KEY DEFAULT gen_random_uuid(),
    username    VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    role        VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL,
    updated_at  TIMESTAMPTZ  NOT NULL
);
--
-- INSERT INTO users (
--     id,
--     username,
--     password,
--     role,
--     created_at,
--     updated_at
-- ) VALUES (
--     '00000000-0000-0000-0000-000000000000',
--     'admin',
--     'admin',
--     NOW(),
--     NOW()
-- );
