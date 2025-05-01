-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS workouts (
    id BIGSERIAL PRIMARY KEY,
    --user id
    description TEXT,
    duration_minutes INTEGER NOT NULL,
    calories_burder INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd 
-- +goose Down
-- +goose StatementBegin
DROP TABLE workouts;
-- +goose StatementEnd
