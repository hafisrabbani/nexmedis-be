-- =============================================
-- EXTENSIONS
-- =============================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- =============================================
-- ENUMS
-- =============================================

DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'client_status') THEN
            CREATE TYPE client_status AS ENUM ('active', 'revoked');
        END IF;
    END
$$;

-- =============================================
-- TABLE: clients
-- =============================================

CREATE TABLE IF NOT EXISTS clients
(
    id           UUID PRIMARY KEY       DEFAULT uuid_generate_v4(),
    client_id    VARCHAR(64)   NOT NULL UNIQUE,
    name         VARCHAR(255)  NOT NULL,
    email        BYTEA         NOT NULL, -- encrypted email
    api_key_hash VARCHAR(255)  NOT NULL,
    status       client_status NOT NULL DEFAULT 'active',
    created_at   TIMESTAMP     NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_clients_api_key_hash
    ON clients (api_key_hash);

-- =============================================
-- TABLE: api_logs (PARTITIONED)
-- =============================================

CREATE TABLE IF NOT EXISTS api_logs
(
    id         BIGSERIAL,
    client_id  UUID         NOT NULL,
    endpoint   VARCHAR(255) NOT NULL,
    ip_address INET         NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id, created_at),
    CONSTRAINT fk_api_logs_client
        FOREIGN KEY (client_id) REFERENCES clients (id)
) PARTITION BY RANGE (created_at);

-- =============================================
-- DEFAULT PARTITION (SAFETY NET)
-- =============================================

CREATE TABLE IF NOT EXISTS api_logs_default
    PARTITION OF api_logs
        DEFAULT;

-- =============================================
-- SAMPLE MONTHLY PARTITIONS
-- (extend via cron / migration)
-- =============================================

CREATE TABLE IF NOT EXISTS api_logs_2025_01
    PARTITION OF api_logs
        FOR VALUES FROM ('2025-01-01') TO ('2025-02-01');

CREATE TABLE IF NOT EXISTS api_logs_2025_02
    PARTITION OF api_logs
        FOR VALUES FROM ('2025-02-01') TO ('2025-03-01');

-- =============================================
-- INDEXES FOR api_logs
-- =============================================

CREATE INDEX IF NOT EXISTS idx_api_logs_client_time
    ON api_logs (client_id, created_at);

CREATE INDEX IF NOT EXISTS idx_api_logs_time
    ON api_logs (created_at);

-- =============================================
-- TABLE: daily_usage
-- =============================================

CREATE TABLE IF NOT EXISTS daily_usage
(
    id             BIGSERIAL PRIMARY KEY,
    client_id      UUID      NOT NULL,
    date           DATE      NOT NULL,
    total_requests BIGINT    NOT NULL DEFAULT 0,
    created_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_daily_usage_client_date UNIQUE (client_id, date),
    CONSTRAINT fk_daily_usage_client
        FOREIGN KEY (client_id) REFERENCES clients (id)
);

CREATE INDEX IF NOT EXISTS idx_daily_usage_date
    ON daily_usage (date);

-- =============================================
-- TRIGGER: auto update updated_at
-- =============================================

CREATE OR REPLACE FUNCTION set_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_daily_usage_updated_at ON daily_usage;

CREATE TRIGGER trg_daily_usage_updated_at
    BEFORE UPDATE
    ON daily_usage
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();


CREATE TABLE IF NOT EXISTS client_ip_whitelists
(
    id         UUID PRIMARY KEY   DEFAULT uuid_generate_v4(),
    client_id  UUID      NOT NULL,
    ip_address INET      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT uq_client_ip UNIQUE (client_id, ip_address),
    CONSTRAINT fk_ip_whitelist_client
        FOREIGN KEY (client_id) REFERENCES clients (id)
);


-- =============================================
-- DONE
-- =============================================
