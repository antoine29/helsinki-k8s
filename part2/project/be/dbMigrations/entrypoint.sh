#!/bin/sh
/liquibase/liquibase --log-level debug --classpath="/" --url "jdbc:postgresql://${PG_HOST}:${PG_PORT}/${PG_DBNAME}?ssl=false" --username="${PG_USERNAME}" --password="${PG_PASSWORD}" --changelog-file="${CHANGE_LOG}" --hub-mode=off update
