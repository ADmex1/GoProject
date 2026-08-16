CREATE TABLE mail_receivers (
    mail_receiver_internal_id BIGSERIAL PRIMARY KEY,
    mail_thread_internal_id BIGINT NOT NULL,
    user_internal_id BIGINT NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_mail_receiver_thread FOREIGN KEY (mail_thread_internal_id) REFERENCES mail_threads (internal_id) ON DELETE CASCADE
);

CREATE INDEX idx_mail_receivers_thread ON mail_receivers (mail_thread_internal_id);

CREATE INDEX idx_mail_receivers_user ON mail_receivers (user_internal_id);