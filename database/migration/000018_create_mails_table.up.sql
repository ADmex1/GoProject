CREATE TABLE mails (
    internal_id BIGSERIAL PRIMARY KEY,
    public_id UUID not NULL default gen_random_uuid (),
    mail_thread_internal_id BIGINT NOT NULL,
    user_internal_id BIGINT NOT NULL,
    user_public_id UUID NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_mail_thread FOREIGN KEY (mail_thread_internal_id) REFERENCES mail_threads (internal_id) ON DELETE CASCADE
);

CREATE INDEX idx_mails_thread ON mails (mail_thread_internal_id);

CREATE INDEX idx_mails_user ON mails (user_internal_id);