# tempest-user-service  
![linting](https://github.com/neverett8fr/tempest-user-service/actions/workflows/golangci-lint.yml/badge.svg)  
  

⚠️ Note: This User Service is no longer actively maintained or used in the Tempest cloud storage solution. The following README is provided for historical reference only.  
  
The User Service was a component of the Tempest cloud storage solution responsible for managing user-related operations, including user registration, authentication, and profile management.  

  
  
# How to run  
  
## Build  
```bash
docker build -t .
 ```
   
 ## Run  
 ```bash
docker run -p 8080:8080 -v . -e ENV_VARIABLE=value .
 ```
   
 ## Stop the container  
 ```bash
 docker stop container-name
 ```
