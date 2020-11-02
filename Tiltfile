yaml_str = """
apiVersion: v1
kind: Pod
  metadata:
    name: app
    labels:
      app: app
"""
k8s_yaml([
    blob(yaml_str)  # Type: Blob
])

# Build: tell Tilt what images to build from which directories
docker_build('logapi', '.')

# Watch: tell Tilt how to connect locally (optional)
k8s_resource(
  'logapi',
  port_forwards=8080
)
