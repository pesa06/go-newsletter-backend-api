INSERT INTO users
    (
        username,
        password,
        role,
        created_at,
        updated_at
    ) VALUES (
        @username,
        @password,
        @role,
        NOW(),
        NOW()
)