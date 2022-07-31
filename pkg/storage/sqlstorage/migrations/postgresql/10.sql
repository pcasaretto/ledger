--statement
CREATE OR REPLACE FUNCTION use_account(_sources varchar[], _destinations varchar[], account varchar)
    RETURNS BOOLEAN
AS $$
SELECT _sources @> (ARRAY[account]) OR _destinations @> (ARRAY[account])
$$ LANGUAGE sql;