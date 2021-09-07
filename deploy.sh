docker build -t us.gcr.io/ferrous-tine-325314/personal-website:latest .
docker push us.gcr.io/ferrous-tine-325314/personal-website:latest
gcloud run services update personal-website --image=us.gcr.io/ferrous-tine-325314/personal-website:latest
