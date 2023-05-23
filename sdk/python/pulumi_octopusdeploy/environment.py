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

__all__ = ['EnvironmentArgs', 'Environment']

@pulumi.input_type
class EnvironmentArgs:
    def __init__(__self__, *,
                 allow_dynamic_infrastructure: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 jira_extension_settings: Optional[pulumi.Input['EnvironmentJiraExtensionSettingsArgs']] = None,
                 jira_service_management_extension_settings: Optional[pulumi.Input['EnvironmentJiraServiceManagementExtensionSettingsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servicenow_extension_settings: Optional[pulumi.Input['EnvironmentServicenowExtensionSettingsArgs']] = None,
                 sort_order: Optional[pulumi.Input[int]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 use_guided_failure: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Environment resource.
        :param pulumi.Input[str] description: The description of this environment.
        :param pulumi.Input['EnvironmentJiraExtensionSettingsArgs'] jira_extension_settings: Provides extension settings for the Jira integration for this environment.
        :param pulumi.Input['EnvironmentJiraServiceManagementExtensionSettingsArgs'] jira_service_management_extension_settings: Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input['EnvironmentServicenowExtensionSettingsArgs'] servicenow_extension_settings: Provides extension settings for the ServiceNow integration for this environment.
        :param pulumi.Input[int] sort_order: The order number to sort an environment.
        :param pulumi.Input[str] space_id: The space ID associated with this environment.
        """
        if allow_dynamic_infrastructure is not None:
            pulumi.set(__self__, "allow_dynamic_infrastructure", allow_dynamic_infrastructure)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if jira_extension_settings is not None:
            pulumi.set(__self__, "jira_extension_settings", jira_extension_settings)
        if jira_service_management_extension_settings is not None:
            pulumi.set(__self__, "jira_service_management_extension_settings", jira_service_management_extension_settings)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if servicenow_extension_settings is not None:
            pulumi.set(__self__, "servicenow_extension_settings", servicenow_extension_settings)
        if sort_order is not None:
            pulumi.set(__self__, "sort_order", sort_order)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if use_guided_failure is not None:
            pulumi.set(__self__, "use_guided_failure", use_guided_failure)

    @property
    @pulumi.getter(name="allowDynamicInfrastructure")
    def allow_dynamic_infrastructure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allow_dynamic_infrastructure")

    @allow_dynamic_infrastructure.setter
    def allow_dynamic_infrastructure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_dynamic_infrastructure", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this environment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="jiraExtensionSettings")
    def jira_extension_settings(self) -> Optional[pulumi.Input['EnvironmentJiraExtensionSettingsArgs']]:
        """
        Provides extension settings for the Jira integration for this environment.
        """
        return pulumi.get(self, "jira_extension_settings")

    @jira_extension_settings.setter
    def jira_extension_settings(self, value: Optional[pulumi.Input['EnvironmentJiraExtensionSettingsArgs']]):
        pulumi.set(self, "jira_extension_settings", value)

    @property
    @pulumi.getter(name="jiraServiceManagementExtensionSettings")
    def jira_service_management_extension_settings(self) -> Optional[pulumi.Input['EnvironmentJiraServiceManagementExtensionSettingsArgs']]:
        """
        Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        """
        return pulumi.get(self, "jira_service_management_extension_settings")

    @jira_service_management_extension_settings.setter
    def jira_service_management_extension_settings(self, value: Optional[pulumi.Input['EnvironmentJiraServiceManagementExtensionSettingsArgs']]):
        pulumi.set(self, "jira_service_management_extension_settings", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="servicenowExtensionSettings")
    def servicenow_extension_settings(self) -> Optional[pulumi.Input['EnvironmentServicenowExtensionSettingsArgs']]:
        """
        Provides extension settings for the ServiceNow integration for this environment.
        """
        return pulumi.get(self, "servicenow_extension_settings")

    @servicenow_extension_settings.setter
    def servicenow_extension_settings(self, value: Optional[pulumi.Input['EnvironmentServicenowExtensionSettingsArgs']]):
        pulumi.set(self, "servicenow_extension_settings", value)

    @property
    @pulumi.getter(name="sortOrder")
    def sort_order(self) -> Optional[pulumi.Input[int]]:
        """
        The order number to sort an environment.
        """
        return pulumi.get(self, "sort_order")

    @sort_order.setter
    def sort_order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sort_order", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this environment.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)

    @property
    @pulumi.getter(name="useGuidedFailure")
    def use_guided_failure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_guided_failure")

    @use_guided_failure.setter
    def use_guided_failure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_guided_failure", value)


@pulumi.input_type
class _EnvironmentState:
    def __init__(__self__, *,
                 allow_dynamic_infrastructure: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 jira_extension_settings: Optional[pulumi.Input['EnvironmentJiraExtensionSettingsArgs']] = None,
                 jira_service_management_extension_settings: Optional[pulumi.Input['EnvironmentJiraServiceManagementExtensionSettingsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servicenow_extension_settings: Optional[pulumi.Input['EnvironmentServicenowExtensionSettingsArgs']] = None,
                 slug: Optional[pulumi.Input[str]] = None,
                 sort_order: Optional[pulumi.Input[int]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 use_guided_failure: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Environment resources.
        :param pulumi.Input[str] description: The description of this environment.
        :param pulumi.Input['EnvironmentJiraExtensionSettingsArgs'] jira_extension_settings: Provides extension settings for the Jira integration for this environment.
        :param pulumi.Input['EnvironmentJiraServiceManagementExtensionSettingsArgs'] jira_service_management_extension_settings: Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input['EnvironmentServicenowExtensionSettingsArgs'] servicenow_extension_settings: Provides extension settings for the ServiceNow integration for this environment.
        :param pulumi.Input[int] sort_order: The order number to sort an environment.
        :param pulumi.Input[str] space_id: The space ID associated with this environment.
        """
        if allow_dynamic_infrastructure is not None:
            pulumi.set(__self__, "allow_dynamic_infrastructure", allow_dynamic_infrastructure)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if jira_extension_settings is not None:
            pulumi.set(__self__, "jira_extension_settings", jira_extension_settings)
        if jira_service_management_extension_settings is not None:
            pulumi.set(__self__, "jira_service_management_extension_settings", jira_service_management_extension_settings)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if servicenow_extension_settings is not None:
            pulumi.set(__self__, "servicenow_extension_settings", servicenow_extension_settings)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)
        if sort_order is not None:
            pulumi.set(__self__, "sort_order", sort_order)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if use_guided_failure is not None:
            pulumi.set(__self__, "use_guided_failure", use_guided_failure)

    @property
    @pulumi.getter(name="allowDynamicInfrastructure")
    def allow_dynamic_infrastructure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allow_dynamic_infrastructure")

    @allow_dynamic_infrastructure.setter
    def allow_dynamic_infrastructure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_dynamic_infrastructure", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this environment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="jiraExtensionSettings")
    def jira_extension_settings(self) -> Optional[pulumi.Input['EnvironmentJiraExtensionSettingsArgs']]:
        """
        Provides extension settings for the Jira integration for this environment.
        """
        return pulumi.get(self, "jira_extension_settings")

    @jira_extension_settings.setter
    def jira_extension_settings(self, value: Optional[pulumi.Input['EnvironmentJiraExtensionSettingsArgs']]):
        pulumi.set(self, "jira_extension_settings", value)

    @property
    @pulumi.getter(name="jiraServiceManagementExtensionSettings")
    def jira_service_management_extension_settings(self) -> Optional[pulumi.Input['EnvironmentJiraServiceManagementExtensionSettingsArgs']]:
        """
        Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        """
        return pulumi.get(self, "jira_service_management_extension_settings")

    @jira_service_management_extension_settings.setter
    def jira_service_management_extension_settings(self, value: Optional[pulumi.Input['EnvironmentJiraServiceManagementExtensionSettingsArgs']]):
        pulumi.set(self, "jira_service_management_extension_settings", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="servicenowExtensionSettings")
    def servicenow_extension_settings(self) -> Optional[pulumi.Input['EnvironmentServicenowExtensionSettingsArgs']]:
        """
        Provides extension settings for the ServiceNow integration for this environment.
        """
        return pulumi.get(self, "servicenow_extension_settings")

    @servicenow_extension_settings.setter
    def servicenow_extension_settings(self, value: Optional[pulumi.Input['EnvironmentServicenowExtensionSettingsArgs']]):
        pulumi.set(self, "servicenow_extension_settings", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)

    @property
    @pulumi.getter(name="sortOrder")
    def sort_order(self) -> Optional[pulumi.Input[int]]:
        """
        The order number to sort an environment.
        """
        return pulumi.get(self, "sort_order")

    @sort_order.setter
    def sort_order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sort_order", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The space ID associated with this environment.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)

    @property
    @pulumi.getter(name="useGuidedFailure")
    def use_guided_failure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_guided_failure")

    @use_guided_failure.setter
    def use_guided_failure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_guided_failure", value)


class Environment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_dynamic_infrastructure: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 jira_extension_settings: Optional[pulumi.Input[pulumi.InputType['EnvironmentJiraExtensionSettingsArgs']]] = None,
                 jira_service_management_extension_settings: Optional[pulumi.Input[pulumi.InputType['EnvironmentJiraServiceManagementExtensionSettingsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servicenow_extension_settings: Optional[pulumi.Input[pulumi.InputType['EnvironmentServicenowExtensionSettingsArgs']]] = None,
                 sort_order: Optional[pulumi.Input[int]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 use_guided_failure: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        This resource manages environments in Octopus Deploy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.Environment("example",
            allow_dynamic_infrastructure=False,
            description="An environment for the development team.",
            jira_extension_settings=octopusdeploy.EnvironmentJiraExtensionSettingsArgs(
                environment_type="unmapped",
            ),
            jira_service_management_extension_settings=octopusdeploy.EnvironmentJiraServiceManagementExtensionSettingsArgs(
                is_enabled=False,
            ),
            servicenow_extension_settings=octopusdeploy.EnvironmentServicenowExtensionSettingsArgs(
                is_enabled=False,
            ),
            use_guided_failure=False)
        ```

        ## Import

        ```sh
         $ pulumi import octopusdeploy:index/environment:Environment [options] octopusdeploy_environment.<name> <environment-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this environment.
        :param pulumi.Input[pulumi.InputType['EnvironmentJiraExtensionSettingsArgs']] jira_extension_settings: Provides extension settings for the Jira integration for this environment.
        :param pulumi.Input[pulumi.InputType['EnvironmentJiraServiceManagementExtensionSettingsArgs']] jira_service_management_extension_settings: Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[pulumi.InputType['EnvironmentServicenowExtensionSettingsArgs']] servicenow_extension_settings: Provides extension settings for the ServiceNow integration for this environment.
        :param pulumi.Input[int] sort_order: The order number to sort an environment.
        :param pulumi.Input[str] space_id: The space ID associated with this environment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EnvironmentArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource manages environments in Octopus Deploy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_octopusdeploy as octopusdeploy

        example = octopusdeploy.Environment("example",
            allow_dynamic_infrastructure=False,
            description="An environment for the development team.",
            jira_extension_settings=octopusdeploy.EnvironmentJiraExtensionSettingsArgs(
                environment_type="unmapped",
            ),
            jira_service_management_extension_settings=octopusdeploy.EnvironmentJiraServiceManagementExtensionSettingsArgs(
                is_enabled=False,
            ),
            servicenow_extension_settings=octopusdeploy.EnvironmentServicenowExtensionSettingsArgs(
                is_enabled=False,
            ),
            use_guided_failure=False)
        ```

        ## Import

        ```sh
         $ pulumi import octopusdeploy:index/environment:Environment [options] octopusdeploy_environment.<name> <environment-id>
        ```

        :param str resource_name: The name of the resource.
        :param EnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_dynamic_infrastructure: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 jira_extension_settings: Optional[pulumi.Input[pulumi.InputType['EnvironmentJiraExtensionSettingsArgs']]] = None,
                 jira_service_management_extension_settings: Optional[pulumi.Input[pulumi.InputType['EnvironmentJiraServiceManagementExtensionSettingsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 servicenow_extension_settings: Optional[pulumi.Input[pulumi.InputType['EnvironmentServicenowExtensionSettingsArgs']]] = None,
                 sort_order: Optional[pulumi.Input[int]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 use_guided_failure: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

            __props__.__dict__["allow_dynamic_infrastructure"] = allow_dynamic_infrastructure
            __props__.__dict__["description"] = description
            __props__.__dict__["jira_extension_settings"] = jira_extension_settings
            __props__.__dict__["jira_service_management_extension_settings"] = jira_service_management_extension_settings
            __props__.__dict__["name"] = name
            __props__.__dict__["servicenow_extension_settings"] = servicenow_extension_settings
            __props__.__dict__["sort_order"] = sort_order
            __props__.__dict__["space_id"] = space_id
            __props__.__dict__["use_guided_failure"] = use_guided_failure
            __props__.__dict__["slug"] = None
        super(Environment, __self__).__init__(
            'octopusdeploy:index/environment:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_dynamic_infrastructure: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            jira_extension_settings: Optional[pulumi.Input[pulumi.InputType['EnvironmentJiraExtensionSettingsArgs']]] = None,
            jira_service_management_extension_settings: Optional[pulumi.Input[pulumi.InputType['EnvironmentJiraServiceManagementExtensionSettingsArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            servicenow_extension_settings: Optional[pulumi.Input[pulumi.InputType['EnvironmentServicenowExtensionSettingsArgs']]] = None,
            slug: Optional[pulumi.Input[str]] = None,
            sort_order: Optional[pulumi.Input[int]] = None,
            space_id: Optional[pulumi.Input[str]] = None,
            use_guided_failure: Optional[pulumi.Input[bool]] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this environment.
        :param pulumi.Input[pulumi.InputType['EnvironmentJiraExtensionSettingsArgs']] jira_extension_settings: Provides extension settings for the Jira integration for this environment.
        :param pulumi.Input[pulumi.InputType['EnvironmentJiraServiceManagementExtensionSettingsArgs']] jira_service_management_extension_settings: Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[pulumi.InputType['EnvironmentServicenowExtensionSettingsArgs']] servicenow_extension_settings: Provides extension settings for the ServiceNow integration for this environment.
        :param pulumi.Input[int] sort_order: The order number to sort an environment.
        :param pulumi.Input[str] space_id: The space ID associated with this environment.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnvironmentState.__new__(_EnvironmentState)

        __props__.__dict__["allow_dynamic_infrastructure"] = allow_dynamic_infrastructure
        __props__.__dict__["description"] = description
        __props__.__dict__["jira_extension_settings"] = jira_extension_settings
        __props__.__dict__["jira_service_management_extension_settings"] = jira_service_management_extension_settings
        __props__.__dict__["name"] = name
        __props__.__dict__["servicenow_extension_settings"] = servicenow_extension_settings
        __props__.__dict__["slug"] = slug
        __props__.__dict__["sort_order"] = sort_order
        __props__.__dict__["space_id"] = space_id
        __props__.__dict__["use_guided_failure"] = use_guided_failure
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowDynamicInfrastructure")
    def allow_dynamic_infrastructure(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "allow_dynamic_infrastructure")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of this environment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="jiraExtensionSettings")
    def jira_extension_settings(self) -> pulumi.Output[Optional['outputs.EnvironmentJiraExtensionSettings']]:
        """
        Provides extension settings for the Jira integration for this environment.
        """
        return pulumi.get(self, "jira_extension_settings")

    @property
    @pulumi.getter(name="jiraServiceManagementExtensionSettings")
    def jira_service_management_extension_settings(self) -> pulumi.Output[Optional['outputs.EnvironmentJiraServiceManagementExtensionSettings']]:
        """
        Provides extension settings for the Jira Service Management (JSM) integration for this environment.
        """
        return pulumi.get(self, "jira_service_management_extension_settings")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="servicenowExtensionSettings")
    def servicenow_extension_settings(self) -> pulumi.Output[Optional['outputs.EnvironmentServicenowExtensionSettings']]:
        """
        Provides extension settings for the ServiceNow integration for this environment.
        """
        return pulumi.get(self, "servicenow_extension_settings")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter(name="sortOrder")
    def sort_order(self) -> pulumi.Output[int]:
        """
        The order number to sort an environment.
        """
        return pulumi.get(self, "sort_order")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        The space ID associated with this environment.
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter(name="useGuidedFailure")
    def use_guided_failure(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "use_guided_failure")
