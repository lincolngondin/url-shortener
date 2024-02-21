CREATE TABLE urls (
    original_url text,
    shortened_url text PRIMARY KEY,
    creation_time timestamp,
    last_click timestamp,
    total_clicks integer
);
