# -*- coding: utf-8 -*-
"""
Created on Sat Feb  8 14:55:01 2020

@author: Akshat
"""

from grpc_tools import protoc

protoc.main((
    '',
    '-I../proto',
    '--python_out=.',
    '--grpc_python_out=.',
    '../proto/api.proto',
))