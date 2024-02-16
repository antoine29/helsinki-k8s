# DB commitment   

GCP cloud sql was selected for this exercise.   
why?   
Sometime ago I've heard a phrase telling that you should not spent resources on maitaining any other thing but the main thing wich your bussiness runs. In this case, our project is a simple ToDo app. The DB is a dependency but it's not the main component, hence it's fine to use a provider so we can focus on the product itself.

After creating the DB instance, the DB private ip has to set on BE deployment.   


## ToDo:   
We could create the SQL cloud instance with terraform. The private IP would be an output, could we set this IP with kustomize on CI pipeline?   
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


