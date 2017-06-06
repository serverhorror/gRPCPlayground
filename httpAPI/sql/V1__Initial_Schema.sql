CREATE TABLE samples (
  id BIGSERIAL PRIMARY KEY,
  prefix TEXT NOT NULL,
  provider TEXT NOT NULL,
  site TEXT NOT NULL,
  -- cloud_uri TEXT NOT NULL,

  creator TEXT,
  created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

  UNIQUE (id, prefix)
);

CREATE TABLE sample_files (
  id BIGSERIAL PRIMARY KEY,
  sample_id BIGINT NOT NULL REFERENCES samples(id) ON DELETE CASCADE,
  source_uri TEXT NOT NULL UNIQUE, -- file://, http://, s3://, gcs://, ...
  cloud_uri TEXT NOT NULL UNIQUE, -- http://, s3://, gcs://, ...

  creator TEXT,
  created TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE VIEW v_all_samples AS
 SELECT s.id AS sample_id,
    sf.source_uri,
    sf.cloud_uri,
    s.creator,
    s.created
   FROM samples s
     JOIN sample_files sf ON s.id = sf.sample_id;