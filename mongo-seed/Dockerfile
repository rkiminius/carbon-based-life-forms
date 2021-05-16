FROM mongo

COPY init-mineral-types.json /init-mineral-types.json
CMD mongoimport --host mongodb --db carbon-based --collection mineralTypes --mode=upsert --type json --file /init-mineral-types.json --jsonArray