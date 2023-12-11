CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    NEW.doc_version = NEW.doc_version + 1;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
