# layer-to-csv

This example shows how to dump an entire layer to a CSV file

You would run it with:

```
ARCGIS_USERNAME={username} ARCGIS_PASSWORD={password} ./layet-to-csv -feature-server-url https://services7.arcgis.com/q3SI94vj8qWDxwBr/ArcGIS/rest/services/Public_Parcels/FeatureServer -layer-index 0
```

It would then write an `output.csv` file that you can import with:

```
nidus-sync=> \copy import.parcel_csv from '/home/eliribble/src/nidus-sync/arcgis-go/examples/parcel-mapping/output.csv' delimiters ',' csv header;
```

assuming that path was valid for you and you already had the `import.parcel_csv` table created.

