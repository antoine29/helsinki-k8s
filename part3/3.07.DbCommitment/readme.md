# gcloud setup
gcloud auth login
gcloud config set project gke-test-412223

# enables required services (tf available?)
gcloud services enable compute.googleapis.com
gcloud services enable container.googleapis.com

# kubectl setup
gcloud components install gke-gcloud-auth-plugin
gcloud container clusters get-credentials test --zone us-central1-c --project innate-mapper-411620

# gh action service account setup (tf available?)
gcloud iam service-accounts create github-actions

gcloud iam service-accounts keys create ./private-key.json --iam-account=github-actions@innate-mapper-411620.iam.gserviceaccount.com

gcloud projects add-iam-policy-binding innate-mapper-411620 \
    --member="serviceAccount:github-actions@innate-mapper-411620.iam.gserviceaccount.com" \
    --role="roles/container.serviceAgent"

gcloud projects add-iam-policy-binding innate-mapper-411620 \
    --member="serviceAccount:github-actions@innate-mapper-411620.iam.gserviceaccount.com" \
    --role="roles/storage.admin"


