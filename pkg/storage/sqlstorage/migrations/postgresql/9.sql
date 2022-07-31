--statement
ALTER TABLE "VAR_LEDGER_NAME".transactions
ADD COLUMN IF NOT EXISTS "_sources" varchar[],
ADD COLUMN IF NOT EXISTS "_destinations" varchar[];
--statement
CREATE INDEX IF NOT EXISTS txs_sources ON "VAR_LEDGER_NAME".transactions USING GIN(_sources);
CREATE INDEX IF NOT EXISTS txs_destinations ON "VAR_LEDGER_NAME".transactions USING GIN(_destinations);
--statement
CREATE OR REPLACE FUNCTION "VAR_LEDGER_NAME".handle_log_entry()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
AS
$$
BEGIN
    if NEW.type = 'NEW_TRANSACTION' THEN
        INSERT INTO "VAR_LEDGER_NAME".transactions(
            id,
            timestamp,
            reference,
            postings,
            metadata,
            pre_commit_volumes,
            post_commit_volumes,
            _sources,
            _destinations
        )
        VALUES (
                (NEW.data ->> 'txid')::bigint,
                (NEW.data ->> 'timestamp')::varchar,
                CASE
                    WHEN (NEW.data ->> 'reference')::varchar = '' THEN NULL
                    ELSE (NEW.data ->> 'reference')::varchar END,
                (NEW.data ->> 'postings')::jsonb,
                CASE WHEN (NEW.data ->> 'metadata')::jsonb IS NULL THEN '{}' ELSE (NEW.data ->> 'metadata')::jsonb END,
                (NEW.data ->> 'preCommitVolumes')::jsonb,
                (NEW.data ->> 'postCommitVolumes')::jsonb,
                (WITH t as (SELECT * FROM json_to_recordset((NEW.data ->> 'postings')::json) as x(source varchar)) SELECT ARRAY (SELECT source from t)),
                (WITH t as (SELECT * FROM json_to_recordset((NEW.data ->> 'postings')::json) as x(destination varchar)) SELECT ARRAY (SELECT destination from t))
        );
    END IF;
    if NEW.type = 'SET_METADATA' THEN
        if NEW.data ->> 'targetType' = 'TRANSACTION' THEN
            UPDATE "VAR_LEDGER_NAME".transactions
            SET metadata = metadata || (NEW.data ->> 'metadata')::jsonb
            WHERE id = (NEW.data ->> 'targetId')::bigint;
        END IF;
        if NEW.data ->> 'targetType' = 'ACCOUNT' THEN
            INSERT INTO "VAR_LEDGER_NAME".accounts (address, metadata)
            VALUES ((NEW.data ->> 'targetId')::varchar,
                    (NEW.data ->> 'metadata')::jsonb)
            ON CONFLICT (address) DO UPDATE SET metadata = accounts.metadata || (NEW.data ->> 'metadata')::jsonb;
        END IF;
    END IF;
    RETURN NEW;
END;
$$;
--statement
UPDATE "VAR_LEDGER_NAME".transactions SET _sources = (WITH t as (SELECT * FROM json_to_recordset(postings::json) as x(source varchar)) SELECT array_agg(DISTINCT ARRAY (SELECT source from t))) where id < 50000;
--statement
UPDATE "VAR_LEDGER_NAME".transactions SET _destinations = (WITH t as (SELECT * FROM json_to_recordset(postings::json) as x(destination varchar)) SELECT array_agg(DISTINCT ARRAY (SELECT destination from t))) where id < 50000;
--statement
CREATE OR REPLACE FUNCTION use_account(_sources varchar[], _destinations varchar[], account varchar)
    RETURNS BOOLEAN
AS $$
SELECT _sources @> (ARRAY[account]) OR _destinations @> (ARRAY[account])
$$ LANGUAGE sql;