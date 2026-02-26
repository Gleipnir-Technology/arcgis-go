# parcel-mapping

This example shows how to pull down publicly available parcel information and map some of the attributes.

You would run it with:

```
ARCGIS_USERNAME={username} ARCGIS_PASSWORD={password} ./parcel-mapping -feature-server-url https://services7.arcgis.com/q3SI94vj8qWDxwBr/ArcGIS/rest/services/Public_Parcels/FeatureServer -layer-index 0 -feature-name-apn "APN_ID" -feature-name-description "PropertySitus"
```

It would then print out:

```
9:26PM INF apn=121272016 desc="2530 HEMLOCK DR" geom="poly with 1 rings"
```
