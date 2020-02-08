# -*- coding: utf-8 -*-
"""
Created on Sat Feb  8 15:03:20 2020

@author: Akshat
"""

from concurrent import futures
from os import chdir
import grpc

chdir("../python_proto/")
import api_pb2
import api_pb2_grpc

class Listener(api_pb2_grpc.ProcessServicer):

    def __init__(self):
        pass

    def Troll(self, request_iterator, context):
        for request in request_iterator:
            yield api_pb2.apiResponse()

    def Relevance(self, request_iterator, context):
        raise Exception("Wrong server")

def run():
    server = grpc.server(futures.ThreadPoolExecutor(4))
    api_pb2_grpc.add_ProcessServicer_to_server(Listener(), server)
    server.add_insecure_port() #Add your ip address here
    server.start()
    input("Press any button to stop the server ")
    server.stop(0)



if __name__ == "__main__":
    run()