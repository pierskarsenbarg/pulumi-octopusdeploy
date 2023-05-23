# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AwsElasticContainerRegistryArgs', 'AwsElasticContainerRegistry']

@pulumi.input_type
class AwsElasticContainerRegistryArgs:
    def __init__(__self__, *,
                 access_key: pulumi.Input[str],
                 region: pulumi.Input[str],
                 secret_key: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 space_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AwsElasticContainerRegistry resource.
        :param pulumi.Input[str] access_key: The AWS access key to use when authenticating against Amazon Web Services.
        :param pulumi.Input[str] region: The AWS region where the registry resides.
        :param pulumi.Input[str] secret_key: The AWS secret key to use when authenticating against Amazon Web Services.
        :param pulumi.Input[str] name: A short, memorable, unique name for this feed. Example: ACME Builds.
        :param pulumi.Input[str] space_id: The space ID associated with this feed.
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "secret_key", secret_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if package_acquisition_location_options is not None:
            pulumi.set(__self__, "package_acquisition_location_options", package_acquisition_location_options)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Input[str]:
        """
        The AWS access key to use when authenticating against Amazon Web Services.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The AWS region where the registry resides.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        The AWS secret key to use when authenticating against Amazon Web Services.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)

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
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this feed.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)


@pulumi.input_type
class _AwsElasticContainerRegistryState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AwsElasticContainerRegistry resources.
        :param pulumi.Input[str] access_key: The AWS access key to use when authenticating against Amazon Web Services.
        :param pulumi.Input[str] name: A short, memorable, unique name for this feed. Example: ACME Builds.
        :param pulumi.Input[str] region: The AWS region where the registry resides.
        :param pulumi.Input[str] secret_key: The AWS secret key to use when authenticating against Amazon Web Services.
        :param pulumi.Input[str] space_id: The space ID associated with this feed.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if package_acquisition_location_options is not None:
            pulumi.set(__self__, "package_acquisition_location_options", package_acquisition_location_options)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS access key to use when authenticating against Amazon Web Services.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

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
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS region where the registry resides.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS secret key to use when authenticating against Amazon Web Services.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this feed.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)


class AwsElasticContainerRegistry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource manages an AWS Elastic Container Registry in Octopus Deploy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.AwsElasticContainerRegistry("example",
            access_key="access-key",
            region="us-east-1",
            secret_key="secret-key")
        ```

        ## Import

        ```sh
         $ pulumi import octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry [options] octopusdeploy_aws_elastic_container_registry.<name> <feed-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The AWS access key to use when authenticating against Amazon Web Services.
        :param pulumi.Input[str] name: A short, memorable, unique name for this feed. Example: ACME Builds.
        :param pulumi.Input[str] region: The AWS region where the registry resides.
        :param pulumi.Input[str] secret_key: The AWS secret key to use when authenticating against Amazon Web Services.
        :param pulumi.Input[str] space_id: The space ID associated with this feed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsElasticContainerRegistryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource manages an AWS Elastic Container Registry in Octopus Deploy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.AwsElasticContainerRegistry("example",
            access_key="access-key",
            region="us-east-1",
            secret_key="secret-key")
        ```

        ## Import

        ```sh
         $ pulumi import octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry [options] octopusdeploy_aws_elastic_container_registry.<name> <feed-id>
        ```

        :param str resource_name: The name of the resource.
        :param AwsElasticContainerRegistryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsElasticContainerRegistryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsElasticContainerRegistryArgs.__new__(AwsElasticContainerRegistryArgs)

            if access_key is None and not opts.urn:
                raise TypeError("Missing required property 'access_key'")
            __props__.__dict__["access_key"] = access_key
            __props__.__dict__["name"] = name
            __props__.__dict__["package_acquisition_location_options"] = package_acquisition_location_options
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            __props__.__dict__["space_id"] = space_id
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AwsElasticContainerRegistry, __self__).__init__(
            'octopusdeploy:index/awsElasticContainerRegistry:AwsElasticContainerRegistry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            package_acquisition_location_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None) -> 'AwsElasticContainerRegistry':
        """
        Get an existing AwsElasticContainerRegistry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The AWS access key to use when authenticating against Amazon Web Services.
        :param pulumi.Input[str] name: A short, memorable, unique name for this feed. Example: ACME Builds.
        :param pulumi.Input[str] region: The AWS region where the registry resides.
        :param pulumi.Input[str] secret_key: The AWS secret key to use when authenticating against Amazon Web Services.
        :param pulumi.Input[str] space_id: The space ID associated with this feed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsElasticContainerRegistryState.__new__(_AwsElasticContainerRegistryState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["name"] = name
        __props__.__dict__["package_acquisition_location_options"] = package_acquisition_location_options
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["space_id"] = space_id
        return AwsElasticContainerRegistry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        The AWS access key to use when authenticating against Amazon Web Services.
        """
        return pulumi.get(self, "access_key")

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
    def region(self) -> pulumi.Output[str]:
        """
        The AWS region where the registry resides.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        The AWS secret key to use when authenticating against Amazon Web Services.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        The space ID associated with this feed.
        """
        return pulumi.get(self, "space_id")

