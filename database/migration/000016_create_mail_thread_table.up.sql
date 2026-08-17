CREATE TABLE mail_threads (
    internal_id BIGSERIAL PRIMARY KEY,
    public_id UUID NOT NULL DEFAULT gen_random_uuid (),
    sender_internal_id BIGINT NOT NULL,
    sender_public_id UUID NOT NULL,
    subject VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT mail_threads_public_id_unique UNIQUE (public_id),
    CONSTRAINT fk_mail_threads_sender_internal_id FOREIGN KEY (sender_internal_id) REFERENCES users (internal_id),
    CONSTRAINT fk_mail_threads_sender_public_id FOREIGN KEY (sender_public_id) REFERENCES users (public_id) ON DELETE CASCADE
);