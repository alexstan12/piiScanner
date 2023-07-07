# piiScanner
Scan databases and data warehouses for PII data. Inspired by  [piicatcher](https://tokern.io/piicatcher/).

The main purpose of this library is to be imported and used as such in your own code, so feel free to do so.

A makefile target can be used to test the PII detector against your own live running db. Specifically, it will check column names but not the actual data in the rows.
Running this target requires the following ENV variables to be specified:
- DB_TYPE -> the sql driver
- DB_NAME -> the database name
- SCHEMA_NAME -> name of the schema
- CREDENTIALS -> the actual user credentials
  
For example:
```
make test DB_TYPE=snowflake DB_NAME=snowflake SCHEMA_NAME=schema_example CREDENTIALS='test:test123!@test/foobar'
```

``DISCLAIMER : this was tested only with a Snowflake DB instance and is not guaranteed to work with other implementations !``
