k8s_yaml(kustomize('kustomize'))

# Build: tell Tilt what images to build from which directories
docker_build('yurifl/logapi', '.')

# Watch: tell Tilt how to connect locally (optional)
k8s_resource(
  'logapi',
  port_forwards=8080
)