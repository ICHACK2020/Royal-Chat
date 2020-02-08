# -*- coding: utf-8 -*-
"""
Created on Sun Jan 26 13:22:15 2020

@author: Akshat
"""

import grpc
import time

from os import chdir

chdir("..")
import service_pb2
import service_pb2_grpc

def run():
    with grpc.insecure_channel("localhost:8080") as channel:
        stub = service_pb2_grpc.ProcessStub(channel)
        stream_stub = service_pb2_grpc.OrdChrStub(channel)
        x = []
        while True:
            try:
                response = stub.Predict(service_pb2.Incoming(msg="hello there"))
                for received in stream_stub.Conv(iter(x)):
                    print(received.char)
                for i in "hello there":
                    print(i)
                    x.append(service_pb2.Ord(num=ord(i)))
                # print(response.relevance, response.sentiment)
                time.sleep(0.01)
            except KeyboardInterrupt:
                print("done")
                channel.unsubscribe(channel.close())

if __name__ == "__main__":
    run()
