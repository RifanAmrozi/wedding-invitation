-- Create wedding_users table
CREATE TABLE wedding_users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Create photos table
CREATE TABLE photos (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    photo_url TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES wedding_users(id) ON DELETE CASCADE
);

-- Create comments table
CREATE TABLE comments (
    id UUID PRIMARY KEY,
    photo_id UUID NOT NULL,
    guest_name VARCHAR(255) NOT NULL,
    comment TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (photo_id) REFERENCES photos(id) ON DELETE CASCADE
);
