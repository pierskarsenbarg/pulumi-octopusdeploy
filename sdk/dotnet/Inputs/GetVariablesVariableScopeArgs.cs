// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetVariablesVariableScopeInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("channels")]
        private InputList<string>? _channels;
        public InputList<string> Channels
        {
            get => _channels ?? (_channels = new InputList<string>());
            set => _channels = value;
        }

        [Input("environments")]
        private InputList<string>? _environments;
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        [Input("machines")]
        private InputList<string>? _machines;
        public InputList<string> Machines
        {
            get => _machines ?? (_machines = new InputList<string>());
            set => _machines = value;
        }

        [Input("roles")]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("tenantTags")]
        private InputList<string>? _tenantTags;
        public InputList<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new InputList<string>());
            set => _tenantTags = value;
        }

        public GetVariablesVariableScopeInputArgs()
        {
        }
        public static new GetVariablesVariableScopeInputArgs Empty => new GetVariablesVariableScopeInputArgs();
    }
}
