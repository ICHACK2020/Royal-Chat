# -*- coding: utf-8 -*-
"""
Created on Sat Feb  8 15:03:20 2020

@authors: Akshat, Krishan
"""

from concurrent import futures
from os import chdir
import grpc
import topicFinder

chdir("../python_proto/")
import python_proto.api_pb2 as api_pb2
import python_proto.api_pb2_grpc as api_pb2_grpc

chdir("../relevancy/")
class Listener(api_pb2_grpc.ProcessServicer):

    def __init__(self):
        self.topicFinderObj = topicFinder.TopicFinder()
        print("INIT")

    def Troll(self, request, context):
        print("THESESE")
        raise Exception("Wrong server")

    def Relevance(self, request, context):
        print(request)
        answer = self.topicFinderObj.checkRelevant(request.msg, request.conv_id.lower())
        print(answer)
        return api_pb2.apiResponse(uid=request.uid, conv_id=request.conv_id, score=answer, rolling_score=0)
        # return api_pb2.apiResponse(uid=request.uid, conv_id=request.conv_id, score= 0 ,
        #                            rolling_score=5.3)

def run():
    server = grpc.server(futures.ThreadPoolExecutor(4))
    api_pb2_grpc.add_ProcessServicer_to_server(Listener(), server)
    server.add_insecure_port("146.169.139.247:8081")  # Add your ip address here
    server.start()
    input("Press any button to stop the server ")
    server.stop(0)


if __name__ == "__main__":
    run()
