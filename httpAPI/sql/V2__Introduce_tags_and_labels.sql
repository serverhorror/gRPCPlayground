ALTER TABLE samples ALTER COLUMN provider DROP NOT NULL;

CREATE TABLE tags (
  id BIGSERIAL PRIMARY KEY,
  tag TEXT NOT NULL UNIQUE,
  created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE sample_tags (
  sample_id BIGINT NOT NULL REFERENCES samples(id) ON DELETE CASCADE,
  tag_id BIGINT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,

  PRIMARY KEY(sample_id, tag_id)
);

CREATE TABLE labels (
  id BIGSERIAL PRIMARY KEY,
  key TEXT NOT NULL,
  value TEXT NOT NULL,
  created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE sample_labels (
  sample_id BIGINT NOT NULL REFERENCES samples(id) ON DELETE CASCADE,
  label_id BIGINT NOT NULL REFERENCES labels(id) ON DELETE CASCADE,

  PRIMARY KEY(sample_id, label_id)
);

-- insert into labels (key, value) select distinct initcap('site') AS key, site as value from samples;