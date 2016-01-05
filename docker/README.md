


### Push Image
Docker is instace-based now . If you just rebuild with the same repository name and tag , it will not override the image on private hub(not sure for Docker Hub) , You need to pull it first , rebuild , and then push it.
=======
## NOTE for Docker

### Private Docker registry

Same name and tag will not be override on registry unless you pull the image first. It means now it's instance based , not tag based for registry.
