CREATE TABLE mail_attachment (
    internal_id BIGSERIAL PRIMARY KEY,
    public_idUUID not NULL default gen_random_uuid (),
    mail_internal_id BIGINT NOT NULL,
    user_internal_id BIGINT NOT NULL,
    file TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_mail_attachment_mail FOREIGN KEY (mail_internal_id) REFERENCES mails (internal_id) ON DELETE CASCADE
);

CREATE INDEX idx_mail_attachments_mail ON mail_attachment (mail_internal_id);

CREATE INDEX idx_mail_attachments_user ON mail_attachment (user_internal_id);