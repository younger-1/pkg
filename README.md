
<https://neon.tech/postgresql/postgresql-getting-started/connect-to-postgresql-database>

<http://www.postgres.cn/docs/current/app-pg-ctl.html>

```bash
# 0. Init database
initdb --locale=C -E UTF-8 /usr/local/var/postgresql@14

# 1. Start local server
postgres -D /usr/local/var/postgresql@14
# or
pg_ctl -D '/usr/local/var/postgresql@14' -l pg.log start
pg_ctl -D '/usr/local/var/postgresql@14' -l pg.log stop
# or start daemon (background service)
brew services start postgresql@14

# 2. Start local client
psql -U young -d postgres
# or
psql --username=young --dbname=postgres
```


```sql
-- create the database
psql> create database recordings;
psql> \l+ -- see all databases

-- choose the database
psql> \c recordings
mysql> use recordings;

-- run the script
psql> \i create-tables.sql
mysql> source create-tables.sql

-- show table info
psql> \d+ album
 Column |          Type          | Collation | Nullable |              Default
--------+------------------------+-----------+----------+-----------------------------------
 id     | integer                |           | not null | nextval('album_id_seq'::regclass)
 title  | character varying(128) |           | not null |
 artist | character varying(255) |           | not null |
 price  | numeric(5,2)           |           | not null |
Indexes: "album_pkey" PRIMARY KEY, btree (id)

-- show table detail
psql> select * from album;
 id |     title     |     artist     | price
----+---------------+----------------+-------
  1 | Blue Train    | John Coltrane  | 56.99
  2 | Giant Steps   | John Coltrane  | 63.99
  3 | Jeru          | Gerry Mulligan | 17.99
  4 | Sarah Vaughan | Sarah Vaughan  | 34.98
(4 rows)
```
