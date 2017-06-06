with cte AS (
SELECT
  s.*,
  rank() over (ORDER BY count(DISTINCT l.id) DESC, count(DISTINCT t.id) DESC) as rank
FROM
  samples s 
  JOIN sample_labels sl ON (s.id = sl.sample_id)
  JOIN labels l ON (
    sl.label_id = l.id and 
    (
      -- this will be dynamic depending on number of label (k/v pairs)
      -- specified by user
      -- might be [0,n] conditions
      (l.key = 'narf' AND l.value = 'zord')
      OR (l.key = 'narf' AND l.value = 'bar')
      OR (l.key='foo' AND l.value = 'bar')
    )
  )
  JOIN sample_tags st ON (s.id = st.sample_id)
  JOIN tags t ON (
    st.tag_id = t.id AND
    (
      -- this will be dynamic depending on number of tags specified by user
      -- might be [0,n] conditions
      t.tag = 't1'
      OR t.tag = 't2'
    )
  )
GROUP BY
  s.id, s.prefix
)
SELECT
  cte.*,
  array_agg(DISTINCT t.tag) as tags,
  array_agg(DISTINCT l.key || '=' || l.value) as labels
FROM
  cte
  JOIN sample_labels sl ON (cte.id = sl.sample_id)
  JOIN labels l ON (sl.label_id = l.id)
  JOIN sample_tags st ON (cte.id = st.sample_id)
  JOIN tags t ON (st.tag_id = t.id)
WHERE cte.rank = 1
GROUP BY
  cte.id, cte.prefix, cte.rank