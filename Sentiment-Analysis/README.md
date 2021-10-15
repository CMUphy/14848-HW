1. Build Container Image for Each Service and push them to Docker Hub.
FrontImage:
docker build -f Dockerfile -t haoyup/sentiment-analysis-frontend .
docker run -d -p 80:80 haoyup/sentiment-analysis-frontend
docker push haoyup/sentiment-analysis-frontend
LogicIamge:
docker build -f Dockerfile -t haoyup/sentiment-analysis-logic .
docker run -d -p 5050:5000 haoyup/sentiment-analysis-logic
docker push haoyup/sentiment-analysis-logic
WebIamge:
docker build -f Dockerfile -t haoyup/sentiment-analysis-web-app .
docker run -d -p 8080:8080 -e SA_LOGIC_API_URL='http://172.17.0.3:5000' haoyup/sentiment-analysis-web-app  
docker push haoyup/sentiment-analysis-web-app

2. Pull all images to GCP container registery:

docker pull haoyup/sentiment-analysis-frontend
docker tag haoyup/sentiment-analysis-frontend gcr.io/eco-limiter-326323/haoyup/sentiment-analysis-frontend
docker push gcr.io/eco-limiter-326323/haoyup/sentiment-analysis-frontend

docker pull haoyup/sentiment-analysis-web-app
docker tag haoyup/sentiment-analysis-web-app gcr.io/eco-limiter-326323/haoyup/sentiment-analysis-web-app
docker push gcr.io/eco-limiter-326323/haoyup/sentiment-analysis-web-app


docker pull haoyup/sentiment-analysis-logic
docker tag haoyup/sentiment-analysis-logic gcr.io/eco-limiter-326323/haoyup/sentiment-analysis-logic
docker push gcr.io/eco-limiter-326323/haoyup/sentiment-analysis-logic

3. Create Cluster and deploy images repectively, accroding to the dialog, expose web-app and frontend using LoadBalancer