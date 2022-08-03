docker run --rm \
    --add-host=host.docker.internal:host-gateway \
    -v ${PWD}/../init/migrations:/flyway/sql \
    flyway/flyway:9.1.2-alpine \
    -url=jdbc:postgresql://host.docker.internal:5432/tgbot \
    -user=postgres \
    -password=postgres \
    migrate