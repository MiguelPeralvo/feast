# @generated by generate_proto_mypy_stubs.py.  Do not edit!
import sys
from feast.core.Source_pb2 import (
    Source as feast___core___Source_pb2___Source,
)

from feast.types.Value_pb2 import (
    ValueType as feast___types___Value_pb2___ValueType,
)

from google.protobuf.descriptor import (
    Descriptor as google___protobuf___descriptor___Descriptor,
)

from google.protobuf.duration_pb2 import (
    Duration as google___protobuf___duration_pb2___Duration,
)

from google.protobuf.internal.containers import (
    RepeatedCompositeFieldContainer as google___protobuf___internal___containers___RepeatedCompositeFieldContainer,
)

from google.protobuf.message import (
    Message as google___protobuf___message___Message,
)

from typing import (
    Iterable as typing___Iterable,
    Optional as typing___Optional,
    Text as typing___Text,
)

from typing_extensions import (
    Literal as typing_extensions___Literal,
)


class FeatureSetSpec(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    name = ... # type: typing___Text
    version = ... # type: int

    @property
    def entities(self) -> google___protobuf___internal___containers___RepeatedCompositeFieldContainer[EntitySpec]: ...

    @property
    def features(self) -> google___protobuf___internal___containers___RepeatedCompositeFieldContainer[FeatureSpec]: ...

    @property
    def max_age(self) -> google___protobuf___duration_pb2___Duration: ...

    @property
    def source(self) -> feast___core___Source_pb2___Source: ...

    def __init__(self,
        *,
        name : typing___Optional[typing___Text] = None,
        version : typing___Optional[int] = None,
        entities : typing___Optional[typing___Iterable[EntitySpec]] = None,
        features : typing___Optional[typing___Iterable[FeatureSpec]] = None,
        max_age : typing___Optional[google___protobuf___duration_pb2___Duration] = None,
        source : typing___Optional[feast___core___Source_pb2___Source] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> FeatureSetSpec: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"max_age",u"source"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"entities",u"features",u"max_age",u"name",u"source",u"version"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"max_age",b"max_age",u"source",b"source"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"entities",b"entities",u"features",b"features",u"max_age",b"max_age",u"name",b"name",u"source",b"source",u"version",b"version"]) -> None: ...

class EntitySpec(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    name = ... # type: typing___Text
    value_type = ... # type: feast___types___Value_pb2___ValueType.Enum

    def __init__(self,
        *,
        name : typing___Optional[typing___Text] = None,
        value_type : typing___Optional[feast___types___Value_pb2___ValueType.Enum] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> EntitySpec: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"name",u"value_type"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"name",b"name",u"value_type",b"value_type"]) -> None: ...

class FeatureSpec(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    name = ... # type: typing___Text
    value_type = ... # type: feast___types___Value_pb2___ValueType.Enum

    def __init__(self,
        *,
        name : typing___Optional[typing___Text] = None,
        value_type : typing___Optional[feast___types___Value_pb2___ValueType.Enum] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> FeatureSpec: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"name",u"value_type"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"name",b"name",u"value_type",b"value_type"]) -> None: ...
