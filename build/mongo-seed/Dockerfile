FROM mongo:latest as build

RUN mkdir "/collections"

ENV MONGO_HOST=localhost
ENV MONGO_USER=root
ENV MONGO_PASS=toor
ENV MONGO_FILES=/collections

ADD build/mongo-seed/import.sh import.sh
RUN chmod +x import.sh

## put schemas to be imported in
# /collections/<schema_name>/<collection_name>.json
#copy schemas collections

CMD ["./import.sh"]
FROM build as final