# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetChannelsResult',
    'AwaitableGetChannelsResult',
    'get_channels',
    'get_channels_output',
]

@pulumi.output_type
class GetChannelsResult:
    """
    A collection of values returned by getChannels.
    """
    def __init__(__self__, channels=None, id=None, ids=None, partial_name=None, skip=None, take=None):
        if channels and not isinstance(channels, list):
            raise TypeError("Expected argument 'channels' to be a list")
        pulumi.set(__self__, "channels", channels)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if partial_name and not isinstance(partial_name, str):
            raise TypeError("Expected argument 'partial_name' to be a str")
        pulumi.set(__self__, "partial_name", partial_name)
        if skip and not isinstance(skip, int):
            raise TypeError("Expected argument 'skip' to be a int")
        pulumi.set(__self__, "skip", skip)
        if take and not isinstance(take, int):
            raise TypeError("Expected argument 'take' to be a int")
        pulumi.set(__self__, "take", take)

    @property
    @pulumi.getter
    def channels(self) -> Sequence['outputs.GetChannelsChannelResult']:
        """
        A channel that matches the specified filter(s).
        """
        return pulumi.get(self, "channels")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Optional[Sequence[str]]:
        """
        A filter to search by a list of IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="partialName")
    def partial_name(self) -> Optional[str]:
        """
        A filter to search by the partial match of a name.
        """
        return pulumi.get(self, "partial_name")

    @property
    @pulumi.getter
    def skip(self) -> Optional[int]:
        """
        A filter to specify the number of items to skip in the response.
        """
        return pulumi.get(self, "skip")

    @property
    @pulumi.getter
    def take(self) -> Optional[int]:
        """
        A filter to specify the number of items to take (or return) in the response.
        """
        return pulumi.get(self, "take")


class AwaitableGetChannelsResult(GetChannelsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetChannelsResult(
            channels=self.channels,
            id=self.id,
            ids=self.ids,
            partial_name=self.partial_name,
            skip=self.skip,
            take=self.take)


def get_channels(channels: Optional[Sequence[pulumi.InputType['GetChannelsChannelArgs']]] = None,
                 ids: Optional[Sequence[str]] = None,
                 partial_name: Optional[str] = None,
                 skip: Optional[int] = None,
                 take: Optional[int] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetChannelsResult:
    """
    Provides information about existing channels.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_octopusdeploy as octopusdeploy

    example = octopusdeploy.get_channels(ids=[
            "Channels-123",
            "Channels-321",
        ],
        partial_name="Defau",
        skip=5,
        take=100)
    ```


    :param Sequence[pulumi.InputType['GetChannelsChannelArgs']] channels: A channel that matches the specified filter(s).
    :param Sequence[str] ids: A filter to search by a list of IDs.
    :param str partial_name: A filter to search by the partial match of a name.
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    __args__ = dict()
    __args__['channels'] = channels
    __args__['ids'] = ids
    __args__['partialName'] = partial_name
    __args__['skip'] = skip
    __args__['take'] = take
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('octopusdeploy:index/getChannels:getChannels', __args__, opts=opts, typ=GetChannelsResult).value

    return AwaitableGetChannelsResult(
        channels=__ret__.channels,
        id=__ret__.id,
        ids=__ret__.ids,
        partial_name=__ret__.partial_name,
        skip=__ret__.skip,
        take=__ret__.take)


@_utilities.lift_output_func(get_channels)
def get_channels_output(channels: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetChannelsChannelArgs']]]]] = None,
                        ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        partial_name: Optional[pulumi.Input[Optional[str]]] = None,
                        skip: Optional[pulumi.Input[Optional[int]]] = None,
                        take: Optional[pulumi.Input[Optional[int]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetChannelsResult]:
    """
    Provides information about existing channels.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_octopusdeploy as octopusdeploy

    example = octopusdeploy.get_channels(ids=[
            "Channels-123",
            "Channels-321",
        ],
        partial_name="Defau",
        skip=5,
        take=100)
    ```


    :param Sequence[pulumi.InputType['GetChannelsChannelArgs']] channels: A channel that matches the specified filter(s).
    :param Sequence[str] ids: A filter to search by a list of IDs.
    :param str partial_name: A filter to search by the partial match of a name.
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    ...
