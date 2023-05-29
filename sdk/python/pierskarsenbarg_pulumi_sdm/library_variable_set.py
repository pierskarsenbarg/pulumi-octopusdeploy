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

__all__ = ['LibraryVariableSetArgs', 'LibraryVariableSet']

@pulumi.input_type
class LibraryVariableSetArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 templates: Optional[pulumi.Input[Sequence[pulumi.Input['LibraryVariableSetTemplateArgs']]]] = None):
        """
        The set of arguments for constructing a LibraryVariableSet resource.
        :param pulumi.Input[str] description: The description of this library variable set.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if templates is not None:
            pulumi.set(__self__, "templates", templates)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this library variable set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    def templates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LibraryVariableSetTemplateArgs']]]]:
        return pulumi.get(self, "templates")

    @templates.setter
    def templates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LibraryVariableSetTemplateArgs']]]]):
        pulumi.set(self, "templates", value)


@pulumi.input_type
class _LibraryVariableSetState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 templates: Optional[pulumi.Input[Sequence[pulumi.Input['LibraryVariableSetTemplateArgs']]]] = None,
                 variable_set_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LibraryVariableSet resources.
        :param pulumi.Input[str] description: The description of this library variable set.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)
        if templates is not None:
            pulumi.set(__self__, "templates", templates)
        if variable_set_id is not None:
            pulumi.set(__self__, "variable_set_id", variable_set_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this library variable set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    def templates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LibraryVariableSetTemplateArgs']]]]:
        return pulumi.get(self, "templates")

    @templates.setter
    def templates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LibraryVariableSetTemplateArgs']]]]):
        pulumi.set(self, "templates", value)

    @property
    @pulumi.getter(name="variableSetId")
    def variable_set_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "variable_set_id")

    @variable_set_id.setter
    def variable_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variable_set_id", value)


class LibraryVariableSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 templates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LibraryVariableSetTemplateArgs']]]]] = None,
                 __props__=None):
        """
        This resource manages library variable sets in Octopus Deploy.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this library variable set.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LibraryVariableSetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource manages library variable sets in Octopus Deploy.

        :param str resource_name: The name of the resource.
        :param LibraryVariableSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LibraryVariableSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 templates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LibraryVariableSetTemplateArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LibraryVariableSetArgs.__new__(LibraryVariableSetArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["space_id"] = space_id
            __props__.__dict__["templates"] = templates
            __props__.__dict__["variable_set_id"] = None
        super(LibraryVariableSet, __self__).__init__(
            'octopusdeploy:index/libraryVariableSet:LibraryVariableSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None,
            templates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LibraryVariableSetTemplateArgs']]]]] = None,
            variable_set_id: Optional[pulumi.Input[str]] = None) -> 'LibraryVariableSet':
        """
        Get an existing LibraryVariableSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this library variable set.
        :param pulumi.Input[str] name: The name of this resource.
        :param pulumi.Input[str] space_id: The space ID associated with this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LibraryVariableSetState.__new__(_LibraryVariableSetState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["space_id"] = space_id
        __props__.__dict__["templates"] = templates
        __props__.__dict__["variable_set_id"] = variable_set_id
        return LibraryVariableSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of this library variable set.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        The space ID associated with this resource.
        """
        return pulumi.get(self, "space_id")

    @property
    @pulumi.getter
    def templates(self) -> pulumi.Output[Optional[Sequence['outputs.LibraryVariableSetTemplate']]]:
        return pulumi.get(self, "templates")

    @property
    @pulumi.getter(name="variableSetId")
    def variable_set_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "variable_set_id")
