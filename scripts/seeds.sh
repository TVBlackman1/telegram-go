# start file from current directory: ./

if [ -f ../develop.env ]; then
  export $(cat ../develop.env | xargs)
else
  echo "develop.env no found"
  exit
fi

docker run --rm \
    --add-host=host.docker.internal:host-gateway \
    -v ${PWD}/../init/seeds:/flyway/sql \
    flyway/flyway:9 \
    -url=jdbc:postgresql://host.docker.internal:5432/${POSTGRES_DBNAME} \
    -user=${POSTGRES_USER} \
    -password=${POSTGRES_PASS} \
    -outOfOrder="true" \
    migrate