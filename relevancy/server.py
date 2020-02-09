# -*- coding: utf-8 -*-
"""
Created on Sat Feb  8 15:03:20 2020

@authors: Akshat, George
"""

from concurrent import futures
from os import chdir
import grpc
import topicFinder

chdir("../python_proto/")
import api_pb2
import api_pb2_grpc


class Listener(api_pb2_grpc.ProcessServicer):

    def __init__(self):
        pass

    def Troll(self, request, context):
        raise Exception("Wrong server")

    def Relevance(self, request, context):
        topicFinderObj = topicFinder.TopicFinder()
        answer = topicFinderObj.checkRelvant(request.msg, request.conv_id)
        return api_pb2.apiResponse(uid=request.uid, conv_id=request.conv_id, score=request.score, rolling_score=request.rolling_score, relevance=answer)


def run():
    server = grpc.server(futures.ThreadPoolExecutor(4))
    api_pb2_grpc.add_ProcessServicer_to_server(Listener(), server)
    server.add_insecure_port("146.169.139.247:8080")  # Add your ip address here
    server.start()
    input("Press any button to stop the server ")
    server.stop(0)


if __name__ == "__main__":
    run()
