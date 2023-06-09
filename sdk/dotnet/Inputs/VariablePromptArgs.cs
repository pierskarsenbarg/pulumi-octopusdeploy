// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Inputs
{

    public sealed class VariablePromptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this variable prompt option.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("displaySettings")]
        public Input<Inputs.VariablePromptDisplaySettingsArgs>? DisplaySettings { get; set; }

        [Input("isRequired")]
        public Input<bool>? IsRequired { get; set; }

        [Input("label")]
        public Input<string>? Label { get; set; }

        public VariablePromptArgs()
        {
        }
        public static new VariablePromptArgs Empty => new VariablePromptArgs();
    }
}
