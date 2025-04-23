#!/usr/bin/env bash
# Yields nothing, an empty HTML pagi needing redirect
#export ENDPOINT=https://deltavcd.maps.arcgis.com/arcgis/rest/services

# yields test-data/rest.json
#export ENDPOINT=https://services8.arcgis.com/pV7SH1EgRc6tpxlJ/arcgis/rest/info/
# curl $ENDPOINT/\
# 	-D - \
# 	-d "f=pjson"

# yields test-data/services.json
#export ENDPOINT=https://services8.arcgis.com/pV7SH1EgRc6tpxlJ/arcgis/rest/services/
# curl $ENDPOINT/\
# 	-D - \
# 	-d "f=pjson"

# This gives us a 404, according to the "services" endpoint the only "type" for "BorderDistrict" is "FeatureServer"
#export ENDPOINT=https://services8.arcgis.com/pV7SH1EgRc6tpxlJ/arcgis/rest/services/BorderDistrict

# yields test-data/feature_server.json
#export ENDPOINT=https://services8.arcgis.com/pV7SH1EgRc6tpxlJ/arcgis/rest/services/BorderDistrict/FeatureServer
# curl $ENDPOINT/\
# 	-D - \
# 	-d "f=pjson"

# yields test-data/layer.json
#export ENDPOINT=https://services8.arcgis.com/pV7SH1EgRc6tpxlJ/arcgis/rest/services/BorderDistrict/FeatureServer/6/
# curl $ENDPOINT/\
# 	-D - \
# 	-d "f=pjson"


# yields test-data/query.json
# But you must add '-d "where=1%3D1"' to curl command below.
#export ENDPOINT='https://services8.arcgis.com/pV7SH1EgRc6tpxlJ/arcgis/rest/services/BorderDistrict/FeatureServer/6/query'
# curl $ENDPOINT/\
# 	-D - \
# 	-d "f=pjson" \
#	-d "where=1%3D1"
