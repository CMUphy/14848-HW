# Prepare all the docker files needed:
 Driver : https://hub.docker.com/repository/docker/haoyup/driver  
 JupyterNotebook : https://hub.docker.com/repository/docker/haoyup/jupyter-notebook  
 Snarqube : https://hub.docker.com/repository/docker/haoyup/sonarqube  
 Spark : https://hub.docker.com/r/charlesjiang1997/spark  
# Prepare all the deployment files:
 Refer to the Deployment directory.
# Deploy on the GCP:
## Create a kubernetes cluster on GCP and connect to the cluster.
![image](https://user-images.githubusercontent.com/54975123/142711083-04443020-5f2d-46f8-aee4-5a4b5a09b677.png)
## Clone the repository and direct to the Deployment file.
![image](https://user-images.githubusercontent.com/54975123/142711269-6c4cd9f8-8058-46d9-8a0b-f2e1931aec4c.png)
![image](https://user-images.githubusercontent.com/54975123/142711275-845eba26-0507-41a6-85ca-ce09db619f1a.png)

## Apply deployment.yaml and service.yaml
### Deiver:
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

## Result of deployment
We can run ```kubectl get pods``` to see the running pods.  
![image](https://user-images.githubusercontent.com/54975123/142711498-82825c3a-c26b-4d64-9a21-e6555d589ee0.png)  
Or we can use the GCP dashboard to see the pods and services.  
![image](https://user-images.githubusercontent.com/54975123/142711294-b221cedb-b82e-41ef-bc0d-86dfb5ad3446.png)
![image](https://user-images.githubusercontent.com/54975123/142711299-fe3128b1-5088-4bf5-bac0-4aa5b1016e18.png)

# Get the ip address of the driver and change the address of server.go
At first, the ip address of server.go is listend on localhost, once it is exposed, we need to change it to the external ip address.  
![image](https://user-images.githubusercontent.com/54975123/142711303-68a19398-0390-4ff1-be35-00ff9d501501.png)

# Rebuild and push driver image
```docker build -t haoyup/driver:v1 .```  
```docker push haoyup/driver:v1```  
# Launch the BigTableTerminal.jar and input the ip address
The GUI send request to the exposed ip address of server, to indicate the ip address, we need to input it into the text feild and click ok.  
![image](https://user-images.githubusercontent.com/54975123/142711326-277facce-7041-434a-88d7-0249fca020ee.png)

# Choose your service
![image](https://user-images.githubusercontent.com/54975123/142711331-ceba2b7f-e050-4b6a-857a-11770105d169.png)
