FROM golang:1.16-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .
EXPOSE 5200
CMD ["/app/main"]

# docker build -t goapi .
# docker image ls
# docker run --name goapi -p 5200:5200 goapi
# docker tag goapi msitproject.azurecr.io/goapi:1.0.2
# az login
# az acr login --name msitproject
# docker push msitproject.azurecr.io/goapi:1.0.2
# az acr repository list -n msitproject
# az webapp create --resource-group MSITPROJECT --plan MSITPROJECT --name msitapi --deployment-container-image-name msitproject.azurecr.io/goapi:latest
# az webapp config appsettings set --resource-group MSITPROJECT --name msitapi --settings WEBSITES_PORT=5200
# az webapp identity assign --resource-group MSITPROJECT --name msitapi --query principalId --output tsv
# az account show --query id --output tsv

#principalId cf2aa752-7ebc-4a5b-8b7f-ce06dc51a99c
#account 7fce95a9-685d-4133-86c0-a312c1e9958e

# az role assignment create --assignee cf2aa752-7ebc-4a5b-8b7f-ce06dc51a99c --scope /subscriptions/7fce95a9-685d-4133-86c0-a312c1e9958e/resourceGroups/MSITPROJECT/providers/Microsoft.ContainerRegistry/registries/MSITPROJECT --role "AcrPull"

# az resource update --ids /subscriptions/7fce95a9-685d-4133-86c0-a312c1e9958e/resourceGroups/msitproject/providers/Microsoft.Web/sites/msitapi/config/web --set properties.acrUseManagedIdentityCreds=True

# az webapp config container set --name msitapi --resource-group MSITPROJECT --docker-custom-image-name msitproject.azurecr.io/goapi:latest --docker-registry-server-url msitproject.azurecr.io

# az webapp log config --name msitapi --resource-group MSITPROJECT --docker-container-logging filesystem

# az webapp log tail --name msitapi --resource-group MSITPROJECT

# az webapp deployment container config --enable-cd true --name msitapi --resource-group MSITPROJECT --query CI_CD_URL --output tsv

# https://$msitapi:Xl5DrnYct59vZ3zWn7LEJWrEWF17gQZyF0oLx43XZupN4qw8G7ZWjoAyPTLw@msitapi.scm.azurewebsites.net/docker/hook

# az acr webhook create --name msitapi --registry msitproject --uri 'https://$msitapi:Xl5DrnYct59vZ3zWn7LEJWrEWF17gQZyF0oLx43XZupN4qw8G7ZWjoAyPTLw@msitapi.scm.azurewebsites.net/docker/hook' --actions push --scope msitproject.azurecr.io/goapi:latest

# eventId=$(az acr webhook ping --name appserviceCD --registry msitproject --query id --output tsv)
# az acr webhook list-events --name appserviceCD --registry msitproject --query "[?id=='68282ae6-989f-4b1c-baa0-9a77eb6ff7a4'].eventResponseMessage"


# https://docs.microsoft.com/en-us/azure/container-registry/container-registry-get-started-docker-cli?tabs=azure-cli

# https://docs.microsoft.com/en-us/azure/app-service/configure-custom-container?pivots=container-linux

# https://docs.microsoft.com/en-us/azure/app-service/tutorial-custom-container?pivots=container-linux