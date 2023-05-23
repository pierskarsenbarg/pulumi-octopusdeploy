# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GithubRepositoryFeedArgs', 'GithubRepositoryFeed']

@pulumi.input_type
class GithubRepositoryFeedArgs:
    def __init__(__self__, *,
                 feed_uri: pulumi.Input[str],
                 download_attempts: Optional[pulumi.Input[int]] = None,
                 download_retry_backoff_seconds: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GithubRepositoryFeed resource.
        :param pulumi.Input[int] download_attempts: The number of times a deployment should attempt to download a package from this feed before failing.
        :param pulumi.Input[int] download_retry_backoff_seconds: The number of seconds to apply as a linear back off between download attempts.
        :param pulumi.Input[str] name: A short, memorable, unique name for this feed. Example: ACME Builds.
        :param pulumi.Input[str] password: The password associated with this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[str] username: The username associated with this resource.
        """
        pulumi.set(__self__, "feed_uri", feed_uri)
        if download_attempts is not None:
            pulumi.set(__self__, "download_attempts", download_attempts)
        if download_retry_backoff_seconds is not None:
            pulumi.set(__self__, "download_retry_backoff_seconds", download_retry_backoff_seconds)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if package_acquisition_location_options is not None:
            pulumi.set(__self__, "package_acquisition_location_options", package_acquisition_location_options)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="feedUri")
    def feed_uri(self) -> pulumi.Input[str]:
        return pulumi.get(self, "feed_uri")

    @feed_uri.setter
    def feed_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "feed_uri", value)

    @property
    @pulumi.getter(name="downloadAttempts")
    def download_attempts(self) -> Optional[pulumi.Input[int]]:
        """
        The number of times a deployment should attempt to download a package from this feed before failing.
        """
        return pulumi.get(self, "download_attempts")

    @download_attempts.setter
    def download_attempts(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "download_attempts", value)

    @property
    @pulumi.getter(name="downloadRetryBackoffSeconds")
    def download_retry_backoff_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds to apply as a linear back off between download attempts.
        """
        return pulumi.get(self, "download_retry_backoff_seconds")

    @download_retry_backoff_seconds.setter
    def download_retry_backoff_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "download_retry_backoff_seconds", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A short, memorable, unique name for this feed. Example: ACME Builds.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="packageAcquisitionLocationOptions")
    def package_acquisition_location_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "package_acquisition_location_options")

    @package_acquisition_location_options.setter
    def package_acquisition_location_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "package_acquisition_location_options", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password associated with this resource.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this resource.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username associated with this resource.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _GithubRepositoryFeedState:
    def __init__(__self__, *,
                 download_attempts: Optional[pulumi.Input[int]] = None,
                 download_retry_backoff_seconds: Optional[pulumi.Input[int]] = None,
                 feed_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GithubRepositoryFeed resources.
        :param pulumi.Input[int] download_attempts: The number of times a deployment should attempt to download a package from this feed before failing.
        :param pulumi.Input[int] download_retry_backoff_seconds: The number of seconds to apply as a linear back off between download attempts.
        :param pulumi.Input[str] name: A short, memorable, unique name for this feed. Example: ACME Builds.
        :param pulumi.Input[str] password: The password associated with this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[str] username: The username associated with this resource.
        """
        if download_attempts is not None:
            pulumi.set(__self__, "download_attempts", download_attempts)
        if download_retry_backoff_seconds is not None:
            pulumi.set(__self__, "download_retry_backoff_seconds", download_retry_backoff_seconds)
        if feed_uri is not None:
            pulumi.set(__self__, "feed_uri", feed_uri)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if package_acquisition_location_options is not None:
            pulumi.set(__self__, "package_acquisition_location_options", package_acquisition_location_options)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="downloadAttempts")
    def download_attempts(self) -> Optional[pulumi.Input[int]]:
        """
        The number of times a deployment should attempt to download a package from this feed before failing.
        """
        return pulumi.get(self, "download_attempts")

    @download_attempts.setter
    def download_attempts(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "download_attempts", value)

    @property
    @pulumi.getter(name="downloadRetryBackoffSeconds")
    def download_retry_backoff_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds to apply as a linear back off between download attempts.
        """
        return pulumi.get(self, "download_retry_backoff_seconds")

    @download_retry_backoff_seconds.setter
    def download_retry_backoff_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "download_retry_backoff_seconds", value)

    @property
    @pulumi.getter(name="feedUri")
    def feed_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "feed_uri")

    @feed_uri.setter
    def feed_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feed_uri", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A short, memorable, unique name for this feed. Example: ACME Builds.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="packageAcquisitionLocationOptions")
    def package_acquisition_location_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "package_acquisition_location_options")

    @package_acquisition_location_options.setter
    def package_acquisition_location_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "package_acquisition_location_options", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password associated with this resource.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this resource.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username associated with this resource.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class GithubRepositoryFeed(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 download_attempts: Optional[pulumi.Input[int]] = None,
                 download_retry_backoff_seconds: Optional[pulumi.Input[int]] = None,
                 feed_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource manages a GitHub repository feed in Octopus Deploy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.GithubRepositoryFeed("example",
            download_attempts=1,
            download_retry_backoff_seconds=30,
            feed_uri="https://api.github.com",
            password="test-password",
            username="test-username")
        ```

        ## Import

        ```sh
         $ pulumi import octopusdeploy:index/githubRepositoryFeed:GithubRepositoryFeed [options] octopusdeploy_github_repository_feed.<name> <feed-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] download_attempts: The number of times a deployment should attempt to download a package from this feed before failing.
        :param pulumi.Input[int] download_retry_backoff_seconds: The number of seconds to apply as a linear back off between download attempts.
        :param pulumi.Input[str] name: A short, memorable, unique name for this feed. Example: ACME Builds.
        :param pulumi.Input[str] password: The password associated with this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[str] username: The username associated with this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GithubRepositoryFeedArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource manages a GitHub repository feed in Octopus Deploy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.GithubRepositoryFeed("example",
            download_attempts=1,
            download_retry_backoff_seconds=30,
            feed_uri="https://api.github.com",
            password="test-password",
            username="test-username")
        ```

        ## Import

        ```sh
         $ pulumi import octopusdeploy:index/githubRepositoryFeed:GithubRepositoryFeed [options] octopusdeploy_github_repository_feed.<name> <feed-id>
        ```

        :param str resource_name: The name of the resource.
        :param GithubRepositoryFeedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GithubRepositoryFeedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 download_attempts: Optional[pulumi.Input[int]] = None,
                 download_retry_backoff_seconds: Optional[pulumi.Input[int]] = None,
                 feed_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GithubRepositoryFeedArgs.__new__(GithubRepositoryFeedArgs)

            __props__.__dict__["download_attempts"] = download_attempts
            __props__.__dict__["download_retry_backoff_seconds"] = download_retry_backoff_seconds
            if feed_uri is None and not opts.urn:
                raise TypeError("Missing required property 'feed_uri'")
            __props__.__dict__["feed_uri"] = feed_uri
            __props__.__dict__["name"] = name
            __props__.__dict__["package_acquisition_location_options"] = package_acquisition_location_options
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["space_id"] = space_id
            __props__.__dict__["username"] = None if username is None else pulumi.Output.secret(username)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password", "username"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(GithubRepositoryFeed, __self__).__init__(
            'octopusdeploy:index/githubRepositoryFeed:GithubRepositoryFeed',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            download_attempts: Optional[pulumi.Input[int]] = None,
            download_retry_backoff_seconds: Optional[pulumi.Input[int]] = None,
            feed_uri: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            password: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'GithubRepositoryFeed':
        """
        Get an existing GithubRepositoryFeed resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] download_attempts: The number of times a deployment should attempt to download a package from this feed before failing.
        :param pulumi.Input[int] download_retry_backoff_seconds: The number of seconds to apply as a linear back off between download attempts.
        :param pulumi.Input[str] name: A short, memorable, unique name for this feed. Example: ACME Builds.
        :param pulumi.Input[str] password: The password associated with this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        :param pulumi.Input[str] username: The username associated with this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GithubRepositoryFeedState.__new__(_GithubRepositoryFeedState)

        __props__.__dict__["download_attempts"] = download_attempts
        __props__.__dict__["download_retry_backoff_seconds"] = download_retry_backoff_seconds
        __props__.__dict__["feed_uri"] = feed_uri
        __props__.__dict__["name"] = name
        __props__.__dict__["package_acquisition_location_options"] = package_acquisition_location_options
        __props__.__dict__["password"] = password
        __props__.__dict__["space_id"] = space_id
        __props__.__dict__["username"] = username
        return GithubRepositoryFeed(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="downloadAttempts")
    def download_attempts(self) -> pulumi.Output[Optional[int]]:
        """
        The number of times a deployment should attempt to download a package from this feed before failing.
        """
        return pulumi.get(self, "download_attempts")

    @property
    @pulumi.getter(name="downloadRetryBackoffSeconds")
    def download_retry_backoff_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The number of seconds to apply as a linear back off between download attempts.
        """
        return pulumi.get(self, "download_retry_backoff_seconds")

    @property
    @pulumi.getter(name="feedUri")
    def feed_uri(self) -> pulumi.Output[str]:
        return pulumi.get(self, "feed_uri")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A short, memorable, unique name for this feed. Example: ACME Builds.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="packageAcquisitionLocationOptions")
    def package_acquisition_location_options(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "package_acquisition_location_options")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password associated with this resource.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        The space ID associated with this resource.
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        The username associated with this resource.
        """
        return pulumi.get(self, "username")

