# grub
a cart management tool written in go and python to make coordinating group orders easy.

It is comprised of two separate RESTful services: a Mux API to provide an interface to a MongoDB Atlas cluster and a Flask web app / API for integrations such as slack. Currently using go routines and the go exec package to manually start the Flask service to interface with slack and serve the UI. However, hoping to move to a GRPC based communication mechanism.

