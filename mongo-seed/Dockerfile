FROM mongo

COPY init-minerals.json /init-minerals.json
COPY init-mineral-types.json /init-mineral-types.json
CMD mongoimport --host mongodb --db carbon-based --collection minerals --mode=upsert --type json --file /init-minerals.json --jsonArray && mongoimport --host mongodb --db carbon-based --collection mineralTypes --mode=upsert --type json --file /init-mineral-types.json --jsonArray
