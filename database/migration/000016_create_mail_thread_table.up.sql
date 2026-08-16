CREATE TABLE mail_threads (
    internal_id BIGSERIAL PRIMARY KEY,
    public_id UUID NOT NULL UNIQUE,
    sender_internal_id BIGINT NOT NULL,
    sender_public_id UUID NOT NULL,
    subject VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);