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
    'GetGitCredentialsResult',
    'AwaitableGetGitCredentialsResult',
    'get_git_credentials',
    'get_git_credentials_output',
]

@pulumi.output_type
class GetGitCredentialsResult:
    """
    A collection of values returned by getGitCredentials.
    """
    def __init__(__self__, git_credentials=None, id=None, name=None, skip=None, take=None):
        if git_credentials and not isinstance(git_credentials, list):
            raise TypeError("Expected argument 'git_credentials' to be a list")
        pulumi.set(__self__, "git_credentials", git_credentials)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if skip and not isinstance(skip, int):
            raise TypeError("Expected argument 'skip' to be a int")
        pulumi.set(__self__, "skip", skip)
        if take and not isinstance(take, int):
            raise TypeError("Expected argument 'take' to be a int")
        pulumi.set(__self__, "take", take)

    @property
    @pulumi.getter(name="gitCredentials")
    def git_credentials(self) -> Sequence['outputs.GetGitCredentialsGitCredentialResult']:
        """
        A list of Git Credentials that match the filter(s).
        """
        return pulumi.get(self, "git_credentials")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        A filter to search by name.
        """
        return pulumi.get(self, "name")

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


class AwaitableGetGitCredentialsResult(GetGitCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGitCredentialsResult(
            git_credentials=self.git_credentials,
            id=self.id,
            name=self.name,
            skip=self.skip,
            take=self.take)


def get_git_credentials(git_credentials: Optional[Sequence[pulumi.InputType['GetGitCredentialsGitCredentialArgs']]] = None,
                        name: Optional[str] = None,
                        skip: Optional[int] = None,
                        take: Optional[int] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGitCredentialsResult:
    """
    Provides information about existing GitCredentials.


    :param Sequence[pulumi.InputType['GetGitCredentialsGitCredentialArgs']] git_credentials: A list of Git Credentials that match the filter(s).
    :param str name: A filter to search by name.
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    __args__ = dict()
    __args__['gitCredentials'] = git_credentials
    __args__['name'] = name
    __args__['skip'] = skip
    __args__['take'] = take
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('octopusdeploy:index/getGitCredentials:getGitCredentials', __args__, opts=opts, typ=GetGitCredentialsResult).value

    return AwaitableGetGitCredentialsResult(
        git_credentials=__ret__.git_credentials,
        id=__ret__.id,
        name=__ret__.name,
        skip=__ret__.skip,
        take=__ret__.take)


@_utilities.lift_output_func(get_git_credentials)
def get_git_credentials_output(git_credentials: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetGitCredentialsGitCredentialArgs']]]]] = None,
                               name: Optional[pulumi.Input[Optional[str]]] = None,
                               skip: Optional[pulumi.Input[Optional[int]]] = None,
                               take: Optional[pulumi.Input[Optional[int]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGitCredentialsResult]:
    """
    Provides information about existing GitCredentials.


    :param Sequence[pulumi.InputType['GetGitCredentialsGitCredentialArgs']] git_credentials: A list of Git Credentials that match the filter(s).
    :param str name: A filter to search by name.
    :param int skip: A filter to specify the number of items to skip in the response.
    :param int take: A filter to specify the number of items to take (or return) in the response.
    """
    ...