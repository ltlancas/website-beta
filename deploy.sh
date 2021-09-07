docker build -t us.gcr.io/<PROJECT-ID>/<IMAGE>:latest .
docker push us.gcr.io/<PROJECT-ID>/<IMAGE>:latest
gcloud run services update <CLOUD RUN SERVICE NAME> --image=us.gcr.io/<PROJECT-ID>/<IMAGE>:latest
