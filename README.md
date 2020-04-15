# gae_basics_webapp_golang

## How to deply the final version.

```
export PROJECT_ID=[Your Project ID]
gcloud config set project $PROJECT_ID
gsutil mb gs://$PROJECT_ID
gsutil iam ch allUsers:objectViewer gs://$PROJECT_ID
gcloud app create

git clone https://github.com/enakai00/gae_basics_webapp_golang.git
cd gae_basics_webapp_golang/guestbook/10_gcs/guestbook_gcs_02/
gcloud app deploy
```
