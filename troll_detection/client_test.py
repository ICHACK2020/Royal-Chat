# -*- coding: utf-8 -*-
"""
Created on Sun Jan 26 13:22:15 2020

"""

import grpc
import time

from os import chdir

import api_pb2
import api_pb2_grpc

msgs = ["I have never seen anything so despicable before",
        "If you walked into a bar you would hit your head",
        "why are you offended you're such a snowflake",
        "When you speak I want to blow your brains out",
        "wow. genius. fucking brilliant",
        "I hope your son gets cancer",
        "I'm going to trace your location and find you",
        "If you don't believe in God, he will punish you in hell forever",
        "We need to take the country forward. make britain great again. We don't need the EU.",
        "That's so gay",
        "Well no obviously not, I think it's a little more subtle than that"]

def run():
    with grpc.insecure_channel("localhost:8080") as channel:
        stream_stub = api_pb2_grpc.ProcessStub(channel)
        x = []
        #while True:
        try:
            x.append(api_pb2.apiCall(uid=True, conv_id='0', msg='You are a fag'))
            x.append(api_pb2.apiCall(uid=False, conv_id='0', msg='I hope your son gets cancer'))
            x.append(api_pb2.apiCall(uid=True, conv_id='1', msg='What is your position on abortion'))
            x.append(api_pb2.apiCall(uid=False, conv_id='1', msg='religious bullshit'))
            x.append(api_pb2.apiCall(uid=True, conv_id='2', msg='oh yeah because if I\'m white I can\'t have any issues'))
            x.append(api_pb2.apiCall(uid=False, conv_id='2', msg='Jesus people like you are ruining america'))
            for received in stream_stub.Troll(iter(x)):
                print(received)
            time.sleep(1.5)
        except KeyboardInterrupt:
            print("done")
            channel.unsubscribe(channel.close())

if __name__ == "__main__":
    run()
