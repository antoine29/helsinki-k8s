# DB commitment   

GCP cloud sql was selected for this exercise.   
why?   
Sometime ago I've heard a phrase telling that you should not spent resources on maitaining any other thing but the main thing wich your bussiness runs. In this case, our project is a simple ToDo app. The DB is a dependency but it's not the main component, hence it's fine to use a provider so we can focus on the product itself.

After creating the DB instance, the DB private ip has to set on BE deployment.   


## ToDo:   
We could create the SQL cloud instance with terraform. The private IP would be an output, could we set this IP with kustomize on CI pipeline?   

