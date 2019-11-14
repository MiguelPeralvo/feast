# @generated by generate_proto_mypy_stubs.py.  Do not edit!
import sys
from feast.types.Value_pb2 import Value as feast___types___Value_pb2___Value

from google.protobuf.descriptor import (
    Descriptor as google___protobuf___descriptor___Descriptor,
)

from google.protobuf.message import Message as google___protobuf___message___Message

from typing import Optional as typing___Optional, Text as typing___Text

from typing_extensions import Literal as typing_extensions___Literal

class Feature(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    name = ...  # type: typing___Text
    @property
    def value(self) -> feast___types___Value_pb2___Value: ...
    def __init__(
        self,
        *,
        value: typing___Optional[feast___types___Value_pb2___Value] = None,
        name: typing___Optional[typing___Text] = None,
    ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> Feature: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(
            self, field_name: typing_extensions___Literal["value"]
        ) -> bool: ...
        def ClearField(
            self, field_name: typing_extensions___Literal["name", "value"]
        ) -> None: ...
    else:
        def HasField(
            self, field_name: typing_extensions___Literal["value", b"value"]
        ) -> bool: ...
        def ClearField(
            self,
            field_name: typing_extensions___Literal["name", b"name", "value", b"value"],
        ) -> None: ...
