// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class VariablePromptDisplaySettingsSelectOptionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public VariablePromptDisplaySettingsSelectOptionGetArgs()
        {
        }
        public static new VariablePromptDisplaySettingsSelectOptionGetArgs Empty => new VariablePromptDisplaySettingsSelectOptionGetArgs();
    }
}