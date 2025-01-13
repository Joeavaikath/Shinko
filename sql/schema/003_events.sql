-- +goose Up
CREATE TABLE action_events (
    id UUID PRIMARY KEY,
    action_id UUID NOT NULL REFERENCES actions(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    executed_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    comment TEXT DEFAULT 'No comments added'
);

-- +goose Down
DROP TABLE action_events;