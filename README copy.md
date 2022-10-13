1. Created simple example of main.go app.
2. Created Dockerfile to build source using multistage builds
3. Created dockercompose.yaml file with services: 
    - task-01
    - nginx
4. nginx.conf file is cofigured to set nginx container as Entry point for a trafic.
5. And the same network was created for both services
6. CI for this project was configured, Pipeline triggers for any changes in main/dev branches
7. Notification
