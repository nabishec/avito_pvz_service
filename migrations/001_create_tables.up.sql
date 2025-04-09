-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL,
    registration_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS pvzs (
    id UUID PRIMARY KEY,
    city TEXT NOT NULL,
    registration_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS receptions (
    id UUID PRIMARY KEY,
    pvz_id UUID NOT NULL REFERENCES PVZs(id) ON DELETE CASCADE,
    status TEXT NOT NULL,
    registration_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    reception_id UUID NOT NULL REFERENCES Receptions(id) ON DELETE CASCADE,
    type TEXT NOT NULL,
    registration_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd