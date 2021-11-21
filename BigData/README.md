# The link of video:

  https://app.box.com/embed_widget/s/w4cjhk46io526kqxqsm995sshf2fxhbj?view=list&sort=name&direction=ASC&theme=dark
  
# Prepare all the docker files needed:
 Driver : https://hub.docker.com/repository/docker/haoyup/driver  
   
 JupyterNotebook : https://hub.docker.com/repository/docker/haoyup/jupyter-notebook  
   
 Snarqube : https://hub.docker.com/repository/docker/haoyup/sonarqube  
   
 Spark : https://hub.docker.com/r/charlesjiang1997/spark  
   
# Prepare all the deployment files:
 Refer to the Deployment directory.
# Deploy on the GCP:
## 1. Create a kubernetes cluster on GCP and connect to the cluster.
![image](https://user-images.githubusercontent.com/54975123/142711083-04443020-5f2d-46f8-aee4-5a4b5a09b677.png)
## 2. Clone the repository and direct to the Deployment file.
![image](https://user-images.githubusercontent.com/54975123/142711269-6c4cd9f8-8058-46d9-8a0b-f2e1931aec4c.png)
![image](https://user-images.githubusercontent.com/54975123/142711275-845eba26-0507-41a6-85ca-ce09db619f1a.png)

## 3. Apply deployment.yaml and service.yaml
### Driver:
```kubectl apply -f driver-deployment.yaml```   
  
```kubectl apply -f driver-service.yaml```  
  
### Spark:
```kubectl apply -f spark-deployment.yaml```  
  
```kubectl apply -f spark-service.yaml```  
  
### JupyterNotebook:
```kubectl apply -f jupyter-notebook-deployment.yaml```  
  
```kubectl apply -f jupyter-notebook-service.yaml```  
  
### Sonarqube:
```kubectl apply -f sonarqube-deployment.yaml```  
  
```kubectl apply -f sonarqube-service.yaml```  
  
### Hadoop:
```kubectl create -f testhadoop.yaml```  
  
![image](https://user-images.githubusercontent.com/54975123/142711283-26e3c781-3750-4285-8a78-d6b2bbfb4c80.png)

## 4. Result of deployment
We can run ```kubectl get pods``` to see the running pods.  
  ![image](https://user-images.githubusercontent.com/54975123/142711837-49666268-49e7-4f1e-9f8e-bb041be2a9da.png)  
  
Or we can use the GCP dashboard to see the pods and services.  
  
![image](https://user-images.githubusercontent.com/54975123/142711859-e4d47001-b008-4338-b6a8-641ced20fd39.png)
![image](https://user-images.githubusercontent.com/54975123/142711862-84ac4f5e-c6d1-41ef-a769-6df66081bc11.png)

# Get the ip address of the driver and change the address of server.go
At first, the ```currentIP``` of server.go is listend on localhost, once it is exposed, we need to change it to the external ip address.  
  
![image](https://user-images.githubusercontent.com/54975123/142713479-defa6017-f9ef-4a43-8108-116fa9359b11.png)
# Rebuild and push driver image
```docker build -t haoyup/driver:v1 .```  
  
```docker push haoyup/driver:v1```  
  
# Redeploy driver
  ```kubectl delete deployment my-driver-deployment```  
  ```kubectl apply -f driver-deployment.yaml```  
# Launch the BigTableTerminal.jar and input the ip address
The GUI send request to the exposed ip address of server, to indicate the ip address, we need to input it into the text feild and click ok.  
  
![image](https://user-images.githubusercontent.com/54975123/142711326-277facce-7041-434a-88d7-0249fca020ee.png)

# Choose your service
![image](https://user-images.githubusercontent.com/54975123/142711331-ceba2b7f-e050-4b6a-857a-11770105d169.png)
