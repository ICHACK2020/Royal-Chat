# -*- coding: utf-8 -*-
"""
Created on Sat Feb  8 15:03:20 2020

@authors: Akshat, George
"""

from concurrent import futures
from os import chdir
import grpc
from tracker import Detector

chdir("../python_proto")
import python_proto.api_pb2 as api_pb2
import python_proto.api_pb2_grpc as api_pb2_grpc


class Listener(api_pb2_grpc.ProcessServicer):

    def __init__(self):
        self.detector = Detector()

    def Troll(self, request, context):
        self.detector.request(request.msg, request.uid, request.conv_id)
        id = Detector.generate_id(request.conv_id, request.uid)
        score = self.detector.get_recent_score(id)
        rolling_score = self.detector.get_rolling_score(id)

        print('Request recieved:')
        print('Message:\t', request.msg)
        print('Score:\t', score)
        print('Single score metric:', self.detector.get_rolling_score_as_score(id))
        print('\n\n')

        return api_pb2.apiResponse(uid=request.uid, conv_id=request.conv_id, score=score, rolling_score=rolling_score)


    def Relevance(self, request_iterator, context):
        raise Exception("Wrong server")


def run():
    server = grpc.server(futures.ThreadPoolExecutor(4))
    api_pb2_grpc.add_ProcessServicer_to_server(Listener(), server)
    server.add_insecure_port("146.169.157.5:8080")  # Add your ip address here
    server.start()
    input("Press any button to stop the server ")
    server.stop(0)


if __name__ == "__main__":
    run()
